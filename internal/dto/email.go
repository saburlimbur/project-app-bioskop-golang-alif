package dto

import "time"

type VerifyEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp" validate:"required,len=6"`
}

type VerifyEmailResponse struct {
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	VerifiedAt    time.Time `json:"verified_at"`
}
