package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
)

var repository = repositories.NewShippingRepository()
var shippingMapper = mappers.ShippingMapper{}
var camundaMapper = mappers.CamundaMapper{}

// TODO comment service

func CreateShipping(shipping *entities.ShippingInDTO) dtos.Response {

	// Convert the dto to an entity
	shippingModel := shippingMapper.FromDTO(shipping)

	query, err := repository.Create(shippingModel)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	newCamundaProcDTO, err := CreateNewCamundaProcInstance()
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Add Camunda Process ID
	query.ProcessID = newCamundaProcDTO.ID
	queryWithProcID, err := repository.Save(&query)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Send ShippingDetailsAdded message to Camunda
	SendMessageToCamundaProcess(newCamundaProcDTO.ID, "ShippingDetailsAdded")

	// Mapping into a dto
	shippingDto := shippingMapper.ToDTO(queryWithProcID)

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
		dtosArr = append(dtosArr, shippingMapper.ToDTO(&v))
	}

	return dtos.Response{Success: true, Data: dtosArr}

}

func FindShippingById(id int) dtos.Response {

	query, err := repository.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	shippingDto := shippingMapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: shippingDto}

}

func UpdateShippingById(id int, dto entities.ShippingInPutDTO) dtos.Response {

	query, err := repository.UpdateById(id, &dto)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	shippingDto := shippingMapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: shippingDto}
}

func DeleteShippingById(id int) dtos.Response {

	err := repository.DeleteById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	return dtos.Response{Success: true}

}

func FindShippingStateById(id int) dtos.Response {

	query, err := repository.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	camundaID := query.ProcessID
	stateDto, err := GetProcInstanceState(camundaID)
	if err != nil {
		return dtos.Response{}
	}

	return dtos.Response{Success: true, Data: stateDto}

}

func UpdateShippingState(shippingId int, message string) dtos.Response {

	// Bring shipping
	query, err := repository.FindOneById(shippingId)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Bring proc id
	camundaID := query.ProcessID

	// Send msg to proc
	return SendMessageToCamundaProcess(camundaID, message)
}
