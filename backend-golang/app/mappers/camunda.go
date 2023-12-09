package mappers

import (
	"bytes"
	"encoding/json"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"log"
)

type CamundaMapper struct{}

// NewCamundaMapper Returns a new instance
func NewCamundaMapper() *CamundaMapper {
	return &CamundaMapper{}
}

func (m *CamundaMapper) JsonBodyToCamundaProcessDTO(jsonResponse []byte) (*dtos.BasicCamundaProcessDTO, error) {
	var data dtos.BasicCamundaProcessDTO
	err := json.Unmarshal(jsonResponse, &data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &data, nil
}

func (m *CamundaMapper) JsonBodyToDetailedProcessDTO(jsonResponse []byte) (dtos.DetailedCamundaProcessDTO, error) {
	var data dtos.DetailedCamundaProcessDTO
	err := json.Unmarshal(jsonResponse, &data)
	if err != nil {
		log.Fatal(err)
		return *new(dtos.DetailedCamundaProcessDTO), err
	}
	return data, nil
}

func (m *CamundaMapper) PayloadToJsonBody(payload interface{}) (*bytes.Buffer, error) {

	encodedPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewBuffer(encodedPayload)

	return reader, nil
}
