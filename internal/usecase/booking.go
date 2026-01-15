package usecase

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"alfdwirhmn/bioskop/internal/data/repository"
	"alfdwirhmn/bioskop/internal/dto"
	"alfdwirhmn/bioskop/pkg/database"
	"alfdwirhmn/bioskop/pkg/utils"
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type BookingServiceCase interface {
	CreateBooking(ctx context.Context, userID int, req dto.CreateBookingRequest) (*entity.Booking, error)
	FindByID(ctx context.Context, id int) (*entity.Booking, error)
}

type bookingServiceCase struct {
	Repo      repository.Repository
	TxManager database.TxManager
	Logger    *zap.Logger
	Config    utils.Configuration
}

func NewBookingServiceCase(repo repository.Repository, txMgr database.TxManager, log *zap.Logger, conf utils.Configuration) BookingServiceCase {
	return &bookingServiceCase{
		Repo:      repo,
		TxManager: txMgr,
		Logger:    log,
		Config:    conf,
	}
}

func (bs *bookingServiceCase) CreateBooking(ctx context.Context, userID int, req dto.CreateBookingRequest) (*entity.Booking, error) {

	var result *entity.Booking

	err := bs.TxManager.WithTransaction(ctx, func(tx pgx.Tx) error {
		// repo versi transaction
		repoTx := bs.Repo.WithTx(tx)

		// cek showtime
		showtime, err := repoTx.ShowtimeRepo.FindByID(ctx, req.ShowtimeID)
		if err != nil {
			return err
		}

		// cek seat sudah dibooking atau belum
		exists, err := repoTx.BookingRepo.IsSeatBooked(
			ctx,
			req.ShowtimeID,
			req.SeatID,
		)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("seat already booked")
		}

		booking := &entity.Booking{
			UserID:        userID,
			ShowtimeID:    req.ShowtimeID,
			SeatID:        req.SeatID,
			BookingCode:   utils.GenerateBookingCode(),
			BookingStatus: "pending",
			TotalPrice:    showtime.Price,
			ExpiredAt:     time.Now().Add(15 * time.Minute),
		}

		// insert booking
		created, err := repoTx.BookingRepo.CreateBooking(ctx, booking)
		if err != nil {
			return err
		}

		result = created
		return nil
	})

	if err != nil {
		bs.Logger.Error("failed create booking", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (bs *bookingServiceCase) FindByID(ctx context.Context, id int) (*entity.Booking, error) {

	booking, err := bs.Repo.BookingRepo.FindBookingById(ctx, id)
	if err != nil {
		bs.Logger.Error("booking not found", zap.Error(err))
		return nil, err
	}

	return booking, nil
}
