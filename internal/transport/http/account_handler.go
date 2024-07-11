package http

import (
	"github.com/gin-gonic/gin"
)

type accountHandler struct {
}

func NewAccountHandler() *accountHandler {
	return &accountHandler{}
}

func (h *accountHandler) RegistrationHandler(ctx *gin.Context) {

}

func (h *accountHandler) LoginHandler(ctx *gin.Context) {

}

func (h *accountHandler) LogoutHandler(ctx *gin.Context) {

}

func (h *accountHandler) RefreshTokenHandler(ctx *gin.Context) {

}

func (h *accountHandler) CheckSession(ctx *gin.Context) {

}
