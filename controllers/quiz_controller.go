package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quizzup/models"
)

// GetAllQuizzes gives all Products from Model
func GetAllQuizzes(c *gin.Context) {
	err, allQuizzes := models.GetAllQuizzes()
	if err == 0 {
		c.JSON(404, gin.H{"error": "Error Finding Products"})
	} else {
		c.JSON(200, allQuizzes)
	}
}

// GetQuiz gets single product by id
func GetQuiz(c *gin.Context) {
	id := c.Param("id")
	quiz := models.Quiz{}
	statusCode := quiz.GetQuiz(id)
	if statusCode == 1 {
		c.JSON(200, quiz)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Product with ID found"})
	}
}
