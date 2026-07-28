package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireGateway() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("X-Forwarded-By") != "api-gateway" {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": "Direct access is not allowed. Use API Gateway.",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
