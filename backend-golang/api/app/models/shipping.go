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
	SelectedOfferID    int // Selected Offer
	SenderID           int
	RecipientID        int
	ProcessID          string
	PickedUpAt         time.Time
	DeliveredAt        time.Time
	OriginAddress      string
	DestinationAddress string
	Offers             []Offer       `gorm:"PRELOAD:true"` // Has many
	Route              []RouteDetail `gorm:"PRELOAD:true"` // Has many
	SelectedOffer      Offer         `gorm:"PRELOAD:true"` // Has one
	Sender             User          `gorm:"PRELOAD:true"` // Has one
	Recipient          User          `gorm:"PRELOAD:true"` // Has one
	Package            Package       `gorm:"PRELOAD:true"` // Has one
}
