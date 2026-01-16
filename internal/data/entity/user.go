package entity

import "time"

type Users struct {
	ID          int
	Username    string `json:"username" db:"username"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"-" db:"password_hash"`
	FullName    string `json:"full_name" db:"full_name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`

	IsVerified      bool       `db:"is_verified"`
	EmailVerifiedAt *time.Time `db:"email_verified_at"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
