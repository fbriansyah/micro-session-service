package port

import (
	"context"
	"time"

	dmsession "github.com/fbriansyah/micro-session-service/internal/application/domain/session"
	dmtoken "github.com/fbriansyah/micro-session-service/internal/application/domain/token"
)

type SessionServicePort interface {
	// CreateSession new session and save it to cache database.
	CreateSession(ctx context.Context, userID string, accessDuration, refreshDuration time.Duration) (dmsession.Session, error)
	// RefreshToken generate new access token if token still valid
	RefreshToken(ctx context.Context, sessionID string, duration time.Duration) (dmsession.Session, error)
	// GetPayloadFromToken validate the token and get payload data
	GetPayloadFromToken(ctx context.Context, token string) (dmtoken.Payload, error)
	// DeleteSession from cache database
	DeleteSession(ctx context.Context, sessionID string) (string, error)
}
