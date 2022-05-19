package shipping

type MessageToProc struct {
	MessageName       string `json:"messageName"`
	ProcessInstanceId string `json:"processInstanceId"`
}
