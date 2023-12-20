package config

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/docs"
	"os"
)

func LoadSwaggerConfig() {

	LoadEnv()

	docs.SwaggerInfo.Title = "Celer crowdshipping API"
	docs.SwaggerInfo.Description = "Crowdshipping in the whole Argentine Republic."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "www.google.com" - TODO deploy to heroku
	docs.SwaggerInfo.BasePath = os.Getenv("BASE_PATH")
}
