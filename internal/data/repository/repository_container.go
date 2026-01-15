package repository

import (
	"alfdwirhmn/bioskop/pkg/database"

	"go.uber.org/zap"
)

type Repository struct {
	DB   database.PgxIface
	Logg *zap.Logger

	UserRepo          UserRepository
	CinemaRepo        CinemaRepository
	BookingRepo       BookingRepository
	ShowtimeRepo      ShowtimeRepository
	PaymentRepo       PaymentRepository
	PaymentMethodRepo PaymentMethodRepository

	SessionRepo SessionRepository
}

// constructor utama untuk abstract query
func NewRepository(db database.PgxIface, log *zap.Logger) Repository {
	return Repository{
		DB:   db,
		Logg: log,

		UserRepo:          NewUserRepository(db, log),
		SessionRepo:       NewSessionRepository(db, log),
		CinemaRepo:        NewCinemaRepository(db, log),
		ShowtimeRepo:      NewShowtimeRepository(db, log),
		BookingRepo:       NewBookingRepository(db, log),
		PaymentRepo:       NewPaymentRepository(db, log),
		PaymentMethodRepo: NewPaymentMethodRepository(db, log),
	}
}

// transaction
func (r Repository) WithTx(tx database.PgxIface) Repository {
	return Repository{
		DB:   tx,
		Logg: r.Logg,

		UserRepo:          NewUserRepository(tx, r.Logg),
		SessionRepo:       NewSessionRepository(tx, r.Logg),
		CinemaRepo:        NewCinemaRepository(tx, r.Logg),
		ShowtimeRepo:      NewShowtimeRepository(tx, r.Logg),
		BookingRepo:       NewBookingRepository(tx, r.Logg),
		PaymentRepo:       NewPaymentRepository(tx, r.Logg),
		PaymentMethodRepo: NewPaymentMethodRepository(tx, r.Logg),
	}
}
