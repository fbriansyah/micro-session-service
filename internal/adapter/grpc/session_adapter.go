package grpc

import (
	"context"
	"fmt"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/session"
	"github.com/fbriansyah/micro-session-service/util"
	"google.golang.org/grpc/codes"
)

func (a *GrpcServerAdapter) CreateSession(ctx context.Context, UserID *session.UserID) (*session.Session, error) {
	sessionDomain, err := a.service.CreateSession(ctx, UserID.UserId, a.accessTokenDuration, a.refeshTokenDuration)
	if err != nil {
		return nil, generateError(
			codes.FailedPrecondition,
			fmt.Sprintf("error creating session %v", err),
		)
	}

	return &session.Session{
		Id:                    sessionDomain.ID,
		UserId:                sessionDomain.UserID,
		AccessToken:           sessionDomain.AccessToken,
		RefreshToken:          sessionDomain.RefreshToken,
		AccessTokenExpiresAt:  util.ToDateTime(sessionDomain.AccessTokenExpiresAt),
		RefreshTokenExpiresAt: util.ToDateTime(sessionDomain.RefreshTokenExpiresAt),
	}, nil
}
func (a *GrpcServerAdapter) RefreshToken(context.Context, *session.SessionID) (*session.Session, error) {
	return &session.Session{}, nil
}
func (a *GrpcServerAdapter) DeleteSession(context.Context, *session.SessionID) (*session.SessionID, error) {
	return &session.SessionID{}, nil
}
func (a *GrpcServerAdapter) GetPayloadFromToken(context.Context, *session.Token) (*session.Payload, error) {
	return &session.Payload{}, nil
}
