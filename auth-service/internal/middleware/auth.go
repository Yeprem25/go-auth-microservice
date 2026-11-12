package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type TokenParser interface {
	ParseAccessToken(token string) (int64, error)
}

func Auth(parser TokenParser) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")
		userID, err := parser.ParseAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Set("userId", userID)
		c.Next()
	}
}
