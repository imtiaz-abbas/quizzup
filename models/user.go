package models

//User Information
type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	EmailID     string `json:"email_id"`
	Password    string `json:"password"`
}
