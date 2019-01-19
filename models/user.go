package models

//User Information
type User struct {
	Name     string `json:"name"`
	EmailID  string `json:"email_id" gorm:"primary_key;unique;not null; unique_index"`
	Password string `json:"password"`
}
