package mappers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
)

type PackageMapper struct {
}

// NewPackageMapper Returns a new instance
func NewPackageMapper() *PackageMapper {
	return &PackageMapper{}
}

// FromDTO Return a model with the dto that comes from the service
func (m PackageMapper) FromDTO(dto *entities.PackageInDTO) models.Package {

	return models.Package{
		Name:        dto.Name,
		ImageURL:    dto.ImageURL,
		Description: dto.Description,
		Dimensions:  dto.Dimensions,
		Value:       dto.Value,
	}

}

// ToDTO Return a dto with the model that comes from the database
func (m PackageMapper) ToDTO(model *models.Package) interface{} {

	// If model is empty, it returns an empty interface
	if model.ID == 0 {
		return nil
	}

	return entities.PackageOutDTO{
		ID:          model.ID,
		Name:        model.Name,
		ImageURL:    model.ImageURL,
		Description: model.Description,
		Dimensions:  model.Dimensions,
		Value:       model.Value,
		//CreatedAt:   model.CreatedAt,
		//UpdatedAt:   model.UpdatedAt,
		//DeletedAt:   model.DeletedAt.Time,
	}

}
