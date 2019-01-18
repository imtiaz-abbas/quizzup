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

// CreateQuizRequestBody s
type CreateQuizRequestBody struct {
	Content []QuestionBody `json:"content"`
}

// QuestionBody asd
type QuestionBody struct {
	QuestionText string       `json:"question"`
	Options      []OptionBody `json:"options"`
}

//OptionBody adf
type OptionBody struct {
	Option string `json:"option"`
}

// CreateQuiz creates quiz
func CreateQuiz(c *gin.Context) {
	q := CreateQuizRequestBody{}
	err := c.Bind(&q)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Invalid request",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Quiz submitted successfully",
		})
	}
}
