package models

import (
	"gorm.io/gorm"
	"time"
)

// Shipping TODO: Field-Level Permission gorm
type Shipping struct {
	gorm.Model
	ID                 int `gorm:"primarykey"`
	Details            string
	PackageID          int
	SelectedOfferID    *int `sql:"default:null"` // Selected Offer
	SenderID           int
	RecipientID        int
	ProcessID          string
	PickedUpAt         time.Time
	DeliveredAt        time.Time
	OriginAddress      string
	DestinationAddress string
	Offers             []Offer       `gorm:"preload:true"` // Has many
	Route              []RouteDetail `gorm:"preload:true"` // Has many
	SelectedOffer      Offer         `gorm:"preload:true"` // Has one
	Sender             User          `gorm:"preload:true"` // Has one
	Recipient          User          `gorm:"preload:true"` // Has one
	Package            Package       `gorm:"preload:true"` // Has one
}
