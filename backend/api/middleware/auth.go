package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kimashii-dan/food-delivery-app/backend/pkg"
)

func CheckAuth(jwtService *pkg.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("accessToken")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "access token not found"})
			c.Abort()
			return
		}

		claims, err := jwtService.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}
