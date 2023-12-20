package models

import (
	"gorm.io/gorm"
)

type Package struct {
	gorm.Model
	ID          int `gorm:"primarykey"`
	Name        string
	ImageURL    string
	Description string
	Dimensions  string
	Value       float64 `json:"value"` // Monetary value
}
