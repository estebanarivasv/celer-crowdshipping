package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
	"time"
)

var shippingRepository = repositories.NewShippingRepository()
var shippingMapper = mappers.ShippingMapper{}

// TODO: comment service

func CreateShipping(shipping *entities.ShippingInDTO) dtos.Response {

	// Convert the dto to an entity
	shippingModel := shippingMapper.FromDTO(shipping)

	query, err := shippingRepository.Create(shippingModel)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	newCamundaProcDTO, err := CreateNewCamundaProcInstance()
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Add Camunda Process ID
	query.ProcessID = newCamundaProcDTO.ID
	queryWithProcID, err := shippingRepository.Save(&query)
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

	response, err := shippingRepository.FindAll()
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Convert and append all models into dtos
	for _, v := range response {
		dtosArr = append(dtosArr, shippingMapper.ToBasicDTO(&v))
	}

	return dtos.Response{Success: true, Data: dtosArr}

}

func FindShippingById(id int) dtos.Response {

	query, err := shippingRepository.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	shippingDto := shippingMapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: shippingDto}

}

func DeleteShippingById(id int) dtos.Response {

	err := shippingRepository.DeleteById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	return dtos.Response{Success: true}

}

func FindShippingStateById(id int) dtos.Response {

	query, err := shippingRepository.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	camundaID := query.ProcessID

	stateDto, err := GetProcInstanceState(camundaID)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}
	if stateDto.State == "COMPLETED" {
		return dtos.Response{Success: true, Error: "the process has been successfully completed"}
	}

	currentTaskDto, err := GetProcInstanceCurrentTask(camundaID)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	return dtos.Response{Success: true, Data: currentTaskDto}

}

func UpdateShippingState(shippingId int, message string) dtos.Response {

	// Bring shipping
	query, err := shippingRepository.FindOneById(shippingId)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Update PICKED_UP_AT or DELIVERED_AT
	if message == "PackageInTransit" {
		query.PickedUpAt = time.Now()
		query.UpdatedAt = time.Now()
		_, err := shippingRepository.Save(&query)
		if err != nil {
			return dtos.Response{Success: false, Error: err.Error()}
		}
	} else if message == "DeliveredToRecipient" {
		query.DeliveredAt = time.Now()
		query.UpdatedAt = time.Now()
		_, err := shippingRepository.Save(&query)
		if err != nil {
			return dtos.Response{Success: false, Error: err.Error()}
		}
	}

	// Bring proc id
	camundaID := query.ProcessID

	// Send msg to proc
	return SendMessageToCamundaProcess(camundaID, message)
}

func UpdateSelectedOffer(shippingId int, dto entities.ShippingInPatchDTO) dtos.Response {

	offerDto := FindOfferById(dto.SelectedOfferID)
	if !offerDto.Success {
		return offerDto
	}

	query, err := shippingRepository.UpdateById(shippingId, dto)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Send OfferSelected message to Camunda
	SendMessageToCamundaProcess(query.ProcessID, "OfferSelected")

	// Mapping into a dto
	shippingDto := shippingMapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: shippingDto}

}

func FindShippingRequests() dtos.Response {

	conditions := make(map[string]interface{})
	conditions["selected_offer_id"] = nil

	requests, err := shippingRepository.FindFilteredShippings(conditions)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Map all models into a dto
	var requestsArr []interface{}
	for _, r := range requests {
		requestsArr = append(requestsArr, shippingMapper.ToBasicDTO(&r))
	}

	return dtos.Response{Success: true, Data: requestsArr}

}

func FindActiveShippings() dtos.Response {
	requests, err := shippingRepository.FindCurrentShippings()
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Map all models into a dto
	var requestsArr []interface{}
	for _, r := range requests {
		requestsArr = append(requestsArr, shippingMapper.ToBasicDTO(&r))
	}

	return dtos.Response{Success: true, Data: requestsArr}

}
