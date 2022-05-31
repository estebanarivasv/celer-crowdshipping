package entities

type NewOfferInDTO struct {
	ShippingCost  float64 `json:"shippingCost"`
	Message       string  `json:"message"`
	Duration      float64 `json:"duration"`
	DistributorID int     `json:"distributorId"`
}

type OfferInDTO struct {
	ShippingCost  float64 `json:"shippingCost"`
	Message       string  `json:"message"`
	Duration      float64 `json:"duration"`
	ShippingID    int     `json:"requestId"`
	DistributorID int     `json:"distributorId"`
}

type OfferOutDTO struct {
	ID            int     `json:"id"`
	ShippingCost  float64 `json:"shippingCost"`
	Message       string  `json:"message"`
	Duration      float64 `json:"duration"`
	ShippingID    int     `json:"requestId"`
	DistributorID int     `json:"distributorId"`
	// CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt time.Time `json:"updatedAt"`
	// DeletedAt time.Time `json:"deletedAt"`
}
