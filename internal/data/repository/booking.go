package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type BookingRepository interface {
}

type bookingRepo struct {
	db  *sql.DB
	log *zap.Logger
}

func NewBookingRepository(db *sql.DB, log *zap.Logger) BookingRepository {
	return &bookingRepo{
		db:  db,
		log: log,
	}
}
