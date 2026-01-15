package repository

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"alfdwirhmn/bioskop/pkg/database"
	"context"
	"errors"

	"go.uber.org/zap"
)

type ShowtimeRepository interface {
	FindByID(ctx context.Context, id int) (*entity.Showtime, error)
}

type showtimeRepo struct {
	DB   database.PgxIface
	Logg *zap.Logger
}

func NewShowtimeRepository(db database.PgxIface, log *zap.Logger) ShowtimeRepository {
	return &showtimeRepo{
		DB:   db,
		Logg: log,
	}
}

func (sr *showtimeRepo) FindByID(ctx context.Context, id int) (*entity.Showtime, error) {
	query := `
		SELECT id, cinema_id, movie_id, show_date, show_time, price
		FROM showtimes
		WHERE id = $1
	`

	var s entity.Showtime
	err := sr.DB.QueryRow(ctx, query, id).Scan(
		&s.ID,
		&s.CinemaID,
		&s.MovieID,
		&s.ShowDate,
		&s.ShowTime,
		&s.Price,
	)
	if err != nil {
		return nil, errors.New("showtime not found")
	}

	return &s, nil
}
