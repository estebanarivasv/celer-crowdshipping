package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
)

var routeRepository = repositories.NewRouteRepository()
var routeMapper = mappers.RouteMapper{}

func AddNewShippingCoordinates(shippingId int, dto entities.RouteDetailInDTO) dtos.Response {

	model := routeMapper.FromDTO(&dto)
	model.ShippingID = shippingId

	query, err := routeRepository.Create(model)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	responseDto := routeMapper.ToDTO(&query)
	return dtos.Response{Success: true, Data: responseDto}
}

func FindRouteByShippingID(id int) dtos.Response {

	// Get shipping
	shippingDao, err := shippingRepository.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	conditions := make(map[string]interface{})
	conditions["shipping_id"] = shippingDao.ID

	routeDetails, err := routeRepository.FindFilteredRoutes(conditions)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Map all models into a dto
	var routeArr []interface{}
	for _, detail := range routeDetails {
		routeArr = append(routeArr, routeMapper.ToDTO(&detail))
	}

	return dtos.Response{Success: true, Data: routeArr}
}
