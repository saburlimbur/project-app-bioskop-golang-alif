package usecase

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"alfdwirhmn/bioskop/internal/data/repository"
	"alfdwirhmn/bioskop/internal/dto"
	"alfdwirhmn/bioskop/pkg/utils"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceCase interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*entity.Users, error)
	Login(ctx context.Context, req dto.LoginRequest, ip, devinf string) (*dto.LoginResponse, error)
	Logout(ctx context.Context, token string) error
	VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (*dto.UserResponse, error)
}

type userServiceCase struct {
	Repo   repository.Repository
	Logger *zap.Logger
	Config utils.Configuration
}

func NewUserServiceCase(repo repository.Repository, log *zap.Logger, conf utils.Configuration) UserServiceCase {
	return &userServiceCase{
		Repo:   repo,
		Logger: log,
		Config: conf,
	}
}

func (us *userServiceCase) Register(ctx context.Context, req dto.RegisterRequest) (*entity.Users, error) {
	exists, err := us.Repo.UserRepo.IsEmailExists(ctx, req.Email, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already exists")
	}

	// cek username
	exists, err = us.Repo.UserRepo.IsUsernameExists(ctx, req.Username, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("username already exists")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &entity.Users{
		Username:    req.Username,
		Email:       req.Email,
		Password:    string(hashedPassword),
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
	}

	// return us.Repo.UserRepo.Register(ctx, user)
	tx, err := us.Repo.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	repoTx := us.Repo.WithTx(tx)

	newUser, err := repoTx.UserRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	// generate otp
	otpCode, err := utils.GenerateOTP()
	if err != nil {
		us.Logger.Error("failed to generate otp", zap.Error(err))
		return nil, err
	}

	// save to db
	otp := &entity.OTPVerification{
		UserID:    newUser.ID,
		OTPCode:   otpCode,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}
	if err := repoTx.OTPRepo.CreateOTP(ctx, otp); err != nil {
		return nil, err
	}

	// commit, so if it fails, the rollback runs and the user OTP is not stuck
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return newUser, nil

}

func (us *userServiceCase) Login(ctx context.Context, req dto.LoginRequest, ip, devinf string) (*dto.LoginResponse, error) {
	user, err := us.Repo.UserRepo.FindByIdentifier(ctx, req.Identifier)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsVerified {
		return nil, errors.New("email not verified")
	}

	// generate token
	token := uuid.NewString()

	expiresAt := time.Now().Add(24 * time.Hour)

	session := &entity.Session{
		UserID:     user.ID,
		Token:      token,
		ExpiresAt:  expiresAt,
		IPAddress:  ip,
		DeviceInfo: devinf,
	}

	if err := us.Repo.SessionRepo.Create(ctx, session); err != nil {
		us.Logger.Error("failed to create session", zap.Error(err))
		return nil, err
	}

	user.Password = ""

	return &dto.LoginResponse{
		User:      dto.ToUserResponse(user),
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

func (us *userServiceCase) Logout(ctx context.Context, token string) error {
	// cek session valid
	session, err := us.Repo.SessionRepo.FindValidSession(ctx, token)
	if err != nil {
		return errors.New("invalid session")
	}

	// revoke session
	if err := us.Repo.SessionRepo.Revoke(ctx, session.Token); err != nil {
		us.Logger.Error("failed to revoke session", zap.Error(err))
		return err
	}

	return nil
}

func (us *userServiceCase) VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (*dto.UserResponse, error) {

	user, err := us.Repo.UserRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.EmailVerifiedAt != nil {
		return nil, errors.New("email already verified")
	}

	tx, err := us.Repo.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	repoTx := us.Repo.WithTx(tx)

	// verify OTP
	if _, err := repoTx.OTPRepo.VerifyOTP(ctx, user.ID, req.OTP); err != nil {
		return nil, err
	}

	if err := repoTx.UserRepo.VerifyEmail(ctx, user.ID); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	verifiedUser, err := us.Repo.UserRepo.FindByIdentifier(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	return dto.ToUserResponse(verifiedUser), nil
}
