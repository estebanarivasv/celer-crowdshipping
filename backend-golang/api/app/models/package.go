package models

import "gorm.io/gorm"

// Package TODO: Field-Level Permission gorm
type Package struct {
	gorm.Model
	ID          int     `gorm:"primarykey" json:"id"`
	Name        string  `json:"name"`
	ImageURL    string  `json:"image_url"`
	Description string  `json:"description"`
	Dimensions  string  `json:"dimensions"`
	Value       float64 `json:"value"` // Monetary value
}
