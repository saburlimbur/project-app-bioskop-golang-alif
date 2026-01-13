package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	UserRepo   UserRepository
	CinemaRepo CinemaRepository

	SessionRepo SessionRepository
}

func NewRepository(db *pgxpool.Pool, log *zap.Logger) Repository {
	return Repository{
		UserRepo:    NewUserRepository(db, log),
		SessionRepo: NewSessionRepository(db, log),
		CinemaRepo:  NewCinemaRepository(db, log),
	}
}
