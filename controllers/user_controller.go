package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

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
func AuthorizeUser(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.ParseWithClaims(strings.Split(tokenString, " ")[1], &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("QUIZZUP_SECRET_KEY")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token is invalid"})
		return
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		fmt.Println(claims.UserID)
		c.Set("userId", claims.UserID)
		c.Next()
		return
	}
	fmt.Println(err)
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token is invalid"})
	return
}

// LoginInput type
type LoginInput struct {
	EmailID  string `json:"email_id"`
	Password string `json:"password"`
}

// MyClaims s
type MyClaims struct {
	UserID uint
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{user.ID, jwt.StandardClaims{}})
	signedToken, _ := token.SignedString([]byte(os.Getenv("QUIZZUP_SECRET_KEY")))

	c.JSON(200, gin.H{"token": signedToken})
}
