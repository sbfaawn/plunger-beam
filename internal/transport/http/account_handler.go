package http

import (
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{}
}

func (h *AccountHandler) RegistrationHandler(ctx *gin.Context) {

}

func (h *AccountHandler) LoginHandler(ctx *gin.Context) {

}

func (h *AccountHandler) LogoutHandler(ctx *gin.Context) {

}

func (h *AccountHandler) RefreshTokenHandler(ctx *gin.Context) {

}

func (h *AccountHandler) CheckSession(ctx *gin.Context) {

}
