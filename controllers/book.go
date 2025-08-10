package controllers

import (
	"book-manager/config"
	"book-manager/models"
	"book-manager/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, "Failed to fetch books")
		return
	}
	response.Success(c, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		response.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	book.Available = true
	if err := config.DB.Create(&book).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, "Failed to create book")
		return
	}
	response.Success(c, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		response.Fail(c, http.StatusNotFound, "Book not found")
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		response.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Save(&book).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, "Failed to update book")
		return
	}

	response.Success(c, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Fail(c, http.StatusNotFound, "Book not found")
		} else {
			response.Fail(c, http.StatusInternalServerError, "Database error")
		}
		return
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, "Failed to delete book")
		return
	}

	response.Success(c, nil)
}

func BorrowBook(c *gin.Context) {
	id := c.Param("id")
	username, exists := c.Get("username")
	if !exists {
		response.Fail(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Fail(c, http.StatusNotFound, "Book not found")
		} else {
			response.Fail(c, http.StatusInternalServerError, "Database error")
		}
		return
	}

	if !book.Available {
		response.Fail(c, http.StatusBadRequest, "Book not available")
		return
	}

	book.Available = false
	book.Borrower = username.(string)

	if err := config.DB.Save(&book).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, "Failed to borrow book")
		return
	}

	response.Success(c, book)
}

func ReturnBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Fail(c, http.StatusNotFound, "Book not found")
		} else {
			response.Fail(c, http.StatusInternalServerError, "Database error")
		}
		return
	}

	if book.Available {
		response.Fail(c, http.StatusBadRequest, "Book is not borrowed")
		return
	}

	book.Available = true
	book.Borrower = ""

	if err := config.DB.Save(&book).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, "Failed to return book")
		return
	}

	response.Success(c, book)
}
