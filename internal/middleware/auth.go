package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakewd7/internal/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		authHeader := c.GetHeader("Authorization")
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}

		if tokenString == "" {
			tokenString, _ = c.Cookie("access_token")
		}

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No token provided"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)

		if claims.RoleID != nil {
			c.Set("roleID", *claims.RoleID)
		} else {
			c.Set("roleID", "")
		}

		c.Next()
	}
}
