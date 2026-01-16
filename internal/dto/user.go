package dto

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"time"
)

type UserResponse struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	FullName      string `json:"full_name"`
	PhoneNumber   string `json:"phone_number"`
	EmailVerified bool   `json:"email_verified"`

	CreatedAt time.Time `json:"created_at"`
}

type RegisterRequest struct {
	Username    string `json:"username" validate:"required,min=3,max=50,alphanum"`
	Email       string `json:"email" validate:"required,email,max=100"`
	Password    string `json:"password" validate:"required,min=8,max=100"`
	FullName    string `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=20"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User      *UserResponse `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt time.Time     `json:"expires_at"`
}

func ToUserResponse(user *entity.Users) *UserResponse {
	return &UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		FullName:      user.FullName,
		PhoneNumber:   user.PhoneNumber,
		EmailVerified: user.EmailVerifiedAt != nil,
		CreatedAt:     user.CreatedAt,
	}
}
