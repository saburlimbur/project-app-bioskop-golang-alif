package repository

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"alfdwirhmn/bioskop/pkg/database"
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
)

type OTPRepository interface {
	CreateOTP(ctx context.Context, otp *entity.OTPVerification) error
	VerifyOTP(ctx context.Context, userID int, code string) (int, error)
}

type otpRepo struct {
	DB   database.PgxIface
	Logg *zap.Logger
}

func NewOTPRepository(db database.PgxIface, log *zap.Logger) OTPRepository {
	return &otpRepo{
		DB:   db,
		Logg: log,
	}
}

func (or *otpRepo) CreateOTP(ctx context.Context, otp *entity.OTPVerification) error {
	query := `
		INSERT INTO otp_verifications (user_id, otp_code, expires_at)
		VALUES ($1, $2, $3) 
	`

	_, err := or.DB.Exec(
		ctx,
		query,
		otp.UserID,
		otp.OTPCode,
		otp.ExpiresAt,
	)
	return err
}

func (ur *otpRepo) VerifyOTP(ctx context.Context, userID int, code string) (int, error) {

	var id int
	var expiresAt time.Time
	var isUsed bool

	query := `
		SELECT id, expires_at, is_used
		FROM otp_verifications
		WHERE user_id=$1 AND otp_code=$2
		FOR UPDATE
	`

	err := ur.DB.QueryRow(ctx, query, userID, code).
		Scan(&id, &expiresAt, &isUsed)

	if err != nil {
		return 0, err
	}

	if isUsed || time.Now().After(expiresAt) {
		return 0, errors.New("otp expired or already used")
	}

	_, err = ur.DB.Exec(
		ctx,
		`UPDATE otp_verifications SET is_used=true WHERE id=$1`,
		id,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}
