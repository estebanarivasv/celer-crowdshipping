package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
)

var offerRepository = repositories.NewOfferRepository()
var offerMapper = mappers.OfferMapper{}

func FindOfferById(id int) dtos.Response {
	query, err := offerRepository.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	dto := offerMapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: dto}

}

func CreateOffer(offerDTO *entities.OfferInDTO) dtos.Response {

	// Get shipping
	shippingDao, err := shippingRepository.FindOneById(offerDTO.ShippingID)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Verify Shipping has selected_offer_id null
	if shippingDao.SelectedOfferID != nil {
		return dtos.Response{Success: false, Error: "this shipping id does not refer to a shipping request"}
	}

	// Convert the dto to an entity
	model := offerMapper.FromDTO(offerDTO)

	query, err := offerRepository.Create(model)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	dto := offerMapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: dto}
}
