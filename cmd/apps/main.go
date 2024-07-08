package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	fmt.Println("hello")

	group := server.Group("/api/v1")
	group.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"status":  "Success",
			"data":    "",
			"message": "",
		})
	})

	server.Run(":8026")
}
