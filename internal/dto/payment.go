package dto

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"time"
)

type PaymentResponse struct {
	ID              int        `json:"id"`
	BookingID       int        `json:"booking_id"`
	PaymentMethodID int        `json:"payment_method_id"`
	Amount          float64    `json:"amount"`
	PaymentStatus   string     `json:"payment_status"`
	PaymentDate     *time.Time `json:"payment_date,omitempty"`
	TransactionID   *string    `json:"transaction_id,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`

	PaymentMethod PaymentMethodResponse `json:"payment_method"`
}

type PaymentMethodResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type CreatePaymentRequest struct {
	BookingID       int `json:"booking_id" validate:"required,gt=0"`
	PaymentMethodID int `json:"payment_method_id" validate:"required,gt=0"`
}

func ToPaymentResponse(pay *entity.Payment) PaymentResponse {
	return PaymentResponse{
		ID:              pay.ID,
		BookingID:       pay.BookingID,
		PaymentMethodID: pay.PaymentMethodID,
		Amount:          pay.Amount,
		PaymentStatus:   pay.PaymentStatus,
		PaymentDate:     pay.PaymentDate,
		TransactionID:   pay.TransactionID,
		CreatedAt:       pay.CreatedAt,
		PaymentMethod: PaymentMethodResponse{
			ID:   pay.PaymentMethod.ID,
			Name: pay.PaymentMethod.Name,
			Code: pay.PaymentMethod.Code,
		},
	}
}

func ToPaymentMethodResponses(methods []entity.PaymentMethod) []PaymentMethodResponse {
	res := make([]PaymentMethodResponse, 0, len(methods))
	for _, m := range methods {
		res = append(res, PaymentMethodResponse{
			ID:   m.ID,
			Name: m.Name,
			Code: m.Code,
		})
	}
	return res
}
