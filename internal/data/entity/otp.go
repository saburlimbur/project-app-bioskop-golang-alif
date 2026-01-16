package entity

import "time"

type OTPVerification struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	OTPCode   string    `db:"otp_code"`
	ExpiresAt time.Time `db:"expires_at"`
	IsUsed    bool      `db:"is_used"`
	CreatedAt time.Time `db:"created_at"`
}
