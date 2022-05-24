package repositories

import (
	"encoding/json"
	"fmt"
	shipping "github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}
var mapper = mappers.CamundaMapper{}

func SendNewInstanceRequest(httpMethod string, url string, body io.Reader) (*shipping.BasicCamundaProcessDTO, error) {

	// Define the request
	request, err := http.NewRequest(httpMethod, url, body)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	// Make a request and return a response
	response, err := client.Do(request)
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
		log.Fatal(err)
		return nil, err
	}

	dto, err := mapper.JsonBodyToCamundaProcessDTO(dao)
	if err != nil {
		return nil, err
	}
	return dto, nil

}

func SendMessageToProcessRequest(httpMethod string, url string, body io.Reader) error {

	request, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
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
		return http.ErrBodyNotAllowed
	}
	return nil
}
