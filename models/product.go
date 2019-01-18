package models

import (
	"time"
)

// Product Information
type Product struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	Category         Category    `json:"category"`
	Images           []string    `json:"images" gorm:"type:text[]"`
	Value            int         `json:"value"`
	Weight           float64     `json:"weight"`
	UnavailableDates []time.Time `json:"unavailableDates"`
	TotalRating      float64     `json:"totalRating"`
	// Reviews          []Review    `json:"reviews"`
}

func getProducts() []Product {
	allProducts := []Product{
		Product{
			ID:          "0001",
			Name:        "Gopro",
			Description: "This is the description for gopro",
			Category: Category{
				ID:          "category001",
				Name:        "Photography",
				Image:       "",
				Description: "This is the Description for the Category",
			},
			Images:           nil,
			Value:            26000,
			Weight:           0.4,
			UnavailableDates: nil,
			TotalRating:      3.8,
			// Reviews:          nil,
		},
		Product{
			ID:          "0002",
			Name:        "Gimbal",
			Description: "This is the description for GIMBAL",
			Category: Category{
				ID:          "category001",
				Name:        "Photography",
				Image:       "",
				Description: "This is the Description for the Category",
			},
			Images:           nil,
			Value:            44000,
			Weight:           1.6,
			UnavailableDates: nil,
			TotalRating:      4.5,
			// Reviews:          nil,
		},
	}
	return allProducts
}

// GetProduct gets a product by id
func (product *Product) GetProduct(id string) int {
	allProducts := getProducts()
	statusCode := 0
	for _, item := range allProducts {
		if id == item.ID {
			statusCode = 1
			*product = item
		}
	}
	return statusCode
}

// GetAllProducts func gets all the products in the database
func GetAllProducts() (int, []Product) {
	allProducts := getProducts()
	return 1, allProducts
}
