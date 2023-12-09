package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/camunda"
	"net/http"
)

type CamundaService struct {
	camundaRepo *repositories.CamundaRepository
}

// NewCamundaService Returns a new instance
func NewCamundaService() *CamundaService {

	var repo = repositories.NewCamundaRepository()

	return &CamundaService{
		camundaRepo: repo,
	}
}

// CreateNewCamundaProcInstance Camunda service that creates a new process instance
func (s *CamundaService) CreateNewCamundaProcInstance() (*dtos.BasicCamundaProcessDTO, error) {

	url := camunda.GetCreateShippingProcURL()

	procDTO, err := s.camundaRepo.SendNewInstanceRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	return procDTO, nil
}

// SendMessageToCamundaProcess Camunda service that sends a message to a process instance
func (s *CamundaService) SendMessageToCamundaProcess(procId string, msgName string) dtos.Response {

	url := camunda.GetMessageProcURL()
	payload := dtos.MessageToProcessDTO{
		ProcessInstanceId: procId,
		MessageName:       msgName,
	}
	body, err := camunda.PayloadToJsonBody(payload)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	err = s.camundaRepo.SendMessageToProcessRequest(http.MethodPost, url, body)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	return dtos.Response{Success: true}

}

func (s *CamundaService) GetProcInstanceCurrentTask(procId string) (dtos.DetailedCamundaProcessDTO, error) {

	url := camunda.GetProcCurrentActivityURL(procId)

	procDTO, err := s.camundaRepo.GetProcCurrentActivity(http.MethodGet, url)
	if err != nil {
		return *new(dtos.DetailedCamundaProcessDTO), err
	}
	return procDTO, nil
}

func (s *CamundaService) GetProcInstanceState(procId string) (dtos.DetailedCamundaProcessDTO, error) {

	url := camunda.GetProcStateURL(procId)

	procDTO, err := s.camundaRepo.GetProcCurrentActivity(http.MethodGet, url)
	if err != nil {
		return *new(dtos.DetailedCamundaProcessDTO), err
	}
	return procDTO, nil
}
