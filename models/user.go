package models

import "time"

//User Information
type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Location     string    `json:"location"`
	PhoneNumber  string    `json:"phoneNumber"`
	EmailID      string    `json:"emailId"`
	Image        string    `json:"image"`
	DateOfBirth  time.Time `json:"dateOfBirth"`
	AadharNumber string    `json:"aadharNumber"`
	OtherID      string    `json:"otherId"`
	OtherIDType  string    `json:"otherIdType"`
	CreatedAt    time.Time `json:"createdAt"`
}
