package repositories

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"gorm.io/gorm"
)

type RouteRepository struct {
	db *gorm.DB
}

// NewRouteRepository Returns a new instance
func NewRouteRepository() *RouteRepository {
	var db = config.ConnectToDb()
	return &RouteRepository{db: db}
}

func (r *RouteRepository) Create(dao models.RouteDetail) (models.RouteDetail, error) {
	err := r.db.Create(&dao).Error
	if err != nil {
		return *new(models.RouteDetail), err
	}

	return r.FindOneById(dao.ID)
}

// FindOneById Get one route by ID from the database
func (r *RouteRepository) FindOneById(id int) (models.RouteDetail, error) {
	var packModel models.RouteDetail

	err := r.db.Model(models.RouteDetail{}).Where("id = ?", id).Take(&packModel).Error
	if err != nil {
		return *new(models.RouteDetail), err
	}

	return packModel, nil
}

// DeleteById Delete a route by ID from the database
func (r *RouteRepository) DeleteById(id int) error {

	err := r.db.Delete(&models.RouteDetail{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
