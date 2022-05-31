package controllers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// InDTO Generic interface that makes implement a generic method available for all the Controllers
type InDTO interface {
	interface{}
}

// ShouldBindDTO Generic method that maps the JSON body to the generic InputDTO defined below
func ShouldBindDTO[T InDTO](c *gin.Context, dto T) (boundDTO T) {
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			dtos.Response{Error: "the server will not process the request due to a malformed " +
				"request syntax, an invalid request message framing, or a deceptive request routing"})
		return
	}
	return dto
}

func ConvertParamToInt(param string) (int, error) {
	p, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return p, nil
}

func ContainsValidSenderMsg(v string) bool {
	elems := []string{"OfferSelected", "PackageOk"}
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func ContainsValidDistributorMsg(v string) bool {
	elems := []string{"PackageInTransit", "DeliveredToRecipient"}
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
