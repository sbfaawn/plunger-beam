package http

import (
	"plunger-beam/internal/dto"

	"github.com/gin-gonic/gin"
)

type BasicHandler struct {
}

func NewBasicHandler() *BasicHandler {
	return &BasicHandler{}
}

func (h *BasicHandler) NoRouteHandler(ctx *gin.Context) {
	ctx.JSON(404, dto.Response{
		Message: "",
		Data:    "",
		Error:   "404 Endpoint not found",
	})
}

func (h *BasicHandler) NoMethodAllowed(ctx *gin.Context) {
	ctx.JSON(400, dto.Response{
		Message: "",
		Data:    "",
		Error:   "No Method Allowed",
	})
}

func (h *BasicHandler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, dto.Response{
		Message: "",
		Data:    "",
		Error:   "Chat Message API is Up",
	})
}
