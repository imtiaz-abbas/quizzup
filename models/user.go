package models

import (
	"github.com/jinzhu/gorm"
)

//User Information
type User struct {
	gorm.Model
	Name     string   `json:"name"`
	EmailID  string   `json:"email_id" gorm:"unique;not null; unique_index"`
	Password string   `json:"password"`
	Results  []Result `json:"results"`
}
