package services

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/mappers"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/repositories"
)

type OfferService struct {
	offerRepo    *repositories.OfferRepository
	shippingRepo *repositories.ShippingRepository
	userRepo     *repositories.UserRepository
	mapper       *mappers.OfferMapper
}

// NewOfferService Returns a new instance
func NewOfferService() *OfferService {

	var shippingRepository = repositories.NewShippingRepository()
	var offerRepository = repositories.NewOfferRepository()
	var userRepository = repositories.NewUserRepository()
	var offerMapper = mappers.NewOfferMapper()

	return &OfferService{
		offerRepo:    offerRepository,
		shippingRepo: shippingRepository,
		userRepo:     userRepository,
		mapper:       offerMapper,
	}
}

func (s *OfferService) FindOfferById(id int) dtos.Response {
	query, err := s.offerRepo.FindOneById(&id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	dto := s.mapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: dto}

}

func (s *OfferService) CreateOffer(offerDTO *entities.OfferInDTO, distributorIDFromContext int) dtos.Response {

	// Get shipping
	shippingDao, err := s.shippingRepo.FindOneById(offerDTO.ShippingID)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Verify Shipping has selected_offer_id null
	if shippingDao.SelectedOfferID != nil {
		return dtos.Response{Success: false, Error: "this shipping id does not refer to a shipping request"}
	}

	// Convert the dto to an entity
	offerModel := s.mapper.FromDTO(offerDTO)
	offerModel.DistributorID = distributorIDFromContext // Set distributor from the context

	query, err := s.offerRepo.Create(offerModel)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Mapping into a dto
	dto := s.mapper.ToDTO(&query)

	return dtos.Response{Success: true, Data: dto}
}

func (s *OfferService) FindOffersByRequestID(id int) dtos.Response {

	// Get shipping
	shippingDao, err := s.shippingRepo.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Verify Shipping has selected_offer_id null
	if shippingDao.SelectedOfferID != nil {
		return dtos.Response{Success: false, Error: "this shipping id does not refer to a shipping request"}
	}

	var dtosArr []interface{}

	conditions := make(map[string]interface{})
	conditions["shipping_id"] = id

	offers, err := s.offerRepo.FindAllRequestOffers(conditions)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Convert and append all bpmn-models into dtos
	for _, offer := range offers {
		dtosArr = append(dtosArr, s.mapper.ToDTO(&offer))
	}

	return dtos.Response{Success: true, Data: dtosArr}

}

func (s *OfferService) FindOffersByShippingID(id int) dtos.Response {

	var dtosArr []interface{}

	// Get shipping
	shippingDao, err := s.shippingRepo.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	conditions := make(map[string]interface{})
	conditions["shipping_id"] = shippingDao.ID

	offers, err := s.offerRepo.FindAllRequestOffers(conditions)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Convert and append all bpmn-models into dtos
	for _, offer := range offers {
		dtosArr = append(dtosArr, s.mapper.ToDTO(&offer))
	}

	return dtos.Response{Success: true, Data: dtosArr}

}

func (s *OfferService) DeleteOfferById(id int) dtos.Response {
	err := s.offerRepo.DeleteById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	return dtos.Response{Success: true}
}

func (s *OfferService) FindOffersByDistributorID(id int) dtos.Response {

	var dtosArr []interface{}

	// Verify user exists
	user, err := s.userRepo.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	conditions := make(map[string]interface{})
	conditions["distributor_id"] = user.ID

	offers, err := s.offerRepo.FindAllRequestOffers(conditions)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Convert and append all bpmn-models into dtos
	for _, offer := range offers {
		dtosArr = append(dtosArr, s.mapper.ToDTO(&offer))
	}

	return dtos.Response{Success: true, Data: dtosArr}

}

func (s *OfferService) FindSelectedOfferByShippingId(id int) dtos.Response {

	// Verify shipping exists
	shipping, err := s.shippingRepo.FindOneById(id)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	// Verify selected offer is not null
	if shipping.SelectedOfferID == nil {
		return dtos.Response{Success: false, Error: "there is not a selected offer in this shipping instance."}
	}

	dao, err := s.offerRepo.FindOneById(shipping.SelectedOfferID)
	if err != nil {
		return dtos.Response{Success: false, Error: err.Error()}
	}

	dto := s.mapper.ToDetailedDTO(&dao)

	return dtos.Response{Success: true, Data: dto}
}
