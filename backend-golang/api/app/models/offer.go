package models

import "gorm.io/gorm"

// TODO: Field-Level Permission gorm
type Offer struct {
	gorm.Model
	ID            int     `gorm:"primarykey" json:"id"`
	ShippingCost  float64 `json:"shipping_cost"`
	Message       string  `json:"message"`
	Duration      float64 `json:"duration"`   // Estimated duration - hours
	ShippingID    int     `json:"request_id"` // Has one relationship - Shipping
	DistributorID int     `json:"distributor_id"`
	Distributor   User
	Shipping      *Shipping // Avoid recursive types
}
