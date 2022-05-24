package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/camunda"
	"net/http"
)

// CreateNewCamundaProcInstance Camunda service that creates a new process instance
func CreateNewCamundaProcInstance() (*dtos.BasicCamundaProcessDTO, error) {

	url := camunda.GetCreateShippingProcURL()

	procDTO, err := repositories.SendNewInstanceRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	return procDTO, nil
}

// SendMessageToCamundaProcess Camunda service that sends a message to a process instance
func SendMessageToCamundaProcess(msgName string, procId string) dtos.Response {

	url := camunda.GetMessageProcURL()
	payload := dtos.MessageToProcessDTO{
		MessageName:       msgName,
		ProcessInstanceId: procId,
	}
	body, err := camunda.PayloadToJsonBody(payload)
	if err != nil {
		return dtos.Response{Success: true, Error: err.Error()}
	}

	err = repositories.SendMessageToProcessRequest(http.MethodPost, url, body)
	if err != nil {
		return dtos.Response{Success: true, Error: err.Error()}
	}

	return dtos.Response{Success: true}

}
