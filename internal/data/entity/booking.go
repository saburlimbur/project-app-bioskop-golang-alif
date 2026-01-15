package entity

import "time"

type Booking struct {
	ID            int       `db:"id" json:"id"`
	UserID        int       `db:"user_id" json:"user_id"`
	ShowtimeID    int       `db:"showtime_id" json:"showtime_id"`
	SeatID        int       `db:"seat_id" json:"seat_id"`
	BookingCode   string    `db:"booking_code" json:"booking_code"`
	BookingStatus string    `db:"booking_status" json:"booking_status"`
	TotalPrice    float64   `db:"total_price" json:"total_price"`
	ExpiredAt     time.Time `db:"expired_at" json:"expired_at"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
