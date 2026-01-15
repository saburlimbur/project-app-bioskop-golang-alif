package repository

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"alfdwirhmn/bioskop/internal/dto"
	"alfdwirhmn/bioskop/pkg/database"
	"context"
	"database/sql"
	"errors"
	"time"

	"go.uber.org/zap"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, bk *entity.Booking) (*entity.Booking, error)
	FindBookingById(ctx context.Context, id int) (*entity.Booking, error)
	UpdateBookingStatus(ctx context.Context, id int, status string) error
	FindUserBookingHistory(ctx context.Context, userID int) ([]dto.UserBookingResponse, error)

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

func (br *bookingRepo) FindUserBookingHistory(ctx context.Context, userID int) ([]dto.UserBookingResponse, error) {

	query := `
			SELECT
			    b.id,
			    b.booking_code,
			    b.booking_status,
			    b.total_price,
			    b.expired_at,
			    b.created_at,

			    m.title,
			    c.name,

			    s.show_date,
			    s.show_time,

			    st.seat_number,
			    st.row_number,
			    st.seat_type,

			    CASE
			        WHEN p.id IS NULL THEN 'pending'
			        ELSE p.payment_status
			    END AS payment_status,

			    COALESCE(p.amount, 0)
			FROM bookings b
			JOIN showtimes s ON s.id = b.showtime_id
			JOIN cinemas c ON c.id = s.cinema_id
			JOIN movies m ON m.id = s.movie_id
			JOIN seats st ON st.id = b.seat_id
			LEFT JOIN payments p ON p.booking_id = b.id
			WHERE b.user_id = $1
			ORDER BY b.created_at DESC;
	`

	rows, err := br.DB.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dto.UserBookingResponse

	for rows.Next() {
		var r dto.UserBookingResponse
		var showDate time.Time
		var showTime time.Time

		err := rows.Scan(
			&r.ID,
			&r.BookingCode,
			&r.BookingStatus,
			&r.TotalPrice,
			&r.ExpiredAt,
			&r.CreatedAt,

			&r.MovieTitle,
			&r.CinemaName,

			&showDate,
			&showTime,

			&r.SeatNumber,
			&r.SeatRow,
			&r.SeatType,

			&r.PaymentStatus,
			&r.PaymentAmount,
		)
		if err != nil {
			return nil, err
		}

		r.ShowDate = showDate.Format("2006-01-02")
		r.ShowTime = showTime.Format("15:04")

		result = append(result, r)
	}

	return result, nil
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
