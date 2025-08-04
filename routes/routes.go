package routes

import (
	"book-manager/controllers"
	"book-manager/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/api")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		auth.GET("/books", controllers.GetBooks)
		auth.POST("/books", controllers.CreateBook)
		auth.PUT("/books/:id", controllers.UpdateBook)
		auth.DELETE("/books/:id", controllers.DeleteBook)
		auth.POST("/books/:id/borrow", controllers.BorrowBook)
		auth.POST("/books/:id/return", controllers.ReturnBook)
	}

	return r
}
