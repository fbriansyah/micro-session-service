package application

import (
	"encoding/json"
	"errors"
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

func (s *SessionService) CreateSession(
	userID string, accessDuration, refreshDuration time.Duration) (dmsession.Session, error) {

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
		sessionKey,
		string(sessionJson),
		refreshDuration,
	)

	if err != nil {
		return dmsession.Session{}, err
	}

	return sessionData, nil
}
func (s *SessionService) GetPayloadFromToken(token string) (dmtoken.Payload, error) {
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
func (s *SessionService) RefreshToken(sessionID string) (dmsession.Session, error) {
	return dmsession.Session{}, errors.New("refresh token not yet implemented")
}
