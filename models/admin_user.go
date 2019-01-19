package models

import (
	"github.com/jinzhu/gorm"
)

// AdminUser holds Admin data
type AdminUser struct {
	gorm.Model
	Name        string `json:"name"`
	EmailID     string `json:"email_id" gorm:"primary_key;unique;unique_index"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
