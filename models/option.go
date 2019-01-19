package models

// Option struct
type Option struct {
	ID         uint64 `json:"id" gorm:"primary_key"`
	OptionText string `json:"option"`
}
