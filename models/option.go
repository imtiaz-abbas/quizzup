package models

// Option struct
type Option struct {
	ID         uint   `json:"id"`
	OptionText string `json:"option"`
	QuestionID uint   `json:"question_id"`
}
