package models

import (
	"gorm.io/gorm"
)

// TODO: Field-Level Permission gorm
type Offer struct {
	gorm.Model
	ID            int     `gorm:"primarykey" json:"id"`
	ShippingCost  float64 `json:"shipping_cost"`
	Message       string  `json:"message"`
	Duration      string  `json:"duration"`   // Estimated duration
	ShippingID    uint    `json:"request_id"` // Has one relationship - Shipping
	DistributorID int     `json:"distributor_id"`
	Distributor   User
}
