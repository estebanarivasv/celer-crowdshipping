package shipping

type BasicShippingProc struct {
	Id           string `json:"id"`
	DefinitionId string `json:"definitionId"`
	Ended        bool   `json:"ended"`
	Suspended    bool   `json:"suspended"`
	Links        []struct {
		Method string `json:"method"`
		Href   string `json:"href"`
		Rel    string `json:"rel"`
	} `json:"links"`
}
