package models

import (
	"fmt"

	"github.com/quizzup/db"
)

// Quiz Information
type Quiz struct {
	ID      uint64     `json:"id" gorm:"primary_key"`
	Content []Question `json:"content"`
}

func getQuizzes() []Quiz {
	var allQuizzes []Quiz
	if err := db.Get().Find(&allQuizzes).Error; err != nil {
		fmt.Println("==== error ====", err)
		return []Quiz{}
	}
	fmt.Println("===== found users ")
	fmt.Println("===== found users ", allQuizzes)
	fmt.Println("===== found users ")
	return allQuizzes

}

// GetQuiz gets a Quiz by id
func (quiz *Quiz) GetQuiz(id uint64) int {
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
