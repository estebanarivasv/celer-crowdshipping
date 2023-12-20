package mappers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
)

type OfferMapper struct {
}

// NewOfferMapper Returns a new instance
func NewOfferMapper() *OfferMapper {
	return &OfferMapper{}
}

// FromDTO Return a model with the dto that comes from the service
func (m OfferMapper) FromDTO(dto *entities.OfferInDTO) models.Offer {

	return models.Offer{
		ShippingCost: dto.ShippingCost,
		Message:      dto.Message,
		Duration:     dto.Duration,
		ShippingID:   dto.ShippingID,
		//DistributorID: dto.DistributorID,
	}

}

// ToDTO Return a dto with the model that comes from the database
func (m OfferMapper) ToDTO(model *models.Offer) interface{} {

	// If model is empty, it returns an empty interface
	if model.ID == 0 {
		return nil
	}

	return entities.OfferOutDTO{
		ID:            model.ID,
		ShippingCost:  model.ShippingCost,
		Message:       model.Message,
		Duration:      model.Duration,
		ShippingID:    model.ShippingID,
		DistributorID: model.DistributorID,
		//CreatedAt:   model.CreatedAt,
		//UpdatedAt:   model.UpdatedAt,
		//DeletedAt:   model.DeletedAt.Time,
	}

}

func (m OfferMapper) ToDetailedDTO(model *models.Offer) interface{} {

	// If model is empty, it returns an empty interface
	if model.ID == 0 {
		return nil
	}

	return entities.DetailedOfferOutDTO{
		ID:           model.ID,
		ShippingCost: model.ShippingCost,
		Message:      model.Message,
		Duration:     model.Duration,
		ShippingID:   model.ShippingID,
		Distributor:  UserMapper{}.ToDTO(&model.Distributor),
		//CreatedAt:   model.CreatedAt,
		//UpdatedAt:   model.UpdatedAt,
		//DeletedAt:   model.DeletedAt.Time,
	}

}
