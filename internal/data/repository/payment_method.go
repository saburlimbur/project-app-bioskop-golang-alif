package repository

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"alfdwirhmn/bioskop/pkg/database"
	"context"
	"errors"

	"go.uber.org/zap"
)

type PaymentMethodRepository interface {
	FindByID(ctx context.Context, id int) (*entity.PaymentMethod, error)
	ListActive(ctx context.Context) ([]entity.PaymentMethod, error)
}

type paymentMethodRepo struct {
	DB   database.PgxIface
	Logg *zap.Logger
}

func NewPaymentMethodRepository(db database.PgxIface, log *zap.Logger) PaymentMethodRepository {
	return &paymentMethodRepo{
		DB:   db,
		Logg: log,
	}
}

func (pmr *paymentMethodRepo) ListActive(ctx context.Context) ([]entity.PaymentMethod, error) {
	query := `
		SELECT id, name, code
		FROM payment_methods
		WHERE is_active = true
		ORDER BY name ASC
	`

	rows, err := pmr.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var methods []entity.PaymentMethod
	for rows.Next() {
		var pm entity.PaymentMethod
		if err := rows.Scan(&pm.ID, &pm.Name, &pm.Code); err != nil {
			return nil, err
		}
		methods = append(methods, pm)
	}

	return methods, nil
}

func (pmr *paymentMethodRepo) FindByID(ctx context.Context, id int) (*entity.PaymentMethod, error) {
	query := `
		SELECT id, name, code, description, is_active, created_at
		FROM payment_methods
		WHERE id = $1
	`

	var pm entity.PaymentMethod
	err := pmr.DB.QueryRow(ctx, query, id).Scan(
		&pm.ID,
		&pm.Name,
		&pm.Code,
		&pm.Description,
		&pm.IsActive,
		&pm.CreatedAt,
	)

	if err != nil {
		return nil, errors.New("payment method not found")
	}

	return &pm, nil
}
