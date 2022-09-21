package mappers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
)

type ShippingMapper struct {
}

// NewShippingMapper Returns a new instance
func NewShippingMapper() *ShippingMapper {
	return &ShippingMapper{}
}

// FromDTO Return a model with the dto that comes from the service
func (m ShippingMapper) FromDTO(dto *entities.ShippingInDTO) models.Shipping {

	return models.Shipping{
		Details:            dto.Details,
		SenderID:           dto.SenderID,
		PackageID:          dto.PackageID,
		RecipientID:        dto.RecipientID,
		OriginAddress:      dto.OriginAddress,
		DestinationAddress: dto.DestinationAddress,
	}

}

// ToDTO Return a dto with the model that comes from the database
func (m ShippingMapper) ToDTO(model *models.Shipping) interface{} {

	return entities.ShippingOutDTO{
		ID:                 model.ID,
		Details:            model.Details,
		Package:            PackageMapper{}.ToDTO(&model.Package),
		OriginAddress:      model.OriginAddress,
		DestinationAddress: model.DestinationAddress,
		ProcessID:          model.ProcessID,
		SelectedOfferID:    model.SelectedOfferID,
		Sender:             UserMapper{}.ToDTO(&model.Sender),
		Recipient:          UserMapper{}.ToDTO(&model.Recipient),
		PickedUpAt:         model.PickedUpAt,
		DeliveredAt:        model.DeliveredAt,
		CreatedAt:          model.CreatedAt,
		//UpdatedAt:   model.UpdatedAt,
		//DeletedAt:   model.DeletedAt.Time,
	}

}

// ToBasicDTO Return a dto with the model that comes from the database
func (m ShippingMapper) ToBasicDTO(model *models.Shipping) interface{} {

	// If model is empty, it returns an empty interface
	if model.ID == 0 {
		return nil
	}

	return entities.ShippingOutBasicDTO{
		ID:                 model.ID,
		Details:            model.Details,
		OriginAddress:      model.OriginAddress,
		DestinationAddress: model.DestinationAddress,
		PickedUpAt:         model.PickedUpAt,
		DeliveredAt:        model.DeliveredAt,
	}

}
