package auth

import (
	"time"
)

type Schema struct {
	ID           int
	Email        string
	OTP          string
	OTPExpiresAt time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
