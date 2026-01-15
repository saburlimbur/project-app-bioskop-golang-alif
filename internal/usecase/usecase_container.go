package usecase

import (
	"alfdwirhmn/bioskop/internal/data/repository"
	"alfdwirhmn/bioskop/pkg/database"
	"alfdwirhmn/bioskop/pkg/utils"

	"go.uber.org/zap"
)

type Service struct {
	UserService    UserServiceCase
	CinemaService  CinemasServiceCase
	BookingService BookingServiceCase
	PaymentService PaymentServiceCase
}

func NewService(
	repo repository.Repository,
	txMgr database.TxManager,
	log *zap.Logger,
	conf utils.Configuration,
) Service {
	return Service{
		UserService:    NewUserServiceCase(repo, log, conf),
		CinemaService:  NewCinemaServiceCase(repo, log, conf),
		BookingService: NewBookingServiceCase(repo, txMgr, log, conf),
		PaymentService: NewPaymentServiceCase(repo, txMgr, log, conf),
	}
}
