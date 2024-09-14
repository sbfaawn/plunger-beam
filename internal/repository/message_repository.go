package repository

import (
	"errors"
	"plunger-beam/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MessageRepository interface {
	SaveMessage(ctx *gin.Context, message *models.Message) error
	GetMessagesByConversation(ctx *gin.Context, usernameOne string, usernameTwo string) ([]models.Message, error)
	DeleteMessageById(ctx *gin.Context, messageId int) error
}

type messageRepositoryImpl struct {
	db *gorm.DB
}

func NewMessageRepository(DB *gorm.DB) MessageRepository {
	return &messageRepositoryImpl{
		db: DB,
	}
}

func (storage *messageRepositoryImpl) SaveMessage(ctx *gin.Context, message *models.Message) error {
	var err error
	db := storage.db

	tx := db.Begin()
	err = tx.Create(&message).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (storage *messageRepositoryImpl) GetMessagesByConversation(ctx *gin.Context, usernameOne string, usernameTwo string) ([]models.Message, error) {
	db := storage.db
	var err error
	var messages []models.Message

	tx := db.Begin()
	err = tx.Select("message_id", "sender", "receiver", "message", "created_at").
		Where("(sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?)", usernameOne, usernameTwo, usernameTwo, usernameOne).
		Find(&messages).
		Error

	if err != nil {
		tx.Rollback()
		return messages, err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return messages, err
	}

	return messages, nil
}

func (storage *messageRepositoryImpl) DeleteMessageById(ctx *gin.Context, messageId int) error {
	var err error
	db := storage.db
	message := models.Message{
		Id: messageId,
	}

	tx := db.Begin()
	delete := tx.Model(&message).Delete(&message)

	if delete.Error != nil {
		tx.Rollback()
		return err
	}

	result := delete.Commit().WithContext(ctx)

	if result.Error != nil {
		return err
	}

	if delete.RowsAffected == result.RowsAffected {
		tx.Rollback()
		return errors.New("record with messageId is not found")
	}

	return nil
}
