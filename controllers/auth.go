package controllers

import (
	"book-manager/config"
	"book-manager/models"
	"book-manager/utils"
	"book-manager/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Fail(c,http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		response.Fail(c,http.StatusBadRequest, "Username already exists")
		return
	}
	response.Success(c,user)

}

func Login(c *gin.Context) {
	var user models.User
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Where("username = ? AND password = ?", input.Username, input.Password).First(&user).Error; err != nil {
		response.Fail(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	token, _ := utils.GenerateToken(user.Username)
	response.Success(c,token)
}
