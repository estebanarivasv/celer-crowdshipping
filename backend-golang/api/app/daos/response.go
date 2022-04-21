package daos

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"message"`
	Data    interface{} `json:"data"`
}
