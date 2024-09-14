package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		ctx.Next()

		// Stop timer
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Log format
		log.Printf("| %3d | %13v | %15s | %-7s  %#v\n",
			ctx.Writer.Status(),
			latency,
			ctx.ClientIP(),
			ctx.Request.Method,
			ctx.Request.URL.Path,
		)
	}
}
