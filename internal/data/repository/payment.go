package repository

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"alfdwirhmn/bioskop/pkg/database"
	"context"

	"go.uber.org/zap"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, py *entity.Payment) (*entity.Payment, error)
	FindPaymentWithMethod(ctx context.Context, paymentID int) (*entity.Payment, error)
}

type paymentRepo struct {
	DB   database.PgxIface
	Logg *zap.Logger
}

func NewPaymentRepository(db database.PgxIface, log *zap.Logger) PaymentRepository {
	return &paymentRepo{
		DB:   db,
		Logg: log,
	}
}

func (pr *paymentRepo) CreatePayment(ctx context.Context, py *entity.Payment) (*entity.Payment, error) {
	query := `
		INSERT INTO payments (booking_id, payment_method_id, amount, payment_status, payment_date, transaction_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING
		id, booking_id, payment_method_id, amount, payment_status,
		payment_date, transaction_id, created_at
	`

	var payment entity.Payment
	err := pr.DB.QueryRow(ctx, query,
		py.BookingID,
		py.PaymentMethodID,
		py.Amount,
		py.PaymentStatus,
		py.PaymentDate,
		py.TransactionID,
	).Scan(
		&payment.ID,
		&payment.BookingID,
		&payment.PaymentMethodID,
		&payment.Amount,
		&payment.PaymentStatus,
		&payment.PaymentDate,
		&payment.TransactionID,
		&payment.CreatedAt,
	)

	if err != nil {
		pr.Logg.Error("failed to create payment", zap.Error(err))
		return nil, err
	}

	pr.Logg.Info("payment created successfully")
	return &payment, nil
}

func (pr *paymentRepo) FindPaymentMethodByID(
	ctx context.Context,
	id int,
) (*entity.PaymentMethod, error) {

	query := `
		SELECT id, name, code, is_active, created_at
		FROM payment_methods
		WHERE id = $1
	`

	var pm entity.PaymentMethod
	err := pr.DB.QueryRow(ctx, query, id).Scan(
		&pm.ID,
		&pm.Name,
		&pm.Code,
		&pm.IsActive,
		&pm.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &pm, nil
}

func (pr *paymentRepo) FindPaymentWithMethod(ctx context.Context, id int) (*entity.Payment, error) {

	query := `
		SELECT
			p.id,
			p.booking_id,
			p.amount,
			p.payment_status,
			p.payment_date,
			p.transaction_id,
			p.created_at,

			pm.id,
			pm.name,
			pm.code
		FROM payments p
		JOIN payment_methods pm ON pm.id = p.payment_method_id
		WHERE p.id = $1
	`

	var pay entity.Payment
	var method entity.PaymentMethod

	err := pr.DB.QueryRow(ctx, query, id).Scan(
		&pay.ID,
		&pay.BookingID,
		&pay.Amount,
		&pay.PaymentStatus,
		&pay.PaymentDate,
		&pay.TransactionID,
		&pay.CreatedAt,

		&method.ID,
		&method.Name,
		&method.Code,
	)

	if err != nil {
		return nil, err
	}

	pay.PaymentMethod = &method
	return &pay, nil
}
