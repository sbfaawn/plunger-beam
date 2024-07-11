package http

import (
	"plunger-beam/api/dto"

	"github.com/gin-gonic/gin"
)

type basicHandler struct {
}

func NewBasicHandler() *basicHandler {
	return &basicHandler{}
}

func (h *basicHandler) NoRouteHandler(ctx *gin.Context) {
	ctx.JSON(404, dto.Response{
		Message: "",
		Data:    "",
		Error:   "404 Endpoint not found",
	})
}

func (h *basicHandler) NoMethodAllowed(ctx *gin.Context) {
	ctx.JSON(400, dto.Response{
		Message: "",
		Data:    "",
		Error:   "No Method Allowed",
	})
}

func (h *basicHandler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, dto.Response{
		Message: "",
		Data:    "",
		Error:   "Chat Message API is Up",
	})
}
