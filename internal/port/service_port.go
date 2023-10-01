package port

import (
	"context"
	"time"

	dmsession "github.com/fbriansyah/micro-session-service/internal/application/domain/session"
	dmtoken "github.com/fbriansyah/micro-session-service/internal/application/domain/token"
)

type SessionServicePort interface {
	CreateSession(ctx context.Context, userID string, accessDuration, refreshDuration time.Duration) (dmsession.Session, error)
	RefreshToken(ctx context.Context, sessionID string, duration time.Duration) (dmsession.Session, error)
	GetPayloadFromToken(ctx context.Context, token string) (dmtoken.Payload, error)
	DeleteSession(ctx context.Context, sessionID string) (string, error)
}
