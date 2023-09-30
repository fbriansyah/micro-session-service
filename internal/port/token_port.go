package port

import (
	"time"

	dmtoken "github.com/fbriansyah/micro-session-service/internal/application/domain/token"
)

// Maker is an interface for managing tokens
type TokenMakerPort interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(mode string, userID, duration time.Duration) (string, *dmtoken.Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*dmtoken.Payload, error)
}
