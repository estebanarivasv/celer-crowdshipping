package models

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"gorm.io/gorm"
)

type RouteDetail struct {
	gorm.Model
	ID          int       `gorm:"primarykey" json:"id"`
	Coordinates string    `json:"coordinates"`
	ShippingID  int       `json:"shipping_id"`
	Shipping    *Shipping // Avoid recursive types
}

func (model RouteDetail) FromDTO(dto entities.RouteDetailInDTO) interface{} {

	return RouteDetail{
		Coordinates: dto.Coordinates,
		ShippingID:  dto.ShippingID,
	}

}

func (model RouteDetail) ToDTO() entities.RouteDetailOutDTO {

	return entities.RouteDetailOutDTO{
		ID:          model.ID,
		Coordinates: model.Coordinates,
		ShippingID:  model.ShippingID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		DeletedAt:   model.DeletedAt.Time,
	}

}
