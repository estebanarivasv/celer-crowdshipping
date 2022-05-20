package models

import "time"

type OfferInDTO struct {
	ShippingCost  float64 `json:"shipping_cost"`
	Message       string  `json:"message"`
	Duration      float64 `json:"duration"`
	ShippingID    int     `json:"request_id"`
	DistributorID int     `json:"distributor_id"`
}

type OfferOutDTO struct {
	ID           int        `json:"id"`
	ShippingCost float64    `json:"shipping_cost"`
	Message      string     `json:"message"`
	Duration     float64    `json:"duration"`
	ShippingID   int        `json:"request_id"`
	Distributor  UserOutDTO `json:"distributor"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    time.Time  `json:"deleted_at"`
}
