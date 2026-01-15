package entity

import "time"

type Showtime struct {
	ID        int       `db:"id" json:"id"`
	CinemaID  int       `db:"cinema_id" json:"cinema_id"`
	MovieID   int       `db:"movie_id" json:"movie_id"`
	ShowDate  time.Time `db:"show_date" json:"show_date"` // DATE
	ShowTime  time.Time `db:"show_time" json:"show_time"` // TIME
	Price     float64   `db:"price" json:"price"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
