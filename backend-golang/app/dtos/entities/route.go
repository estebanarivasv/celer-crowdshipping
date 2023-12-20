package entities

type NewRouteDetailInDTO struct {
	Coordinates string `json:"coordinates"`
}

type RouteDetailInDTO struct {
	Coordinates string `json:"coordinates"`
	ShippingID  int    `json:"shippingId"`
}

type RouteDetailOutDTO struct {
	ID          int    `json:"id"`
	Coordinates string `json:"coordinates"`
	ShippingID  int    `json:"shippingId"`
}
