package http

import (
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) SendMessages(ctx *gin.Context) {

}

func (h *MessageHandler) GetConversation(ctx *gin.Context) {

}

func (h *MessageHandler) DeleteMessage(ctx *gin.Context) {

}
