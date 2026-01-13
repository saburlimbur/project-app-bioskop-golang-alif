package entity

import "time"

type Cinemas struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Address     string `db:"address" json:"address"`
	City        string `db:"city" json:"city"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	TotalSeats  int    `db:"total_seats" json:"total_seats"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
