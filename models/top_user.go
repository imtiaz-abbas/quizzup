package models

// TopUser struct
type TopUser struct {
	User  User `json:"user"`
	Score int  `json:"score"`
}
