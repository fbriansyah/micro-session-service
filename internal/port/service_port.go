package port

import (
	dmsession "github.com/fbriansyah/micro-session-service/internal/application/domain/session"
	dmtoken "github.com/fbriansyah/micro-session-service/internal/application/domain/token"
)

type SessionServicePort interface {
	CreateSession(userID string) (dmsession.Session, error)
	RefreshToken(sessionID string) (dmsession.Session, error)
	GetPayloadFromToken(token string) (dmtoken.Payload, error)
}
