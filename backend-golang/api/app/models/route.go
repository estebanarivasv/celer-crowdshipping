package models

import "gorm.io/gorm"

type RouteDetail struct {
	gorm.Model
	ID          int    `gorm:"primarykey" json:"id"`
	Coordinates string `json:"coordinates"`
	ShippingID  int    `json:"shipping_id"`
}
