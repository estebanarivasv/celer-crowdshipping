package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"regexp"
)

// Directory where .env is located
const projectDirName = "backend-golang"

func LoadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()

	// Obtain the route where the .env file is, based in currentWorkDirectory
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	// Load env variables
	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		fmt.Println("Error loading .env file" + " - path: " + string(rootPath))
	}
}
