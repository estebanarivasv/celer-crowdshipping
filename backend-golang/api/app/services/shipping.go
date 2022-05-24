package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
)

var repository = repositories.NewShippingRepository()
var mapper = mappers.ShippingMapper{}

// TODO comment service

func CreateShipping(shipping *entities.ShippingInDTO) dtos.Response {

	// Convert the dto to an entity
	shippingModel := mapper.FromDTO(shipping)

	query, err := repository.Create(shippingModel)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	newCamundaProcDTO, err := CreateNewCamundaProcInstance()
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	query.ProcessID = newCamundaProcDTO.ID
	queryWithProcID, err := repository.Save(&query)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	shippingDto := mapper.ToDTO(queryWithProcID)

	return dtos.Response{Success: true, Data: shippingDto}
}

func FindAllShippings() dtos.Response {

	var dtosArr []interface{}

	response, err := repository.FindAll()
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Convert and append all models into dtos
	for _, v := range response {
		dtosArr = append(dtosArr, mapper.ToDTO(&v))
	}

	return dtos.Response{Success: true, Data: dtosArr}

}

func FindShippingById(id int) dtos.Response {

	query, err := repository.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	shippingDto := mapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: shippingDto}

}

func UpdateShippingById(id int, dto entities.ShippingInPutDTO) dtos.Response {

	query, err := repository.UpdateById(id, &dto)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	shippingDto := mapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: shippingDto}
}

func DeleteShippingById(id int) dtos.Response {

	err := repository.DeleteById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	return dtos.Response{Success: true}

}
