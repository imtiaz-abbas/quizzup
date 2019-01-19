package models

// Question struct
type Question struct {
	ID           uint     `json:"id"`
	QuizID       uint     `json:"quiz_id"`
	QuestionText string   `json:"question"`
	Options      []Option `json:"options"`
}
