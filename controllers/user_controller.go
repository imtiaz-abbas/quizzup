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

// UserInput struct
type UserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	EmailID  string `json:"email_id"`
}

// CreateUser method Creates a User
func CreateUser(c *gin.Context) {
	user := &models.User{}
	var userInput UserInput
	if err := c.Bind(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	user.Name = userInput.Name
	user.EmailID = userInput.EmailID
	user.Password = userInput.Password

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
		accounts["one"] = "test"
		accounts["two"] = "test"
		accounts["three"] = "test"
	}
	return gin.BasicAuth(accounts)
}
