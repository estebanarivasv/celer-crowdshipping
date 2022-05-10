package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/shipping"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

const baseUrl = "http://localhost:8080/engine-rest"

func SendMessage(msgName string, procId string) error {

	const url = baseUrl + "/message"
	payload := dtos.MessageToProc{
		MessageName:       msgName,
		ProcessInstanceId: procId,
	}
	encodedPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(encodedPayload))
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

// TODO COMMENT AND REFACTOR THIS, HANDLE ERRORS
func CreateShippingProcInstance() (*shipping.BasicShippingProc, error) {

	var url = baseUrl + `/process-definition/key/ShippingDelivery/start`
	emptyJson := bytes.NewBuffer([]byte(`{}`))

	request, err := http.NewRequest(http.MethodPost, url, emptyJson)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	shippingResp, err := mappers.JsonToShippingStruct(body)
	return shippingResp, nil
}
