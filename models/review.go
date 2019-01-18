package models

import (
	"time"
)

// Review Information
type Review struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Images    []string  `json:"images" gorm:"type:text[]"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"createdAt"`
	Rating    int       `json:"rating"`
	UserID    string    `json:"userId"`
	BlogID    string    `json:"blogId"`
}
