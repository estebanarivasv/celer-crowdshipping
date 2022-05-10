package shipping

type DetailedShippingProc struct {
	Id                     string                 `json:"id"`
	DefinitionId           string                 `json:"processDefinitionId"`
	Name                   string                 `json:"name"`
	ActivityId             string                 `json:"activityId"`
	ActivityName           string                 `json:"activityName"`
	ChildActivityInstances []DetailedShippingProc `json:"childActivityInstances"`
	Incidents              []interface{}          `json:"incidents"`
}
