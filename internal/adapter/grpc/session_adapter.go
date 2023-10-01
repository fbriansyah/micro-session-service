package grpc

import (
	"context"
	"fmt"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/session"
	"github.com/fbriansyah/micro-session-service/internal/application"
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
func (a *GrpcServerAdapter) RefreshToken(ctx context.Context, sess *session.SessionID) (*session.Session, error) {
	s, err := a.service.RefreshToken(ctx, sess.SessionId, a.refeshTokenDuration)
	if err != nil {
		return nil, generateError(
			codes.FailedPrecondition,
			fmt.Sprintf("error refresh token: %v", err),
		)
	}

	return &session.Session{
		Id:                    s.ID,
		UserId:                s.UserID,
		AccessToken:           s.AccessToken,
		RefreshToken:          s.RefreshToken,
		AccessTokenExpiresAt:  util.ToDateTime(s.AccessTokenExpiresAt),
		RefreshTokenExpiresAt: util.ToDateTime(s.RefreshTokenExpiresAt),
	}, nil
}
func (a *GrpcServerAdapter) DeleteSession(ctx context.Context, sess *session.SessionID) (*session.SessionID, error) {

	_, err := a.service.DeleteSession(ctx, sess.SessionId)
	if err != nil {
		return nil, generateError(
			codes.FailedPrecondition,
			fmt.Sprintf("error delete session: %v", err),
		)
	}

	return &session.SessionID{
		SessionId: sess.SessionId,
	}, nil
}
func (a *GrpcServerAdapter) GetPayloadFromToken(ctx context.Context, tkn *session.Token) (*session.Payload, error) {
	payload, err := a.service.GetPayloadFromToken(ctx, tkn.AccessToken)
	if err != nil {
		return nil, generateError(
			codes.FailedPrecondition,
			fmt.Sprintf("error get payload from token: %v", err),
		)
	}

	if payload.Mode == application.MODE_REFRESH_TOKEN {
		return nil, generateError(
			codes.InvalidArgument,
			"cannot get payload using reffresh token",
		)
	}

	return &session.Payload{
		UserId: payload.UserID,
	}, nil
}
