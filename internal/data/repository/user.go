package repository

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"alfdwirhmn/bioskop/pkg/database"
	"context"

	"go.uber.org/zap"
)

type UserRepository interface {
	Register(ctx context.Context, user *entity.Users) (*entity.Users, error)
	FindByIdentifier(ctx context.Context, identifier string) (*entity.Users, error)
	FindByEmail(ctx context.Context, email string) (*entity.Users, error)
	VerifyEmail(ctx context.Context, userID int) error
	IsEmailVerified(ctx context.Context, userID int) (bool, error)

	IsEmailExists(ctx context.Context, email string, excludeID int) (bool, error)
	IsUsernameExists(ctx context.Context, username string, excludeID int) (bool, error)
}

type userRepo struct {
	DB   database.PgxIface
	Logg *zap.Logger
}

func NewUserRepository(db database.PgxIface, log *zap.Logger) UserRepository {
	return &userRepo{
		DB:   db,
		Logg: log,
	}
}

func (ur *userRepo) Register(ctx context.Context, user *entity.Users) (*entity.Users, error) {
	query := `
		INSERT INTO users (username, email, password, full_name, phone_number)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, username, email, password, full_name, phone_number, is_verified, email_verified_at, created_at, updated_at
	`

	var usr entity.Users
	err :=
		ur.DB.QueryRow(
			ctx,
			query,
			user.Username,
			user.Email,
			user.Password,
			user.FullName,
			user.PhoneNumber,
		).Scan(
			&usr.ID,
			&usr.Username,
			&usr.Email,
			&usr.Password,
			&usr.FullName,
			&usr.PhoneNumber,
			&usr.IsVerified,
			&usr.EmailVerifiedAt,
			&usr.CreatedAt,
			&usr.UpdatedAt,
		)

	if err != nil {
		ur.Logg.Error("Failed to register", zap.Error(err))
		return nil, err
	}

	ur.Logg.Info("Registered succesfully")
	return &usr, nil
}

func (r *userRepo) FindByIdentifier(ctx context.Context, identifier string) (*entity.Users, error) {
	query := `
		SELECT id, username, email, password, full_name, phone_number, email_verified_at, is_verified, created_at, updated_at
		FROM users
		WHERE email = $1 OR username = $1
	`

	var user entity.Users
	err := r.DB.QueryRow(
		ctx,
		query,
		identifier,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.PhoneNumber,
		&user.EmailVerifiedAt,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepo) FindByEmail(ctx context.Context, email string) (*entity.Users, error) {
	query := `
		SELECT id, email, email_verified_at
		FROM users
		WHERE email=$1
	`

	var user entity.Users
	err := ur.DB.QueryRow(ctx, query, email).
		Scan(&user.ID, &user.Email, &user.EmailVerifiedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepo) VerifyEmail(ctx context.Context, userID int) error {
	query := `
		UPDATE users
		SET email_verified_at = NOW(),
			is_verified = TRUE,
			updated_at = NOW()
		WHERE id = $1
	`
	_, err := ur.DB.Exec(ctx, query, userID)
	return err
}

func (ur *userRepo) IsEmailVerified(ctx context.Context, userID int) (bool, error) {
	query := `
		SELECT email_verified_at IS NOT NULL
		FROM users
		WHERE id = $1
	`

	var verified bool
	err := ur.DB.QueryRow(ctx, query, userID).Scan(&verified)
	return verified, err
}

func (ur *userRepo) IsEmailExists(ctx context.Context, email string, excludeID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND id != $2)`

	var exists bool
	err := ur.DB.QueryRow(
		ctx,
		query,
		email,
		excludeID,
	).Scan(&exists)
	if err != nil {
		ur.Logg.Error("Failed to check email existence", zap.Error(err))
		return false, err
	}

	return exists, nil
}

func (ur *userRepo) IsUsernameExists(ctx context.Context, username string, excludeID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 AND id != $2)`

	var exists bool
	err := ur.DB.QueryRow(
		ctx,
		query,
		username,
		excludeID,
	).Scan(&exists)
	if err != nil {
		ur.Logg.Error("Failed to check username existence", zap.Error(err))
		return false, err
	}

	return exists, nil
}
