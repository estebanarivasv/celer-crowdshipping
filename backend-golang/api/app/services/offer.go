package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
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
