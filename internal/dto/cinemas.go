package dto

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"time"
)

type CinemaResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	PhoneNumber string    `json:"phone_number"`
	TotalSeats  int       `json:"total_seats"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateCinemaRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Address     string `json:"address" validate:"required"`
	City        string `json:"city" validate:"required,min=3,max=50"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,min=10,max=20"`
	TotalSeats  int    `json:"total_seats" validate:"required,min=1"`
}

type UpdateCinemaRequest struct {
	Name        string `json:"name" validate:"omitempty,min=3,max=100"`
	Address     string `json:"address" validate:"omitempty"`
	City        string `json:"city" validate:"omitempty,min=3,max=50"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,min=10,max=20"`
	TotalSeats  int    `json:"total_seats" validate:"omitempty,min=1"`
}

func ToCinemaResponse(cinema *entity.Cinemas) CinemaResponse {
	return CinemaResponse{
		ID:          cinema.ID,
		Name:        cinema.Name,
		Address:     cinema.Address,
		City:        cinema.City,
		PhoneNumber: cinema.PhoneNumber,
		TotalSeats:  cinema.TotalSeats,
		CreatedAt:   cinema.CreatedAt,
	}
}
