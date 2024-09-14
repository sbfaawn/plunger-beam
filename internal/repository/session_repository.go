package repository

import (
	"plunger-beam/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type SessionRepository interface {
	SaveSession(ctx *gin.Context, session *models.Session) error
	GetSessionById(ctx *gin.Context, sessionId string) (string, error)
	DeleteSession(ctx *gin.Context, sessionId string) error
	UpdateSessionExpiration(ctx *gin.Context, session *models.Session) error
}

type sessionRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewSessionRepository(DB *gorm.DB, redisClient *redis.Client) SessionRepository {
	return &sessionRepository{
		db:          DB,
		redisClient: redisClient,
	}
}

func (storage *sessionRepository) SaveSession(ctx *gin.Context, session *models.Session) error {
	var err error
	redisClient := storage.redisClient

	setCmd := redisClient.Set(ctx, session.Id, session.Username, session.TTL)
	if err = setCmd.Err(); err != nil {
		return err
	}

	return nil
}

func (storage *sessionRepository) GetSessionById(ctx *gin.Context, sessionId string) (string, error) {
	var err error
	redisClient := storage.redisClient

	getCmd := redisClient.Get(ctx, sessionId)
	if err = getCmd.Err(); err != nil {
		return "", err
	}

	res, err := getCmd.Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (storage *sessionRepository) DeleteSession(ctx *gin.Context, sessionId string) error {
	var err error
	redisClient := storage.redisClient

	delCmd := redisClient.Del(ctx, sessionId)
	if err = delCmd.Err(); err != nil {
		return err
	}

	return nil
}

func (storage *sessionRepository) UpdateSessionExpiration(ctx *gin.Context, session *models.Session) error {
	var err error
	redisClient := storage.redisClient

	setCmd := redisClient.Expire(ctx, session.Id, session.TTL)
	if err = setCmd.Err(); err != nil {
		return err
	}

	return nil
}
