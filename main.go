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
	// db.Get().DropTableIfExists(&models.Quiz{}, &models.Question{}, &models.Option{}, &models.User{}, &models.AdminUser{}, &models.Result{})
	// db.Get().CreateTable(&models.Quiz{}, &models.Question{}, &models.Option{}, &models.User{}, &models.AdminUser{}, &models.Result{})
	db.Get().DropTableIfExists(&models.Result{})
	db.Get().CreateTable(&models.Result{})
	db.Get().AutoMigrate(&models.Quiz{}, &models.Question{}, &models.Option{}, &models.User{}, &models.AdminUser{})

	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })
	router.GET("/quizzes", controllers.GetAllQuizzes)
	router.GET("/quizzes/:id", controllers.GetQuiz)
	router.GET("/top_users", controllers.GetTopUsers)
	router.GET("/top_users/:id", controllers.GetTopUsersOfQuiz)

	router.POST("/create_user", controllers.CreateUser)
	router.POST("/create_admin_user", controllers.CreateAdminUser)

	adminAuthorized := router.Group("/", controllers.AuthorizeAdmin())
	adminAuthorized.POST("/create_quiz", controllers.CreateQuiz)

	userAuthorized := router.Group("/", controllers.AuthorizeUser())
	userAuthorized.POST("/quiz", controllers.PostResults)

	router.Run(":" + port)
}
