package repository

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"alfdwirhmn/bioskop/pkg/database"
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type SessionRepository interface {
	FindValidSession(ctx context.Context, token string) (*entity.Session, error)
	Create(ctx context.Context, s *entity.Session) error
	Revoke(ctx context.Context, token string) error
}

type Session struct {
	ID        int
	UserID    int
	Token     string
	ExpiresAt time.Time
}

type sessionRepository struct {
	DB  database.PgxIface
	Log *zap.Logger
}

func NewSessionRepository(db database.PgxIface, log *zap.Logger) SessionRepository {
	return &sessionRepository{
		DB:  db,
		Log: log,
	}
}

func (r *sessionRepository) FindValidSession(ctx context.Context, token string) (*entity.Session, error) {

	var s entity.Session

	query := `
		SELECT id, user_id, token, expires_at, revoked_at,
		       ip_address, device_info, created_at
		FROM sessions
		WHERE token = $1
		  AND expires_at > NOW()
		  AND revoked_at IS NULL
	`

	err := r.DB.QueryRow(ctx, query, token).Scan(
		&s.ID,
		&s.UserID,
		&s.Token,
		&s.ExpiresAt,
		&s.RevokedAt,
		&s.IPAddress,
		&s.DeviceInfo,
		&s.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("session not found or invalid")
		}
		return nil, err
	}

	return &s, nil
}

func (r *sessionRepository) Create(ctx context.Context, s *entity.Session) error {
	query := `
		INSERT INTO sessions (
			user_id, token, expires_at,
			ip_address, device_info
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.DB.Exec(ctx, query,
		s.UserID,
		s.Token,
		s.ExpiresAt,
		s.IPAddress,
		s.DeviceInfo,
	)

	return err
}

func (r *sessionRepository) Revoke(ctx context.Context, token string) error {
	query := `
		UPDATE sessions
		SET revoked_at = NOW()
		WHERE token = $1
		  AND revoked_at IS NULL
	`
	_, err := r.DB.Exec(ctx, query, token)
	return err
}
