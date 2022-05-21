package models

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"gorm.io/gorm"
	"time"
)

// Shipping TODO: Field-Level Permission gorm
type Shipping struct {
	gorm.Model
	ID                 int           `gorm:"primarykey" json:"id"`
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

func (model Shipping) FromDTO(dto entities.ShippingInDTO) interface{} {

	return Shipping{
		Details:            dto.Details,
		PackageID:          dto.PackageID,
		SelectedOfferID:    dto.SelectedOfferID,
		SenderID:           dto.SenderID,
		RecipientID:        dto.RecipientID,
		OriginAddress:      dto.OriginAddress,
		DestinationAddress: dto.DestinationAddress,
	}

}

func (model Shipping) ToDTO() entities.ShippingOutDTO {

	// Each coordinate is mapped into a dto and appended to a slice
	var routeSlice []entities.RouteDetailOutDTO
	for _, r := range model.Route {
		routeSlice = append(routeSlice,
			r.ToDTO())
	}

	// Each coordinate is mapped into a dto and appended to a slice
	var offersSlice []entities.OfferOutDTO
	for _, o := range model.Offers {
		offersSlice = append(offersSlice,
			o.ToDTO())
	}

	return entities.ShippingOutDTO{
		ID:                 model.ID,
		Details:            model.Details,
		Package:            model.Package.ToDTO(),
		OriginAddress:      model.OriginAddress,
		DestinationAddress: model.DestinationAddress,
		Offers:             offersSlice,
		SelectedOffer:      model.SelectedOffer.ToDTO(),
		Sender:             model.Sender.ToDTO(),
		Recipient:          model.Recipient.ToDTO(),
		Route:              routeSlice,
		PickedUpAt:         model.PickedUpAt,
		DeliveredAt:        model.DeliveredAt,
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
		DeletedAt:          model.DeletedAt.Time,
	}

}
