package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(ctx *gin.Context) {
	fmt.Printf("[%s] %s\n", ctx.Request.Method, ctx.Request.URL)
	ctx.Next()
}
