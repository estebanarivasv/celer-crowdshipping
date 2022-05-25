package models

import (
	"gorm.io/gorm"
)

type RouteDetail struct {
	gorm.Model
	ID          int `gorm:"primarykey"`
	Coordinates string
	ShippingID  int
	Shipping    *Shipping // Avoid recursive types
}
