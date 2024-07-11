package http

import (
	"github.com/gin-gonic/gin"
)

type messageHandler struct {
}

func NewMessageHandler() *messageHandler {
	return &messageHandler{}
}

func (h *messageHandler) SendMessages(ctx *gin.Context) {

}

func (h *messageHandler) GetConversation(ctx *gin.Context) {

}

func (h *messageHandler) DeleteMessage(ctx *gin.Context) {

}
