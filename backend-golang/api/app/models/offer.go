package models

import (
	"gorm.io/gorm"
)

// TODO: Field-Level Permission gorm
type Offer struct {
	gorm.Model
	ID            int `gorm:"primarykey"`
	ShippingCost  float64
	Message       string
	Duration      float64 // Estimated duration - hours
	ShippingID    int     // Has one relationship - Shipping
	DistributorID int
	Distributor   User      `gorm:"PRELOAD:true"`
	Shipping      *Shipping // Avoid recursive types
}
