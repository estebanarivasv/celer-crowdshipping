package entities

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"time"
)

type ShippingInDTO struct {
	Details            string           `json:"details"`
	PackageID          int              `json:"packageId"`
	OriginAddress      string           `json:"originAddress"`
	DestinationAddress string           `json:"destinationAddress"`
	Recipient          models.Recipient `json:"recipient"`
}

type ShippingInPatchDTO struct {
	SelectedOfferID int `json:"selectedOfferId"`
}

type ShippingInPatchWithDistributorDTO struct {
	SelectedOfferID int `json:"selectedOfferId"`
	DistributorID   int `json:"distributorId"`
}

type ShippingOutDTO struct {
	ID                 int         `json:"id"`
	Details            string      `json:"details"`
	Package            interface{} `json:"package"`
	OriginAddress      string      `json:"originAddress"`
	DestinationAddress string      `json:"destinationAddress"`
	ProcessID          string      `json:"processId"`
	SelectedOfferID    *int        `json:"selectedOfferId"`
	Sender             interface{} `json:"sender"`
	Recipient          interface{} `json:"recipient"`
	PickedUpAt         time.Time   `json:"pickedUpAt"`
	DeliveredAt        time.Time   `json:"deliveredAt"`
	CreatedAt          time.Time   `json:"createdAt"`
}

type ShippingOutBasicDTO struct {
	ID                 int       `json:"id"`
	Details            string    `json:"details"`
	OriginAddress      string    `json:"originAddress"`
	DestinationAddress string    `json:"destinationAddress"`
	PickedUpAt         time.Time `json:"pickedUpAt"`
	DeliveredAt        time.Time `json:"deliveredAt"`
}
