package config

import (
	"github.com/swaggo/swag/example/basic/docs"
	"os"
)

func LoadSwaggerConfig() {

	LoadEnv()

	docs.SwaggerInfo.Title = "Celer crowdshipping API"
	docs.SwaggerInfo.Description = "Crowdshipping in the whole Argentine republic."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "www.google.com" - TODO deploy to heroku
	docs.SwaggerInfo.BasePath = os.Getenv("BASE_PATH")
}
