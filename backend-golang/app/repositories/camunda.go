package repositories

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	shipping "github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/utils/camunda"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
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

	var data any
	err = json.Unmarshal(dao, &data)
	if err != nil {
		fmt.Println("error: ", err)
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
		return shipping.DetailedCamundaProcessDTO{}, err
	}

	response, err := r.client.Do(request)
	if err != nil {
		return shipping.DetailedCamundaProcessDTO{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error: ", err)
		}
	}(response.Body)

	if response.StatusCode == http.StatusNotFound {
		return shipping.DetailedCamundaProcessDTO{}, errors.New("process ID not found")
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

func (r *CamundaRepository) DeployBPMNProcess(httpMethod string, url string, filePath string, deploymentName string) (shipping.DetailedCamundaProcessDTO, error) {

	enableDuplicateFiltering := "true"
	deployChangedOnly := "false"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file: ", err)
		return *new(shipping.DetailedCamundaProcessDTO), err
	}
	defer file.Close()

	// Create a buffer to store the request body
	body := &bytes.Buffer{}

	// Create a multipart form writer to write the request
	writer := multipart.NewWriter(body)

	// Add the file to the multipart form
	fileWriter, err := writer.CreateFormFile("*", file.Name())
	if err != nil {
		fmt.Println("Error adding the BPMN file to the multipart-form: ", err)
		return *new(shipping.DetailedCamundaProcessDTO), err
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Println("Error copying the file to the form: ", err)
		return *new(shipping.DetailedCamundaProcessDTO), err

	}

	// Add other fields to multipart form
	writer.WriteField("deployment-name", deploymentName)
	writer.WriteField("enable-duplicate-filtering", enableDuplicateFiltering)
	writer.WriteField("deploy-changed-only", deployChangedOnly)

	// Close multipart form
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing the multi-part form: ", err)
		return *new(shipping.DetailedCamundaProcessDTO), err
	}

	request, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		fmt.Println("Error creating the HTTP request: ", err)
		return *new(shipping.DetailedCamundaProcessDTO), err
	}

	// Set request header
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Make the request
	response, err := r.client.Do(request)
	if err != nil {
		fmt.Println("Error making the HTTP request:", err)
		return *new(shipping.DetailedCamundaProcessDTO), err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("%s", err)
		}
	}(response.Body)

	if response.StatusCode == http.StatusBadRequest {
		return *new(shipping.DetailedCamundaProcessDTO), errors.New("one of the bpmn resources cannot be parsed")
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

func (r *CamundaRepository) BasicGetRequestForPing() (bool, error) {

	url := camunda.GetURLToPingCamundaEngine()

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}

	response, err := r.client.Do(request)
	if err != nil {
		return false, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("%s", err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected response status code: %d\n", response.StatusCode)
		return false, nil
	}

	// Read the response body
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return false, fmt.Errorf("error reading the response body: %v", err)
	}

	return true, nil
}
