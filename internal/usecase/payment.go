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

type PaymentServiceCase interface {
	CreatePayment(ctx context.Context, userID int, req dto.CreatePaymentRequest) (*entity.Payment, error)
	ListPaymentMethods(ctx context.Context) ([]entity.PaymentMethod, error)
}

type paymentServiceCase struct {
	Repo      repository.Repository
	TxManager database.TxManager
	Logger    *zap.Logger
	Config    utils.Configuration
}

func NewPaymentServiceCase(repo repository.Repository, txMgr database.TxManager, log *zap.Logger, conf utils.Configuration) PaymentServiceCase {
	return &paymentServiceCase{
		Repo:      repo,
		TxManager: txMgr,
		Logger:    log,
		Config:    conf,
	}
}

func (ps *paymentServiceCase) CreatePayment(ctx context.Context, userID int, req dto.CreatePaymentRequest) (*entity.Payment, error) {

	var result *entity.Payment

	err := ps.TxManager.WithTransaction(ctx, func(tx pgx.Tx) error {
		repoTx := ps.Repo.WithTx(tx)

		booking, err := repoTx.BookingRepo.FindBookingById(ctx, req.BookingID)
		if err != nil {
			return err
		}

		if booking.BookingStatus != "pending" {
			return errors.New("booking already paid or cancelled")
		}

		// cek payment method yang ada dan aktif
		method, err := repoTx.PaymentMethodRepo.FindByID(ctx, req.PaymentMethodID)
		if err != nil {
			return err
		}
		if !method.IsActive {
			return errors.New("payment method is inactive")
		}

		now := time.Now()
		trxID := utils.GenerateTransactionID()

		payment := &entity.Payment{
			BookingID:       booking.ID,
			PaymentMethodID: method.ID,
			Amount:          booking.TotalPrice,
			PaymentStatus:   "success",
			PaymentDate:     &now,
			TransactionID:   &trxID,
		}

		created, err := repoTx.PaymentRepo.CreatePayment(ctx, payment)
		if err != nil {
			return err
		}

		// transaction update booking to confirmed
		err = repoTx.BookingRepo.UpdateBookingStatus(ctx, booking.ID, "confirmed")
		if err != nil {
			return err
		}

		// ambil payment dan method
		full, err := repoTx.PaymentRepo.FindPaymentWithMethod(ctx, created.ID)
		if err != nil {
			return err
		}

		result = full
		return nil
	})

	if err != nil {
		ps.Logger.Error("create payment failed", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (ps *paymentServiceCase) ListPaymentMethods(ctx context.Context) ([]entity.PaymentMethod, error) {
	return ps.Repo.PaymentMethodRepo.ListActive(ctx)
}
