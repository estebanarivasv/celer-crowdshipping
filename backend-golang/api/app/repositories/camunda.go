package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	shipping "github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CamundaRepository struct {
	client *http.Client
	mapper *mappers.CamundaMapper
}

// NewCamundaRepository Returns a new instance
func NewCamundaRepository() *CamundaRepository {

	var client = &http.Client{Timeout: 10 * time.Second}
	var mapper = mappers.NewCamundaMapper()

	return &CamundaRepository{
		client: client,
		mapper: mapper,
	}
}

func (r *CamundaRepository) SendNewInstanceRequest(httpMethod string, url string, body io.Reader) (*shipping.BasicCamundaProcessDTO, error) {

	// Define the request
	request, err := http.NewRequest(httpMethod, url, body)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	// Make a request and return a response
	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("%s", err)
		}
	}(response.Body)

	if response.StatusCode == http.StatusBadRequest {
		return nil, http.ErrBodyNotAllowed
	}

	dao, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data any // TODO DATA NO HACE NADA WTF
	err = json.Unmarshal(dao, &data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	dto, err := r.mapper.JsonBodyToCamundaProcessDTO(dao)
	if err != nil {
		return nil, err
	}
	return dto, nil

}

func (r *CamundaRepository) SendMessageToProcessRequest(httpMethod string, url string, body io.Reader) error {

	request, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := r.client.Do(request)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("%s", err)
		}
	}(response.Body)

	if response.StatusCode == http.StatusBadRequest {
		// 400 - BadRequest
		return errors.New("this process is not awaiting for that message")
	}
	return nil
}

func (r *CamundaRepository) GetProcCurrentActivity(httpMethod string, url string) (shipping.DetailedCamundaProcessDTO, error) {

	request, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return *new(shipping.DetailedCamundaProcessDTO), err
	}

	response, err := r.client.Do(request)
	if err != nil {
		return *new(shipping.DetailedCamundaProcessDTO), err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("%s", err)
		}
	}(response.Body)

	if response.StatusCode == http.StatusNotFound {
		return *new(shipping.DetailedCamundaProcessDTO), errors.New("process ID not found")
	}

	dao, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return *new(shipping.DetailedCamundaProcessDTO), err
	}

	dto, err := r.mapper.JsonBodyToDetailedProcessDTO(dao)

	if err != nil {
		return *new(shipping.DetailedCamundaProcessDTO), err
	}
	return dto, nil
}
