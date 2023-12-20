package camunda

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"os"
)

func GetProcCurrentActivityURL(procID string) string {
	config.LoadEnv()
	CamundaBaseUrl := os.Getenv("CAMUNDA_BASE_URL")

	return CamundaBaseUrl + "/process-instance/" + procID + "/activity-instances"
}

func GetMessageProcURL() string {
	config.LoadEnv()
	CamundaBaseUrl := os.Getenv("CAMUNDA_BASE_URL")

	return CamundaBaseUrl + "/message"
}

func GetProcStateURL(procID string) string {
	config.LoadEnv()
	CamundaBaseUrl := os.Getenv("CAMUNDA_BASE_URL")
	return CamundaBaseUrl + "/history/process-instance/" + procID
}

func GetCreateShippingProcURL() string {

	config.LoadEnv()
	CamundaBaseUrl := os.Getenv("CAMUNDA_BASE_URL")
	ShippingKey := os.Getenv("CAMUNDA_SHIPPING_KEY")

	return CamundaBaseUrl + "/process-definition/key/" + ShippingKey + "/start"
}

func GetURLToPushBPMNDiagram() string {

	config.LoadEnv()
	CamundaBaseUrl := os.Getenv("CAMUNDA_BASE_URL")

	return CamundaBaseUrl + "/deployment/create"
}

func GetURLToPingCamundaEngine() string {

	config.LoadEnv()
	CamundaBaseUrl := os.Getenv("CAMUNDA_BASE_URL")

	return CamundaBaseUrl + "/engine/"
}
