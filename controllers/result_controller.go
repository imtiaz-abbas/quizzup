package controllers

import (
	"fmt"
	"net/http"

	"github.com/quizzup/db"
	"github.com/quizzup/models"

	"github.com/gin-gonic/gin"
)

// PostQuizRequestBody sfd
type PostQuizRequestBody struct {
	QuizID  uint     `json:"quiz_id"`
	Answers []Answer `json:"answers"`
}

// Answer s
type Answer struct {
	QuestionID uint `json:"question_id"`
	Option     uint `json:"option_id"`
}

// PostResults s
func PostResults(c *gin.Context) {
	username, _, _ := c.Request.BasicAuth()
	req := PostQuizRequestBody{}
	// var users []models.User
	fmt.Println(" ==== username ", username)
	var user models.User
	c.Bind(&req)

	if error := db.Get().Preload("Results").Where("email_id = ?", username).Find(&user).Error; error != nil {
		fmt.Println(" ==== unable to find user with username ", username)
		c.JSON(http.StatusOK, gin.H{
			"message": "NOT OK",
		})
		return
	}

	fmt.Println(" ==== user while posting results", user.EmailID)

	if len(user.Results) == 0 {
		createResult(c, req, user)
	}
	for _, result := range user.Results {
		if result.QuizID == req.QuizID {
			fmt.Println(" ==== user can attend a quiz only once ===")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "You can only submit a quiz once",
			})
			return
		}
		createResult(c, req, user)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

//createResult s
func createResult(c *gin.Context, req PostQuizRequestBody, user models.User) {
	finalResult := &models.Result{}
	finalResult.QuizID = req.QuizID
	score := 0
	for _, ans := range req.Answers {
		var question models.Question
		if error := db.Get().Where("id = ?", ans.QuestionID).Find(&question).Error; error != nil {
			fmt.Println("=== error ", error)
		}
		if question.CorrectOption == ans.Option {
			score += 10
		} else {
			score -= 5
		}
	}
	finalResult.UserID = user.ID
	finalResult.Score = score
	if err := db.Get().Create(&finalResult).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
}
