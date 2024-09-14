package service

import (
	"plunger-beam/internal/models"
	"plunger-beam/internal/repository"

	"github.com/gin-gonic/gin"
)

type MessageService interface {
	SendMessage(ctx *gin.Context, message *models.Message) error
	GetConversation(ctx *gin.Context, usernameOne string, usernameTwo string) ([]models.Message, error)
	DeleteMessage(ctx *gin.Context, messageId int) error
}

type messageServiceImpl struct {
	messageRepository repository.MessageRepository
}

func NewMessageService(messageRepository repository.MessageRepository) MessageService {
	return &messageServiceImpl{
		messageRepository: messageRepository,
	}
}

func (service *messageServiceImpl) SendMessage(ctx *gin.Context, message *models.Message) error {
	err := service.messageRepository.SaveMessage(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

func (service *messageServiceImpl) GetConversation(ctx *gin.Context, usernameOne string, usernameTwo string) ([]models.Message, error) {
	messages, err := service.messageRepository.GetMessagesByConversation(ctx, usernameOne, usernameTwo)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (service *messageServiceImpl) DeleteMessage(ctx *gin.Context, messageId int) error {
	err := service.messageRepository.DeleteMessageById(ctx, messageId)
	if err != nil {
		return err
	}

	return nil
}
