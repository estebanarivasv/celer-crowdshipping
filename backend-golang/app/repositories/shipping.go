package repositories

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ShippingRepository struct {
	db *gorm.DB
}

// NewShippingRepository Returns a new instance
func NewShippingRepository() *ShippingRepository {
	var db = config.ConnectToDb()
	return &ShippingRepository{db: db}
}

// Save Store a new entity in the database
func (r *ShippingRepository) Save(shipping *models.Shipping) (*models.Shipping, error) {
	err := r.db.Save(&shipping).Error
	if err != nil {
		return *new(*models.Shipping), err
	}
	return shipping, nil
}

func (r *ShippingRepository) Create(dao models.Shipping) (models.Shipping, error) {
	err := r.db.Create(&dao).Error
	if err != nil {
		return *new(models.Shipping), err
	}

	return r.FindOneById(dao.ID)
}

// FindAll Get shipping from the database
func (r *ShippingRepository) FindAll() ([]models.Shipping, error) {
	var shippings []models.Shipping

	err := r.db.Preload(clause.Associations).Find(&shippings).Error
	if err != nil {
		return *new([]models.Shipping), err
	}

	return shippings, nil
}

// FindFilteredShippings Get shipping requests from the database
func (r *ShippingRepository) FindFilteredShippings(filter map[string]interface{}) ([]models.Shipping, error) {
	var shippings []models.Shipping

	err := r.db.Preload(clause.Associations).Where(filter).Find(&shippings).Error
	if err != nil {
		return *new([]models.Shipping), err
	}

	return shippings, nil
}

// FindOneById Get one shipping by ID from the database
func (r *ShippingRepository) FindOneById(id int) (models.Shipping, error) {
	var shipping models.Shipping

	err := r.db.Model(models.Shipping{}).Preload(clause.Associations).Where("id = ?", id).Take(&shipping).Error
	if err != nil {
		// TODO HANDLE EMPTY
		return *new(models.Shipping), err
	}

	return shipping, nil
}

// DeleteById Delete a shipping by ID from the database
func (r *ShippingRepository) DeleteById(id int) error {

	err := r.db.Delete(&models.Shipping{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateById Update an entity from the database with a body and an ID
func (r *ShippingRepository) UpdateById(id int, dto interface{}) (models.Shipping, error) {

	err := r.db.Model(models.Shipping{}).Where("id = ?", id).Updates(dto).Error
	if err != nil {
		return *new(models.Shipping), err
	}

	editedModel, err := r.FindOneById(id)
	if err != nil {
		return *new(models.Shipping), err
	}
	editedModel.UpdatedAt = time.Now()

	_, err = r.Save(&editedModel)
	if err != nil {
		return *new(models.Shipping), err
	}

	return editedModel, nil
}

// FindCurrentShippings Get one shipping by ID from the database
func (r *ShippingRepository) FindCurrentShippings() ([]models.Shipping, error) {
	var shipping []models.Shipping

	err := r.db.Model(models.Shipping{}).Preload(clause.Associations).Where("selected_offer_id IS NOT NULL ").Find(&shipping).Error
	if err != nil {
		return *new([]models.Shipping), err
	}

	return shipping, nil
}
