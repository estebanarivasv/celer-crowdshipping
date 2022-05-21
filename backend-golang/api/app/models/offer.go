package models

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"gorm.io/gorm"
)

// TODO: Field-Level Permission gorm
type Offer struct {
	gorm.Model
	ID            int     `gorm:"primarykey" json:"id"`
	ShippingCost  float64 `json:"shipping_cost"`
	Message       string  `json:"message"`
	Duration      float64 `json:"duration"`   // Estimated duration - hours
	ShippingID    int     `json:"request_id"` // Has one relationship - Shipping
	DistributorID int     `json:"distributor_id"`
	Distributor   User
	Shipping      *Shipping // Avoid recursive types
}

func (model Offer) FromDTO(dto entities.OfferInDTO) interface{} {

	return Offer{
		ShippingCost:  dto.ShippingCost,
		Message:       dto.Message,
		Duration:      dto.Duration,
		ShippingID:    dto.ShippingID,
		DistributorID: dto.DistributorID,
	}

}

func (model Offer) ToDTO() entities.OfferOutDTO {

	return entities.OfferOutDTO{
		ID:           model.ID,
		ShippingCost: model.ShippingCost,
		Message:      model.Message,
		Duration:     model.Duration,
		ShippingID:   model.ShippingID,
		Distributor:  model.Distributor.ToDTO(),
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    model.DeletedAt.Time,
	}

}
