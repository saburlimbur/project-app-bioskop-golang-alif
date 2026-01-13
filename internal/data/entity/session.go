package entity

import (
	"time"
)

// type Session struct {
// 	ID        int
// 	UserID    int
// 	Token     string
// 	ExpiresAt time.Time
// }

type Session struct {
	ID         int        `db:"id"`
	UserID     int        `db:"user_id"`
	Token      string     `db:"token"`
	ExpiresAt  time.Time  `db:"expires_at"`
	RevokedAt  *time.Time `db:"revoked_at"`
	IPAddress  string     `db:"ip_address"`
	DeviceInfo string     `db:"device_info"`
	CreatedAt  time.Time  `db:"created_at"`
}

func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}

func (s *Session) IsRevoked() bool {
	return s.RevokedAt != nil
}

func (s *Session) IsValid() bool {
	return !s.IsExpired() && !s.IsRevoked()
}
