package service

import (
	"errors"
	"plunger-beam/internal/models"
	"plunger-beam/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SESSION_TTL = 60 * time.Minute
)

type SessionService interface {
	CreateSession(ctx *gin.Context, username string) (models.Session, error)
	CheckSession(ctx *gin.Context, sessionId string) (string, error)
	RemoveSession(ctx *gin.Context, sessionId string) error
	UpdateSessionExpiration(ctx *gin.Context, sessionId string) (models.Session, error)
}

type sessionServiceImpl struct {
	sessionRepository repository.SessionRepository
}

func NewSessionService(sessionRepository repository.SessionRepository) SessionService {
	return &sessionServiceImpl{
		sessionRepository: sessionRepository,
	}
}

func (s *sessionServiceImpl) CreateSession(ctx *gin.Context, username string) (models.Session, error) {
	sessionToken := uuid

	session := models.Session{
		Id:       sessionToken,
		Username: username,
		TTL:      SESSION_TTL,
	}

	err := s.sessionRepository.SaveSession(ctx, &session)

	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func (s *sessionServiceImpl) CheckSession(ctx *gin.Context, sessionId string) (string, error) {
	username, err := s.sessionRepository.GetSessionById(ctx, sessionId)

	if err != nil {
		return "", errors.New("Session with SessionId " + sessionId + " is not found")
	}

	return username, nil
}

func (s *sessionServiceImpl) RemoveSession(ctx *gin.Context, sessionId string) error {
	err := s.sessionRepository.DeleteSession(ctx, sessionId)

	if err != nil {
		return err
	}

	return nil
}

func (s *sessionServiceImpl) UpdateSessionExpiration(ctx *gin.Context, sessionId string) (models.Session, error) {
	username, err := s.sessionRepository.GetSessionById(ctx, sessionId)

	if err != nil {
		return models.Session{}, errors.New("Session with SessionId " + sessionId + " is not found")
	}

	session := models.Session{
		Username: username,
		TTL:      SESSION_TTL,
	}

	err = s.sessionRepository.UpdateSessionExpiration(ctx, &session)

	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}
