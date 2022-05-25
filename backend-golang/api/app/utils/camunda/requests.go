package camunda

import (
	"bytes"
	"encoding/json"
)

func PayloadToJsonBody(payload any) (*bytes.Buffer, error) {

	encodedPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	reader := GetBufferFromEncodedJson(encodedPayload)

	return reader, nil
}

func GetBufferFromEncodedJson(json []byte) *bytes.Buffer {
	return bytes.NewBuffer(json)
}
