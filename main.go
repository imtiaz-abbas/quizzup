package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
	"github.com/quizzup/controllers"
	"github.com/quizzup/db"
	"github.com/quizzup/models"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	db.Get().DropTableIfExists(&models.Quiz{}, &models.Question{}, &models.Option{}, &models.User{})
	db.Get().CreateTable(&models.Quiz{}, &models.Question{}, &models.Option{}, &models.User{})

	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })
	router.GET("/quizzes", controllers.GetAllQuizzes)
	router.GET("/quizzes/:id", controllers.GetQuiz)

	router.POST("/create_user", controllers.CreateUser)

	adminAuthorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	adminAuthorized.POST("/create_quiz", controllers.CreateQuiz)

	router.Run(":" + port)
}
