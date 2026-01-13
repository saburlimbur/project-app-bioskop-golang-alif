package usecase

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"alfdwirhmn/bioskop/internal/data/repository"
	"alfdwirhmn/bioskop/internal/dto"
	"alfdwirhmn/bioskop/pkg/utils"
	"context"

	"go.uber.org/zap"
)

type CinemasServiceCase interface {
	FindAll(ctx context.Context, page, limit int) (*[]entity.Cinemas, *dto.Pagination, error)
	FindById(ctx context.Context, id int) (*entity.Cinemas, error)
	SeatAvailability(ctx context.Context, cinemaID int, date string, time string) (*dto.SeatAvailabilityResponse, error)
}

type cinemaServiceCase struct {
	Repo   repository.Repository
	Logger *zap.Logger
	Config utils.Configuration
}

func NewCinemaServiceCase(repo repository.Repository, log *zap.Logger, conf utils.Configuration) CinemasServiceCase {
	return &cinemaServiceCase{
		Repo:   repo,
		Logger: log,
		Config: conf,
	}
}

func (cs *cinemaServiceCase) FindAll(ctx context.Context, page, limit int) (*[]entity.Cinemas, *dto.Pagination, error) {
	cinemas, total, err := cs.Repo.CinemaRepo.Lists(ctx, page, limit)

	if err != nil {
		return nil, nil, err
	}

	pagination := dto.Pagination{
		Page:       page,
		Limit:      limit,
		TotalPages: utils.TotalPage(limit, int64(total)),
		TotalRows:  total,
	}

	return &cinemas, &pagination, nil
}

func (cs *cinemaServiceCase) FindById(ctx context.Context, id int) (*entity.Cinemas, error) {
	return cs.Repo.CinemaRepo.FindById(ctx, id)
}

func (cs *cinemaServiceCase) SeatAvailability(ctx context.Context, cinemaID int, date string, time string) (*dto.SeatAvailabilityResponse, error) {

	seats, err := cs.Repo.CinemaRepo.SeatAvailability(ctx, cinemaID, date, time)
	if err != nil {
		return nil, err
	}

	available := 0
	for _, s := range seats {
		if s.Status == "available" {
			available++
		}
	}

	total := len(seats)

	return &dto.SeatAvailabilityResponse{
		CinemaID:       cinemaID,
		Date:           date,
		Time:           time,
		TotalSeats:     total,
		AvailableSeats: available,
		BookedSeats:    total - available,
		Seats:          seats,
	}, nil
}
