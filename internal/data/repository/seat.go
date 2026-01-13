package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type SeatRepository interface {
}

type seatRepo struct {
	db  *sql.DB
	log *zap.Logger
}

func NewSeatRepository(db *sql.DB, log *zap.Logger) SeatRepository {
	return &seatRepo{
		db:  db,
		log: log,
	}
}
