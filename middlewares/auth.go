package middlewares

import (
	"book-manager/utils"
	"book-manager/utils/response"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Fail(c, http.StatusUnauthorized, "Missing or invalid token")			
			
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			response.Fail(c, http.StatusUnauthorized, "Invalid token")
			
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}
