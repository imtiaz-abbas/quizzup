package models

// AdminUser holds Admin data
type AdminUser struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	EmailID     string `json:"emailId"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}
