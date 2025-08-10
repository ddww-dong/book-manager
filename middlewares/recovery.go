package middlewares

import (
	"book-manager/utils/response"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		defer func() {
			if err := recover(); err!=nil{
				log.Printf("panic recovered: %v", err)

				response.Fail(c, http.StatusInternalServerError, "Server Internal Error")

				c.Abort()
			}
		}()
		c.Next()
	}
}