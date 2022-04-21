package repositories

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"gorm.io/gorm"
)

type ShippingRepository struct {
	db *gorm.DB
}

// Todo: check if it works
var db = config.ConnectToDb()

// NewShippingRepository Returns a new instance
func NewShippingRepository() *ShippingRepository {
	return &ShippingRepository{db: db}
}

// SaveOne Store a new shipping in the database
func (r *ShippingRepository) SaveOne(shipping *models.Shipping) QueryResponse {
	err := r.db.Save(shipping).Error
	if err != nil {
		return QueryResponse{Error: err}
	}
	return QueryResponse{Output: shipping}
}

// FindAll Get shipping from the database
func (r *ShippingRepository) FindAll() QueryResponse {
	var shippings []*models.Shipping

	err := r.db.Find(&shippings).Error
	if err != nil {
		return QueryResponse{Error: err}
	}

	return QueryResponse{Output: shippings}
}

// FindOneById Get one shipping by ID from the database
func (r *ShippingRepository) FindOneById(id int) QueryResponse {
	var shipping *models.Shipping

	err := r.db.Where(&models.Shipping{ID: id}).Take(&shipping).Error
	if err != nil {
		return QueryResponse{Error: err}
	}

	return QueryResponse{Output: shipping}
}

// DeleteOneById Delete a shipping by ID from the database
func (r *ShippingRepository) DeleteOneById(id int) QueryResponse {

	err := r.db.Delete(&models.Shipping{ID: id}).Error
	if err != nil {
		return QueryResponse{Error: err}
	}

	return QueryResponse{Output: nil}
}
