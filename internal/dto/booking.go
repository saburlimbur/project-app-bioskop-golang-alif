package dto

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"time"
)

type BookingResponse struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	ShowtimeID    int       `json:"showtime_id"`
	SeatID        int       `json:"seat_id"`
	BookingCode   string    `json:"booking_code"`
	BookingStatus string    `json:"booking_status"`
	TotalPrice    float64   `json:"total_price"`
	ExpiredAt     time.Time `json:"expired_at"`
	CreatedAt     time.Time `json:"created_at"`
}

type UserBookingResponse struct {
	ID            int       `json:"id"`
	BookingCode   string    `json:"booking_code"`
	TotalPrice    float64   `json:"total_price"`
	BookingStatus string    `json:"booking_status"`
	CreatedAt     time.Time `json:"created_at"`

	CinemaName string `json:"cinema_name"`
	MovieTitle string `json:"movie_title"`
	ShowDate   string `json:"show_date"`
	ShowTime   string `json:"show_time"`

	PaymentStatus string  `json:"payment_status"`
	PaymentAmount float64 `json:"payment_amount"`
}

type CreateBookingRequest struct {
	ShowtimeID int `json:"showtime_id" validate:"required"`
	SeatID     int `json:"seat_id" validate:"required"`
}

func ToBookingResponse(booking *entity.Booking) BookingResponse {
	return BookingResponse{
		ID:            booking.ID,
		UserID:        booking.UserID,
		ShowtimeID:    booking.ShowtimeID,
		SeatID:        booking.SeatID,
		BookingCode:   booking.BookingCode,
		BookingStatus: booking.BookingStatus,
		TotalPrice:    booking.TotalPrice,
		ExpiredAt:     booking.ExpiredAt,
		CreatedAt:     booking.CreatedAt,
	}
}
