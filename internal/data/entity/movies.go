package entity

import "time"

type Movie struct {
	ID          int     `db:"id" json:"id"`
	Title       string  `db:"title" json:"title"`
	Genre       *string `db:"genre" json:"genre,omitempty"`
	Duration    int     `db:"duration" json:"duration"` // menit
	Rating      *string `db:"rating" json:"rating,omitempty"`
	Description *string `db:"description" json:"description,omitempty"`
	Thumbnail   string  `db:"thumbnail" json:"thumbnail"`
	Poster      *string `db:"poster" json:"poster,omitempty"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
