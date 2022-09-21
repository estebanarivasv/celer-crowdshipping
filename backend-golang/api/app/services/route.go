package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
)

type RouteService struct {
	routeRepo    *repositories.RouteRepository
	shippingRepo *repositories.ShippingRepository
	mapper       *mappers.RouteMapper
}

// NewRouteService Returns a new instance
func NewRouteService() *RouteService {

	var routeRepository = repositories.NewRouteRepository()
	var shippingRepository = repositories.NewShippingRepository()
	var routeMapper = mappers.NewRouteMapper()

	return &RouteService{
		routeRepo:    routeRepository,
		shippingRepo: shippingRepository,
		mapper:       routeMapper,
	}
}

func (s *RouteService) AddNewShippingCoordinates(shippingId int, dto entities.RouteDetailInDTO) dtos.Response {

	model := s.mapper.FromDTO(&dto)
	model.ShippingID = shippingId

	query, err := s.routeRepo.Create(model)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	responseDto := s.mapper.ToDTO(&query)
	return dtos.Response{Success: true, Data: responseDto}
}

func (s *RouteService) FindRouteByShippingID(id int) dtos.Response {

	// Get shipping
	shippingDao, err := s.shippingRepo.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	conditions := make(map[string]interface{})
	conditions["shipping_id"] = shippingDao.ID

	routeDetails, err := s.routeRepo.FindFilteredRoutes(conditions)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Map all models into a dto
	var routeArr []interface{}
	for _, detail := range routeDetails {
		routeArr = append(routeArr, s.mapper.ToDTO(&detail))
	}

	return dtos.Response{Success: true, Data: routeArr}
}
