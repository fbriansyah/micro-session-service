package application

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	dmsession "github.com/fbriansyah/micro-session-service/internal/application/domain/session"
	dmtoken "github.com/fbriansyah/micro-session-service/internal/application/domain/token"
	"github.com/fbriansyah/micro-session-service/internal/port"
)

const (
	MODE_ACCESS_TOKEN    = "access"
	MODE_REFRESH_TOKEN   = "refresh"
	SESSION_CACHE_PREFIX = "SES"
)

type SessionService struct {
	port.SessionServicePort

	tokenMaker   port.TokenMakerPort
	cacheAdapter port.CacheAdapterPort
}

func generateSessionKey(id string) string {
	return fmt.Sprintf("%s_%s", SESSION_CACHE_PREFIX, id)
}

// CreateSession new session and save it to cache database.
func (s *SessionService) CreateSession(
	ctx context.Context, userID string, accessDuration, refreshDuration time.Duration) (dmsession.Session, error) {

	// create access token
	accessToken, accessPayload, err := s.tokenMaker.CreateToken(
		MODE_ACCESS_TOKEN,
		userID,
		accessDuration,
	)
	if err != nil {
		return dmsession.Session{}, err
	}

	// create refresh token
	refreshToken, refreshPayload, err := s.tokenMaker.CreateToken(
		MODE_REFRESH_TOKEN,
		userID,
		refreshDuration,
	)
	if err != nil {
		return dmsession.Session{}, err
	}

	// create session
	sessionData := dmsession.Session{
		ID:                    refreshPayload.ID.String(),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
	}

	// save session data to cache database
	sessionKey := generateSessionKey(refreshPayload.ID.String())

	sessionJson, err := json.Marshal(sessionData)
	if err != nil {
		return dmsession.Session{}, err
	}

	err = s.cacheAdapter.SetData(
		ctx,
		sessionKey,
		string(sessionJson),
		refreshDuration,
	)

	if err != nil {
		return dmsession.Session{}, err
	}

	return sessionData, nil
}

// GetPayloadFromToken validate the token and get payload data
func (s *SessionService) GetPayloadFromToken(ctx context.Context, token string) (dmtoken.Payload, error) {
	// verify token to get payload data
	payload, err := s.tokenMaker.VerifyToken(token)
	if err != nil {
		return dmtoken.Payload{}, err
	}

	return dmtoken.Payload{
		ID:        payload.ID,
		UserID:    payload.UserID,
		Mode:      payload.Mode,
		IssuedAt:  payload.IssuedAt,
		ExpiredAt: payload.ExpiredAt,
	}, nil
}

// RefreshToken generate new access token if token still valid
func (s *SessionService) RefreshToken(ctx context.Context, sessionID string, duration time.Duration) (dmsession.Session, error) {
	var session dmsession.Session
	sessionKey := generateSessionKey(sessionID)

	strData, err := s.cacheAdapter.GetData(ctx, sessionKey)

	if err != nil {
		return dmsession.Session{}, err
	}

	err = json.Unmarshal([]byte(strData), &session)
	if err != nil {
		return dmsession.Session{}, err
	}

	// create new access token
	accessToken, accessPayload, err := s.tokenMaker.CreateToken(
		MODE_ACCESS_TOKEN,
		session.UserID,
		duration,
	)
	if err != nil {
		return dmsession.Session{}, err
	}

	return dmsession.Session{
		ID:                    session.ID,
		UserID:                session.UserID,
		RefreshToken:          session.RefreshToken,
		AccessToken:           accessToken,
		RefreshTokenExpiresAt: session.RefreshTokenExpiresAt,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
	}, nil
}

// DeleteSession from cache database
func (s *SessionService) DeleteSession(ctx context.Context, sessionID string) (string, error) {
	sessionKey := generateSessionKey(sessionID)

	_, err := s.cacheAdapter.DeleteData(ctx, sessionKey)
	if err != nil {
		return "", err
	}

	return sessionID, nil
}
