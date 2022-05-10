package services

import (
	"fmt"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
	"log"
)

// Todo: make all functions available with a generic interface

var repository = repositories.NewShippingRepository()

func CreateShipping(shipping *models.Shipping) dtos.Response {

	response := repository.SaveOne(shipping)
	if response.Error != nil {
		return dtos.Response{Success: false, Error: response.Error.Error()}
	}

	return dtos.Response{Success: true, Data: response.Output.(*models.Shipping)}
}

func FindAllShippings() dtos.Response {

	response := repository.FindAll()
	if response.Error != nil {
		return dtos.Response{Success: false, Error: response.Error.Error()}
	}
	return dtos.Response{Success: true, Data: response.Output.([]*models.Shipping)}
}

func FindShippingById(id int) dtos.Response {

	response := repository.FindOneById(id)
	if response.Error != nil {
		return dtos.Response{Success: false, Error: response.Error.Error()}
	}

	log.Printf(fmt.Sprintf("%v", response.Output))

	return dtos.Response{Success: true, Data: response.Output.(*models.Shipping)}
}

func UpdateShippingById(id int, newShipping *models.Shipping) dtos.Response {

	existingShippingDao := FindShippingById(id)
	if !existingShippingDao.Success {
		// Returns success as "false" with an error
		return existingShippingDao
	}
	shipping := existingShippingDao.Data.(*models.Shipping)

	// Todo: fix mapping
	shipping.Image = newShipping.Image
	shipping.Details = newShipping.Details
	shipping.StartAddr = newShipping.StartAddr
	shipping.DestinationAddr = newShipping.DestinationAddr

	response := repository.SaveOne(shipping)
	if response.Error != nil {
		return dtos.Response{Success: false, Error: response.Error.Error()}
	}

	return dtos.Response{Success: true, Data: response.Output.(*models.Shipping)}
}

func DeleteShippingById(id int) dtos.Response {

	response := repository.DeleteOneById(id)
	if response.Error != nil {
		return dtos.Response{Success: false, Error: response.Error.Error()}
	}

	return dtos.Response{Success: true}
}
