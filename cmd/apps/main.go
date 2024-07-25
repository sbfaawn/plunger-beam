package main

import (
	"plunger-beam/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	cfg := config.Read()

	group := server.Group("/api/v1")
	group.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"status":  "Success",
			"data":    "",
			"message": "",
		})
	})

	server.Run(":" + cfg.Server.Port)
}
