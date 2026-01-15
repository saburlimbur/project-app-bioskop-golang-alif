package repository

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"alfdwirhmn/bioskop/internal/dto"
	"alfdwirhmn/bioskop/pkg/database"
	"context"

	"go.uber.org/zap"
)

type CinemaRepository interface {
	Lists(ctx context.Context, page, limit int) ([]entity.Cinemas, int, error)
	FindById(ctx context.Context, id int) (*entity.Cinemas, error)
	SeatAvailability(ctx context.Context, cinemaID int, date string, time string) ([]dto.SeatAvailabilityItem, error)
}

type cinemaRepo struct {
	db  database.PgxIface
	log *zap.Logger
}

func NewCinemaRepository(db database.PgxIface, log *zap.Logger) CinemaRepository {
	return &cinemaRepo{
		db:  db,
		log: log,
	}
}

func (cr *cinemaRepo) Lists(ctx context.Context, page, limit int) ([]entity.Cinemas, int, error) {

	offset := (page - 1) * limit

	// total data
	var total int
	countQuery := `SELECT COUNT(*) FROM cinemas`
	if err := cr.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		cr.log.Error("failed to count cinemas", zap.Error(err))
		return nil, 0, err
	}

	// list data
	query := `
		SELECT id, name, address, city, phone_number, total_seats, created_at, updated_at
		FROM cinemas
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := cr.db.Query(ctx, query, limit, offset)
	if err != nil {
		cr.log.Error("failed to fetch cinemas", zap.Error(err))
		return nil, 0, err
	}
	defer rows.Close()

	var cinemas []entity.Cinemas

	for rows.Next() {
		var c entity.Cinemas
		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Address,
			&c.City,
			&c.PhoneNumber,
			&c.TotalSeats,
			&c.CreatedAt,
			&c.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}

		cinemas = append(cinemas, c)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return cinemas, total, nil
}

func (cr *cinemaRepo) FindById(ctx context.Context, id int) (*entity.Cinemas, error) {
	query := `
		SELECT id, name, address, city, phone_number, total_seats, created_at, updated_at
		FROM cinemas
		WHERE id = $1
	`

	var cnm entity.Cinemas
	err := cr.db.QueryRow(ctx, query, id).Scan(
		&cnm.ID,
		&cnm.Name,
		&cnm.Address,
		&cnm.City,
		&cnm.PhoneNumber,
		&cnm.TotalSeats,
		&cnm.CreatedAt,
		&cnm.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &cnm, nil
}

func (cr *cinemaRepo) SeatAvailability(ctx context.Context, cinemaID int, date string, time string) ([]dto.SeatAvailabilityItem, error) {

	query := `
	SELECT
	  s.id,
	  s.row_number,
	  s.seat_number,
	  s.seat_type,
	  CASE
	    WHEN b.id IS NULL THEN 'available'
	    ELSE 'booked'
	  END AS status
	FROM seats s
	JOIN showtimes st ON st.cinema_id = s.cinema_id
	LEFT JOIN bookings b
	  ON b.seat_id = s.id
	  AND b.showtime_id = st.id
	  AND b.booking_status IN ('pending', 'confirmed')
	WHERE st.cinema_id = $1
	  AND st.show_date = $2
	  AND st.show_time = $3
	ORDER BY s.row_number, s.seat_number
	`

	rows, err := cr.db.Query(ctx, query, cinemaID, date, time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []dto.SeatAvailabilityItem

	for rows.Next() {
		var s dto.SeatAvailabilityItem
		if err := rows.Scan(
			&s.ID,
			&s.Row,
			&s.SeatNumber,
			&s.SeatType,
			&s.Status,
		); err != nil {
			return nil, err
		}
		seats = append(seats, s)
	}

	return seats, nil
}
