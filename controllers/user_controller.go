package controllers

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
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

// LoginInput type
type LoginInput struct {
	EmailID  string `json:"email_id"`
	Password string `json:"password"`
}

// MyClaims s
type MyClaims struct {
	User models.User
	jwt.StandardClaims
}

// LoginUser s
func LoginUser(c *gin.Context) {
	userInput := LoginInput{}
	c.Bind(&userInput)

	var user models.User
	if recordNotFound := db.Get().Where("email_id = ? AND password = ?", userInput.EmailID, userInput.Password).Find(&user).RecordNotFound(); recordNotFound {
		c.JSON(401, gin.H{"error": "Invalid Details"})
		return
	}

	fmt.Println(" ==== user ", user.EmailID, user.Password)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{user, jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test"}})
	signedToken, _ := token.SignedString([]byte("secret"))

	c.JSON(200, gin.H{"token": signedToken})
}
