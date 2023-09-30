package dmsession

import "time"

type Session struct {
	ID                    string    `json:"ID"`
	UserID                string    `json:"user_id"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}
