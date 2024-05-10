package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func DummyMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	// Foo()
	return func(c *gin.Context) {
		fmt.Println("I'm a dummy middleware")
		c.Next()
	}
}
