package models

// Question struct
type Question struct {
	ID           string   `json:"id"`
	QuestionText string   `json:"question"`
	Options      []Option `json:"options"`
}
