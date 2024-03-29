package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
	"time"
)

type ShippingService struct {
	shippingRepo   *repositories.ShippingRepository
	offerRepo      *repositories.OfferRepository
	camundaService *CamundaService
	mapper         *mappers.ShippingMapper
}

// NewShippingService Returns a new instance
func NewShippingService() *ShippingService {

	var shippingRepository = repositories.NewShippingRepository()
	var offerRepo = repositories.NewOfferRepository()
	var camundaService = NewCamundaService()
	var shippingMapper = mappers.NewShippingMapper()

	return &ShippingService{
		shippingRepo:   shippingRepository,
		offerRepo:      offerRepo,
		camundaService: camundaService,
		mapper:         shippingMapper,
	}
}

func (s *ShippingService) CreateShipping(shipping *entities.ShippingInDTO, senderIDFromContext int) dtos.Response {

	// Convert the dto to an entity
	shippingModel := s.mapper.FromDTO(shipping)
	shippingModel.SenderID = senderIDFromContext // Set Sender from the context

	query, err := s.shippingRepo.Create(shippingModel)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	newCamundaProcDTO, err := s.camundaService.CreateNewCamundaProcInstance()
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Add Camunda Process ID
	if newCamundaProcDTO.ID == "" {
		return dtos.Response{Success: false, Error: "error generating an instance of process definition"}
	}
	query.ProcessID = newCamundaProcDTO.ID
	queryWithProcID, err := s.shippingRepo.Save(&query)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Send ShippingDetailsAdded message to Camunda
	s.camundaService.SendMessageToCamundaProcess(newCamundaProcDTO.ID, "ShippingDetailsAdded")

	// Mapping into a dto
	shippingDto := s.mapper.ToDTO(queryWithProcID)

	return dtos.Response{Success: true, Data: shippingDto}
}

func (s *ShippingService) FindAllMyShippings(userID int) dtos.Response {

	var dtosArr []interface{}

	response, err := s.shippingRepo.FindAllByRecipientID(userID)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Convert and append all shippings into dtos
	for _, v := range response {
		dtosArr = append(dtosArr, s.mapper.ToBasicDTO(&v))
	}

	return dtos.Response{Success: true, Data: dtosArr}

}

func (s *ShippingService) FindShippingById(id int) dtos.Response {

	query, err := s.shippingRepo.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	shippingDto := s.mapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: shippingDto}

}

func (s *ShippingService) DeleteShippingById(id int) dtos.Response {

	err := s.shippingRepo.DeleteById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	return dtos.Response{Success: true}

}

func (s *ShippingService) FindShippingStateById(id int) dtos.Response {

	query, err := s.shippingRepo.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	camundaID := query.ProcessID

	stateDto, err := s.camundaService.GetProcInstanceState(camundaID)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}
	if stateDto.State == "COMPLETED" {
		return dtos.Response{Success: true, Error: "the process has been successfully completed"}
	}

	currentTaskDto, err := s.camundaService.GetProcInstanceCurrentTask(camundaID)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	return dtos.Response{Success: true, Data: currentTaskDto}

}

func (s *ShippingService) UpdateShippingState(shippingId int, message string) dtos.Response {

	// Bring shipping
	query, err := s.shippingRepo.FindOneById(shippingId)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Update PICKED_UP_AT or DELIVERED_AT
	if message == "PackageInTransit" {
		query.PickedUpAt = time.Now()
		query.UpdatedAt = time.Now()
		_, err := s.shippingRepo.Save(&query)
		if err != nil {
			return dtos.Response{Success: false, Error: err.Error()}
		}
	} else if message == "DeliveredToRecipient" {
		query.DeliveredAt = time.Now()
		query.UpdatedAt = time.Now()
		_, err := s.shippingRepo.Save(&query)
		if err != nil {
			return dtos.Response{Success: false, Error: err.Error()}
		}
	}

	// Bring proc id
	camundaID := query.ProcessID

	// Send msg to proc
	return s.camundaService.SendMessageToCamundaProcess(camundaID, message)
}

func (s *ShippingService) UpdateSelectedOffer(shippingId int, dto entities.ShippingInPatchDTO) dtos.Response {

	var updatesDto = entities.ShippingInPatchWithDistributorDTO{}

	offer, err := s.offerRepo.FindOneById(&dto.SelectedOfferID)
	if err != nil {
		return dtos.Response{Success: false, Error: "could not find offer"}
	}

	updatesDto.SelectedOfferID = dto.SelectedOfferID
	updatesDto.DistributorID = offer.Distributor.ID

	query, err := s.shippingRepo.UpdateById(shippingId, updatesDto)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Send OfferSelected message to Camunda
	s.camundaService.SendMessageToCamundaProcess(query.ProcessID, "OfferSelected")

	// Mapping into a dto
	shippingDto := s.mapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: shippingDto}

}

func (s *ShippingService) FindShippingRequests() dtos.Response {

	conditions := make(map[string]interface{})
	conditions["selected_offer_id"] = nil

	requests, err := s.shippingRepo.FindFilteredShippings(conditions)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Map all bpmn-models into a dto
	var requestsArr []interface{}
	for _, r := range requests {
		requestsArr = append(requestsArr, s.mapper.ToBasicDTO(&r))
	}

	return dtos.Response{Success: true, Data: requestsArr}

}

func (s *ShippingService) FindActiveShippingsByDistributorID(userID int) dtos.Response {
	requests, err := s.shippingRepo.FindCurrentShippingsByUserID(userID)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Map all bpmn-models into a dto
	var requestsArr []interface{}
	for _, r := range requests {
		requestsArr = append(requestsArr, s.mapper.ToBasicDTO(&r))
	}

	return dtos.Response{Success: true, Data: requestsArr}

}
