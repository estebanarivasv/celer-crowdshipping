package models

import (
	"gorm.io/gorm"
	"time"
)

// Shipping TODO: Field-Level Permission gorm
type Shipping struct {
	gorm.Model
	ID                 int           `gorm:"primarykey" json:"id"`
	ImageURL           string        `json:"image_url"`
	Details            string        `json:"details"`
	PackageID          int           `json:"package_id"`
	SelectedOfferID    int           `json:"selected_offer_id"` // Selected Offer
	SenderID           int           `json:"sender_id"`
	RecipientID        int           `json:"recipient_id"`
	PickedUpAt         time.Time     `json:"pickup_at"`
	DeliveredAt        time.Time     `json:"delivered_at"`
	OriginAddress      string        `json:"origin_addr"`
	DestinationAddress string        `json:"dest_addr"`
	Offers             []Offer       `json:"offers"` // Has many
	Route              []RouteDetail `json:"route"`  // Has many
	SelectedOffer      Offer
	Sender             User
	Recipient          User
	Package            Package
}
