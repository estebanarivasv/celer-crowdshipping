package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	// Create a get method associated with a function
	router.GET("/", func(context *gin.Context) {
		// The function returns a json
		context.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.Run(":5000")
}
