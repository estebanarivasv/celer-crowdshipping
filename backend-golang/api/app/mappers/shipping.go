package mappers

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
)

type ShippingMapper struct {
}

// FromDTO Return a model with the dto that comes from the service
func (m ShippingMapper) FromDTO(dto *entities.ShippingInDTO) models.Shipping {

	return models.Shipping{
		Details:            dto.Details,
		PackageID:          dto.PackageID,
		SelectedOfferID:    dto.SelectedOfferID,
		SenderID:           dto.SenderID,
		RecipientID:        dto.RecipientID,
		OriginAddress:      dto.OriginAddress,
		DestinationAddress: dto.DestinationAddress,
	}

}

// ToDTO Return a dto with the model that comes from the database
func (m ShippingMapper) ToDTO(model *models.Shipping) interface{} {

	// Each coordinate is mapped into a dto and appended to a slice
	// var routeArr []interface{}
	// for _, r := range model.Route {
	// 	routeArr = append(routeArr,
	// 		RouteMapper{}.ToDTO(&r))
	// }

	// Each coordinate is mapped into a dto and appended to a slice
	// var offersSlice []interface{}
	// for _, o := range model.Offers {
	// 	offersSlice = append(offersSlice,
	// 		OfferMapper{}.ToDTO(&o))
	// }

	return entities.ShippingOutDTO{
		ID:                 model.ID,
		Details:            model.Details,
		Package:            PackageMapper{}.ToDTO(&model.Package),
		OriginAddress:      model.OriginAddress,
		DestinationAddress: model.DestinationAddress,
		ProcessID:          model.ProcessID,
		SelectedOfferID:    model.SelectedOfferID,
		Sender:             UserMapper{}.ToDTO(&model.Sender),
		Recipient:          UserMapper{}.ToDTO(&model.Recipient),
		PickedUpAt:         model.PickedUpAt,
		DeliveredAt:        model.DeliveredAt,
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
		DeletedAt:          model.DeletedAt.Time,
		// Offers:             offersSlice,
		// Route:              routeArr,
	}

}
