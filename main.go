package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/quizzup/controllers"
)

func dbURL() string {
	var sslmode string
	sslmode = "sslmode=disable"
	dbPort, _ := strconv.Atoi(os.Getenv("DBPORT"))
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s %s",
		os.Getenv("DBHOST"),
		dbPort,
		os.Getenv("DBUSER"),
		os.Getenv("DBNAME"),
		os.Getenv("DBPASSWORD"),
		sslmode,
	)
}

func main() {
	port := os.Getenv("PORT")
	db, err := gorm.Open("postgres", dbURL())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	dbase := db.DB()
	defer dbase.Close()
	err = dbase.Ping()
	if err != nil {
		panic(err.Error())
	}
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })
	router.GET("/quizzes", controllers.GetAllQuizzes)
	router.GET("/quizzes/:id", controllers.GetQuiz)

	adminAuthorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	adminAuthorized.POST("/create_quiz", controllers.CreateQuiz)

	router.Run(":" + port)
}
