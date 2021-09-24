package auth

import (
	"time"
)

type Login struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}
