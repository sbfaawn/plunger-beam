package http

import (
	"plunger-beam/internal/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) NoRouteHandler(ctx *gin.Context) {
	ctx.JSON(404, dto.Response{
		Message: "",
		Data:    "",
		Error:   "404 Endpoint not found",
	})
}

func (h *Handler) NoMethodAllowed(ctx *gin.Context) {
	ctx.JSON(400, dto.Response{
		Message: "",
		Data:    "",
		Error:   "No Method Allowed",
	})
}

func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, dto.Response{
		Message: "",
		Data:    "",
		Error:   "Chat Message API is Up",
	})
}
