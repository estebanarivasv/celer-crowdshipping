package mappers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
)

type RouteMapper struct {
}

// FromDTO Return a model with the dto that comes from the service
func (m RouteMapper) FromDTO(dto *entities.RouteDetailInDTO) models.RouteDetail {

	return models.RouteDetail{
		Coordinates: dto.Coordinates,
		ShippingID:  dto.ShippingID,
	}

}

// ToDTO Return a dto with the model that comes from the database
func (m RouteMapper) ToDTO(model *models.RouteDetail) interface{} {

	// If model is empty, it returns an empty interface
	if model.ID == 0 {
		return nil
	}

	return entities.RouteDetailOutDTO{
		ID:          model.ID,
		Coordinates: model.Coordinates,
		ShippingID:  model.ShippingID,
		//CreatedAt:   model.CreatedAt,
		//UpdatedAt:   model.UpdatedAt,
		//DeletedAt:   model.DeletedAt.Time,
	}

}
