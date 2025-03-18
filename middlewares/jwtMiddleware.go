package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalpesh-vala/go-chat-app/utils"
)

// JWTAuthMiddleware validates JWT tokens in incoming requests
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString) // Using function from utils/jwt.go
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()
	}
}
