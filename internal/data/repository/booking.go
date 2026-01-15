package repository

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"alfdwirhmn/bioskop/pkg/database"
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, bk *entity.Booking) (*entity.Booking, error)
	FindBookingById(ctx context.Context, id int) (*entity.Booking, error)
	UpdateBookingStatus(ctx context.Context, id int, status string) error
	IsSeatBooked(ctx context.Context, showtimeID int, seatID int) (bool, error)
}

type bookingRepo struct {
	// db  *sql.DB
	DB   database.PgxIface
	Logg *zap.Logger
}

func NewBookingRepository(db database.PgxIface, log *zap.Logger) BookingRepository {
	return &bookingRepo{
		DB:   db,
		Logg: log,
	}
}

func (br *bookingRepo) CreateBooking(ctx context.Context, bk *entity.Booking) (*entity.Booking, error) {
	query := `
		INSERT INTO bookings (user_id, showtime_id, seat_id, booking_code, booking_status, total_price, expired_at)
		VALUES ($1, $2, $3, $4, COALESCE($5, 'pending'), $6, $7)
		RETURNING
		    id,
		    user_id,
		    showtime_id,
		    seat_id,
		    booking_code,
		    booking_status,
		    total_price,
		    expired_at,
		    created_at,
		    updated_at
			`

	var booking entity.Booking
	err := br.DB.QueryRow(ctx, query,
		bk.UserID,
		bk.ShowtimeID,
		bk.SeatID,
		bk.BookingCode,
		bk.BookingStatus,
		bk.TotalPrice,
		bk.ExpiredAt,
	).Scan(
		&booking.ID,
		&booking.UserID,
		&booking.ShowtimeID,
		&booking.SeatID,
		&booking.BookingCode,
		&booking.BookingStatus,
		&booking.TotalPrice,
		&booking.ExpiredAt,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		br.Logg.Info("failed to booking seats", zap.Error(err))
		return nil, err
	}

	br.Logg.Info("booking seats created successfully")
	return &booking, nil
}

func (br *bookingRepo) FindBookingById(ctx context.Context, id int) (*entity.Booking, error) {
	query := `
		SELECT id, user_id, showtime_id, seat_id, booking_code, booking_status, total_price, expired_at, created_at, updated_at
		FROM bookings
		WHERE id = $1
	`

	var bk entity.Booking
	err := br.DB.QueryRow(ctx, query, id).Scan(
		&bk.ID,
		&bk.UserID,
		&bk.ShowtimeID,
		&bk.SeatID,
		&bk.BookingCode,
		&bk.BookingStatus,
		&bk.TotalPrice,
		&bk.ExpiredAt,
		&bk.CreatedAt,
		&bk.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &bk, err
}

func (br *bookingRepo) UpdateBookingStatus(ctx context.Context, id int, status string) error {
	query := `
		UPDATE bookings
		SET booking_status = $1,
			updated_at = NOW()
		WHERE id = $2
	`

	res, err := br.DB.Exec(ctx, query, status, id)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return errors.New("booking not found")
	}

	return nil
}

func (br *bookingRepo) IsSeatBooked(ctx context.Context, showtimeID int, seatID int) (bool, error) {

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM bookings
			WHERE showtime_id = $1
			  AND seat_id = $2
			  AND booking_status IN ('pending', 'confirmed')
		)
	`

	var exists bool
	err := br.DB.QueryRow(ctx, query, showtimeID, seatID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
