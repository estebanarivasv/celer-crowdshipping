package controllers

import (
	"github.com/gin-gonic/gin"
)

// InDTO Generic interface that makes implement a generic method available for all the Controllers
type InDTO interface {
	interface{}
}

// ShouldBindDTO Generic method that maps the JSON body to the generic InputDTO defined below
func ShouldBindDTO[T InDTO](c *gin.Context, DTO T) (boundDTO T, e error) {
	err := c.ShouldBindJSON(&DTO)
	if err != nil {
		return nil, err
	}
	return DTO, nil
}
