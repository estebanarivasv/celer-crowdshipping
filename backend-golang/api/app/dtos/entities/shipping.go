package entities

import (
	"time"
)

type ShippingInDTO struct {
	Details            string `json:"details"`
	PackageID          int    `json:"package_id"`
	SelectedOfferID    int    `json:"selected_offer_id"`
	SenderID           int    `json:"sender_id"`
	RecipientID        int    `json:"recipient_id"`
	OriginAddress      string `json:"origin_addr"`
	DestinationAddress string `json:"dest_addr"`
	// PickedUpAt         time.Time        `json:"pickup_at"`
	// DeliveredAt        time.Time        `json:"delivered_at"`
	// Route              []RouteDetailDTO `json:"route"`
}

type ShippingOutDTO struct {
	ID                 int                 `json:"id"`
	Details            string              `json:"details"`
	Package            PackageOutDTO       `json:"package"`
	OriginAddress      string              `json:"origin_addr"`
	DestinationAddress string              `json:"dest_addr"`
	Offers             []OfferOutDTO       `json:"offers"`
	SelectedOffer      OfferOutDTO         `json:"selected_offer"`
	Sender             UserOutDTO          `json:"sender"`
	Recipient          UserOutDTO          `json:"recipient"`
	Route              []RouteDetailOutDTO `json:"route"`
	PickedUpAt         time.Time           `json:"pickup_at"`
	DeliveredAt        time.Time           `json:"delivered_at"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	DeletedAt          time.Time           `json:"deleted_at"`
}
