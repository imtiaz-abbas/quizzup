package models

// AdminUser holds Admin data
type AdminUser struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	EmailID     string `json:"email_id"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
