package models

import "gorm.io/gorm"

type RouteDetail struct {
	gorm.Model
	ID          int `gorm:"primarykey" json:"id"`
	Coordinates string
	ShippingID  int
}
