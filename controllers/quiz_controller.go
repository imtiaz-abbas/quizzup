package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quizzup/db"
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
	i, _ := strconv.ParseUint(id, 10, 8)
	quiz := models.Quiz{}
	statusCode := quiz.GetQuiz(uint(i))
	if statusCode == 1 {
		c.JSON(200, quiz)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Product with ID found"})
	}
}

// CreateQuizRequestBody s
type CreateQuizRequestBody struct {
	Name    string         `json:"name"`
	Content []QuestionBody `json:"content"`
}

// QuestionBody asd
type QuestionBody struct {
	QuestionText  string       `json:"question"`
	Options       []OptionBody `json:"options"`
	CorrectOption uint         `json:"correct_option"`
}

//OptionBody adf
type OptionBody struct {
	Option string `json:"option"`
}

// CreateQuiz creates quiz
func CreateQuiz(c *gin.Context) {
	q := CreateQuizRequestBody{}
	data := &models.Quiz{}
	questions := []models.Question{}

	err := c.Bind(&q)
	for _, question := range q.Content {
		options := []models.Option{}
		for _, option := range question.Options {
			op := models.Option{}
			op.OptionText = option.Option
			options = append(options, op)
		}
		questions = append(questions, models.Question{QuestionText: question.QuestionText, Options: options, CorrectOption: question.CorrectOption})
	}
	data.Name = q.Name
	data.Questions = questions
	if err := db.Get().Model(&models.Quiz{}).Create(data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
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
