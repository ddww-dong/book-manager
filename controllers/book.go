package controllers

import (
	"book-manager/config"
	"book-manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.Available = true
	config.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.ShouldBindJSON(&book)
	config.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Book{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func BorrowBook(c *gin.Context) {
	id := c.Param("id")
	username := c.MustGet("username").(string)
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil || !book.Available {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not available"})
		return
	}
	book.Available = false
	book.Borrower = username
	config.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func ReturnBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil || book.Available {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not borrowed"})
		return
	}
	book.Available = true
	book.Borrower = ""
	config.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}
