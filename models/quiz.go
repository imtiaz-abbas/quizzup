package models

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/quizzup/db"
)

// Quiz Information
type Quiz struct {
	gorm.Model
	Name      string     `json:"name"`
	Questions []Question `json:"content"`
}

func getQuizzes() []Quiz {
	var allQuizzes []Quiz
	if err := db.Get().Preload("Questions.Options").Find(&allQuizzes).Error; err != nil {
		fmt.Println("==== error ====", err)
		return []Quiz{}
	}
	return allQuizzes

}

// GetQuiz gets a Quiz by id
func (quiz *Quiz) GetQuiz(id uint) int {
	allQuizzes := getQuizzes()
	statusCode := 0
	for _, item := range allQuizzes {
		if id == item.ID {
			statusCode = 1
			*quiz = item
		}
	}
	return statusCode
}

// GetAllQuizzes func gets all the products in the database
func GetAllQuizzes() (int, []Quiz) {
	allQuizzes := getQuizzes()
	return 1, allQuizzes
}
