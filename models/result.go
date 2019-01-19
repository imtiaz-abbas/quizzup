package models

import (
	"github.com/jinzhu/gorm"
)

// Result type to store all the results
type Result struct {
	gorm.Model
	UserID uint `json:"user_id" gorm:"foreign_key"`
	QuizID uint `json:"quiz_id" gorm:"foreign_key"`
	Score  int  `json:"score"`
}
