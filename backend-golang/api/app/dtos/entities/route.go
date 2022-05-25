package entities

import "time"

type NewRouteDetailInDTO struct {
	Coordinates string `json:"coordinates"`
}

type RouteDetailInDTO struct {
	Coordinates string `json:"coordinates"`
	ShippingID  int    `json:"shipping_id"`
}

type RouteDetailOutDTO struct {
	ID          int       `json:"id"`
	Coordinates string    `json:"coordinates"`
	ShippingID  int       `json:"shipping_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
