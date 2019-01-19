package controllers

import (
	"fmt"
	"net/http"

	"github.com/quizzup/db"
	"github.com/quizzup/models"

	"github.com/gin-gonic/gin"
)

// GetAdminUsers gets all the admin user from db
func GetAdminUsers() []models.AdminUser {
	var users []models.AdminUser
	if err := db.Get().Find(&users).Error; err != nil {
		return []models.AdminUser{}
	}
	return users
}

// CreateAdminUser method Creates a User
func CreateAdminUser(c *gin.Context) {
	var user models.AdminUser
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

// AuthorizeAdmin d
func AuthorizeAdmin() gin.HandlerFunc {
	accounts := make(map[string]string)
	var users []models.AdminUser
	if err := db.Get().Find(&users).Error; err != nil {
		return gin.BasicAuth(gin.Accounts{"admin": "admin"})
	}
	for _, user := range users {
		accounts[user.EmailID] = user.Password
	}
	fmt.Println(" ==== ", users)
	fmt.Println(" ==== ", accounts)
	if len(accounts) == 0 {
		accounts["admin"] = "admin"
	}
	return gin.BasicAuth(accounts)
}
