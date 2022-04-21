package models

type Shipping struct {
	// TODO: Field-Level Permission gorm
	ID              uint   `gorm:"primaryKey" json:"id"`
	Image           string `json:"image"`
	Details         string `json:"details"`
	StartAddr       string `json:"start_addr"`
	DestinationAddr string `json:"destination_addr"`
}
