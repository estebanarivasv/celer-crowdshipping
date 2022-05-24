package dtos

type BasicCamundaProcessDTO struct {
	ID           string `json:"id"`
	DefinitionId string `json:"definitionId"`
	Ended        bool   `json:"ended"`
	Suspended    bool   `json:"suspended"`
	Links        []struct {
		Method string `json:"method"`
		Href   string `json:"href"`
		Rel    string `json:"rel"`
	} `json:"links"`
}

type DetailedCamundaProcessDTO struct {
	ID                     string                      `json:"id"`
	DefinitionId           string                      `json:"processDefinitionId"`
	Name                   string                      `json:"name"`
	ActivityId             string                      `json:"activityId"`
	ActivityName           string                      `json:"activityName"`
	ChildActivityInstances []DetailedCamundaProcessDTO `json:"childActivityInstances"`
	Incidents              []interface{}               `json:"incidents"`
}

type MessageToProcessDTO struct {
	MessageName       string `json:"messageName"`
	ProcessInstanceId string `json:"processInstanceId"`
}
