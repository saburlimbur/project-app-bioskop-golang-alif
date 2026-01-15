package entity

import "time"

type Payment struct {
	ID              int        `db:"id" json:"id"`
	BookingID       int        `db:"booking_id" json:"booking_id"`
	PaymentMethodID int        `db:"payment_method_id" json:"-"`
	Amount          float64    `db:"amount" json:"amount"`
	PaymentStatus   string     `db:"payment_status" json:"payment_status"`
	PaymentDate     *time.Time `db:"payment_date" json:"payment_date,omitempty"`
	TransactionID   *string    `db:"transaction_id" json:"transaction_id,omitempty"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updated_at"`

	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`
}

type PaymentMethod struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Code        string    `db:"code" json:"code"`
	Description string    `db:"description" json:"description"`
	IsActive    bool      `db:"is_active" json:"-"`
	CreatedAt   time.Time `db:"created_at" json:"-"`
}
