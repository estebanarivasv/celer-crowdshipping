package mappers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/controllers"
)

type OfferMapper struct {
}

// FromDTO Return a model with the dto that comes from the service
func (m OfferMapper) FromDTO(dto *entities.OfferInDTO) models.Offer {

	return models.Offer{
		ShippingCost:  dto.ShippingCost,
		Message:       dto.Message,
		Duration:      dto.Duration,
		ShippingID:    dto.ShippingID,
		DistributorID: dto.DistributorID,
	}

}

// ToDTO Return a dto with the model that comes from the database
func (m OfferMapper) ToDTO(model *models.Offer) interface{} {

	// If model is empty, it returns an empty interface
	if controllers.IsZero[*models.Offer](model) {
		return *new(interface{})
	}

	return entities.OfferOutDTO{
		ID:            model.ID,
		ShippingCost:  model.ShippingCost,
		Message:       model.Message,
		Duration:      model.Duration,
		ShippingID:    model.ShippingID,
		DistributorID: model.DistributorID,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		DeletedAt:     model.DeletedAt.Time,
	}

}
