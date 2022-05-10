package mappers

import (
	"encoding/json"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/shipping"
	"log"
)

func JsonToShippingStruct(jsonResponse []byte) (*shipping.BasicShippingProc, error) {
	data := new(shipping.BasicShippingProc)
	err := json.Unmarshal(jsonResponse, &data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return data, nil
}
