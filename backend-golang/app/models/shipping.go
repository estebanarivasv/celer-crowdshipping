package models

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Shipping struct {
	gorm.Model
	ID                 int           `gorm:"primarykey"`
	SelectedOfferID    *int          `sql:"default:null"` // Selected Offer
	DistributorID      *int          `sql:"default:null"`
	Recipient          Recipient     `gorm:"type:json"`
	Offers             []Offer       `gorm:"preload:true"` // Has many
	Route              []RouteDetail `gorm:"preload:true"` // Has many
	SelectedOffer      Offer         `gorm:"preload:true"` // Has one
	Sender             User          `gorm:"preload:true"` // Has one
	Distributor        User          `gorm:"preload:true"` // Has one
	Package            Package       `gorm:"preload:true"` // Has one
	Details            string
	PackageID          int
	SenderID           int
	ProcessID          string
	PickedUpAt         time.Time
	DeliveredAt        time.Time
	OriginAddress      string
	DestinationAddress string
}

type Recipient struct {
	FirstName string
	LastName  string
}

// Scan - Parse the value from the database into a Recipient struct - []uint8 --> Recipient
func (r *Recipient) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), r)
}

// Value - Convert structure values into a format for db - Recipient --> []uint8
func (r *Recipient) Value() (driver.Value, error) {
	return json.Marshal(r)
}
