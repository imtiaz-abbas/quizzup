package controllers

import (
	"net/http"

	"github.com/quizzup/db"
	"github.com/quizzup/models"

	"github.com/gin-gonic/gin"
)

// GetUsers gets all the user from db
func GetUsers() []models.User {
	var users []models.User
	if err := db.Get().Find(&users).Error; err != nil {
		return []models.User{}
	}
	return users
}

// CreateUser method Creates a User
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	if err := db.Get().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

// AuthorizeUser d
func AuthorizeUser() gin.HandlerFunc {
	accounts := make(map[string]string)
	var users []models.User
	if err := db.Get().Find(&users).Error; err != nil {
		return gin.BasicAuth(gin.Accounts{"test": "test"})
	}
	for _, user := range users {
		accounts[user.EmailID] = user.Password
	}
	if len(accounts) == 0 {
		accounts["test"] = "test"
	}
	return gin.BasicAuth(accounts)
}
