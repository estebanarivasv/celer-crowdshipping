package repositories

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OfferRepository struct {
	db *gorm.DB
}

// NewOfferRepository Returns a new instance
func NewOfferRepository() *OfferRepository {
	var db = config.ConnectToDb()
	return &OfferRepository{db: db}
}

func (r *OfferRepository) Create(dao models.Offer) (models.Offer, error) {
	err := r.db.Create(&dao).Error
	if err != nil {
		return *new(models.Offer), err
	}

	return r.FindOneById(dao.ID)
}

// FindOneById Get one offer by ID from the database
func (r *OfferRepository) FindOneById(id int) (models.Offer, error) {
	var model models.Offer

	err := r.db.First(&model, id).Error
	if err != nil {
		return *new(models.Offer), err
	}

	return model, nil
}

// DeleteById Delete a offer by ID from the database
func (r *OfferRepository) DeleteById(id int) error {

	err := r.db.Delete(&models.Offer{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAllRequestOffers Get offers associated to requests from the database
func (r *OfferRepository) FindAllRequestOffers(conditions map[string]interface{}) ([]models.Offer, error) {
	var offers []models.Offer

	err := r.db.Preload(clause.Associations).Where(conditions).Find(&offers).Error
	if err != nil {
		return *new([]models.Offer), err
	}

	return offers, nil
}
