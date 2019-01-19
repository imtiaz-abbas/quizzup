package models

// Question struct
type Question struct {
	ID           uint64   `json:"id" gorm:"primary_key"`
	QuestionText string   `json:"question"`
	Options      []Option `json:"options"`
}
