package models

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"gorm.io/gorm"
)

// Package TODO: Field-Level Permission gorm
type Package struct {
	gorm.Model
	ID          int     `gorm:"primarykey" json:"id"`
	Name        string  `json:"name"`
	ImageURL    string  `json:"image_url"`
	Description string  `json:"description"`
	Dimensions  string  `json:"dimensions"`
	Value       float64 `json:"value"` // Monetary value
}

func (model Package) FromDTO(dto entities.PackageInDTO) interface{} {

	return Package{
		Name:        dto.Name,
		ImageURL:    dto.ImageURL,
		Description: dto.Description,
		Dimensions:  dto.Dimensions,
		Value:       dto.Value,
	}

}

func (model Package) ToDTO() entities.PackageOutDTO {

	return entities.PackageOutDTO{
		ID:          model.ID,
		Name:        model.Name,
		ImageURL:    model.ImageURL,
		Description: model.Description,
		Dimensions:  model.Dimensions,
		Value:       model.Value,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		DeletedAt:   model.DeletedAt.Time,
	}

}
