package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
)

var offerRepository = repositories.NewOfferRepository()
var userRepository = repositories.NewUserRepository()
var offerMapper = mappers.OfferMapper{}

func FindOfferById(id int) dtos.Response {
	query, err := offerRepository.FindOneById(&id)
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

func FindOffersByRequestID(id int) dtos.Response {

	// Get shipping
	shippingDao, err := shippingRepository.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Verify Shipping has selected_offer_id null
	if shippingDao.SelectedOfferID != nil {
		return dtos.Response{Success: false, Error: "this shipping id does not refer to a shipping request"}
	}

	var dtosArr []interface{}

	conditions := make(map[string]interface{})
	conditions["shipping_id"] = id

	offers, err := offerRepository.FindAllRequestOffers(conditions)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Convert and append all models into dtos
	for _, offer := range offers {
		dtosArr = append(dtosArr, offerMapper.ToDTO(&offer))
	}

	return dtos.Response{Success: true, Data: dtosArr}

}

func DeleteOfferById(id int) dtos.Response {
	err := offerRepository.DeleteById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	return dtos.Response{Success: true}
}

func FindOffersByDistributorID(id int) dtos.Response {

	var dtosArr []interface{}

	// Verify user exists
	user, err := userRepository.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	conditions := make(map[string]interface{})
	conditions["distributor_id"] = user.ID

	offers, err := offerRepository.FindAllRequestOffers(conditions)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Convert and append all models into dtos
	for _, offer := range offers {
		dtosArr = append(dtosArr, offerMapper.ToDTO(&offer))
	}

	return dtos.Response{Success: true, Data: dtosArr}

}

func FindSelectedOfferByShippingId(id int) dtos.Response {

	// Verify shipping exists
	shipping, err := shippingRepository.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Verify selected offer is not null
	if shipping.SelectedOfferID == nil {
		return dtos.Response{Success: false, Error: "there is not a selected offer in this shipping instance."}
	}

	dto, err := offerRepository.FindOneById(shipping.SelectedOfferID)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	return dtos.Response{Success: true, Data: dto}
}
