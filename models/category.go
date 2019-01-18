package models

// Category information
type Category struct {
	ID               string `json:"id"`
	Image            string `json:"image"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	NumberOfProducts int    `json:"noOfProducts"`
}
