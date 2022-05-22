package repositories

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/dtos/entities"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"gorm.io/gorm"
)

type ShippingRepository struct {
	db *gorm.DB
}

// TODO CHANGE TO CONTEXT -- var db = context.Background()
var db = config.ConnectToDb()

// NewShippingRepository Returns a new instance
func NewShippingRepository() *ShippingRepository {
	return &ShippingRepository{db: db}
}

// Save Store a new entity in the database
func (r *ShippingRepository) Save(shipping models.Shipping) (models.Shipping, error) {
	err := r.db.Save(shipping).Error
	if err != nil {
		return *new(models.Shipping), err
	}
	return shipping, nil
}

func (r *ShippingRepository) Create(dao models.Shipping) (models.Shipping, error) {

	err := r.db.Create(&dao).Error
	if err != nil {
		// TODO HANDLE EMPTY
		return *new(models.Shipping), err
	}

	return r.FindOneById(dao.ID)
}

// FindAll Get shipping from the database
func (r *ShippingRepository) FindAll() ([]models.Shipping, error) {
	var shippings []models.Shipping

	err := r.db.Find(&shippings).Error
	if err != nil {
		// TODO HANDLE EMPTY
		return *new([]models.Shipping), err
	}

	return shippings, nil
}

// FindOneById Get one shipping by ID from the database
func (r *ShippingRepository) FindOneById(id int) (models.Shipping, error) {
	var shipping models.Shipping

	err := r.db.Where(&models.Shipping{ID: id}).Take(&shipping).Error
	if err != nil {
		// TODO HANDLE EMPTY
		return *new(models.Shipping), err
	}

	return shipping, nil
}

// DeleteById Delete a shipping by ID from the database
func (r *ShippingRepository) DeleteById(id int) error {

	err := r.db.Delete(&models.Shipping{ID: id}).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateById Update an entity from the database with a body and an ID
func (r *ShippingRepository) UpdateById(id int, dto *entities.ShippingInDTO) (models.Shipping, error) {

	existentModel, err := r.FindOneById(id)
	if err != nil {
		// TODO HANDLE EMPTY
		return *new(models.Shipping), err
	}

	// TODO SUPUESTAMENTE SE ACTUALIZA
	r.db.Model(&existentModel).Select("*").Updates(dto)

	return r.FindOneById(existentModel.ID)
}
