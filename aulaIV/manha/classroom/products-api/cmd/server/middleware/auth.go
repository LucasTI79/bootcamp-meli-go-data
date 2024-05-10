package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/batatinha123/products-api/pkg/web"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("api_token")

		if token == "" {
			web.RespondWithError(c, http.StatusUnauthorized, "api token required")
			return
		}

		if token != requiredToken {
			web.RespondWithError(c, http.StatusUnauthorized, "invalid API token")
			return
		}

		c.Next()
	}
}
