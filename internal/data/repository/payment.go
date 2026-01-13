package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type PaymentRepository interface {
}

type paymentRepo struct {
	db  *sql.DB
	log *zap.Logger
}

func NewPaymentRepository(db *sql.DB, log *zap.Logger) PaymentRepository {
	return &paymentRepo{
		db:  db,
		log: log,
	}
}
