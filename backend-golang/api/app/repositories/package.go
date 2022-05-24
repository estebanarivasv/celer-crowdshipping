package repositories

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"gorm.io/gorm"
)

type PackageRepository struct {
	db *gorm.DB
}

// NewPackageRepository Returns a new instance
func NewPackageRepository() *PackageRepository {
	var db = config.ConnectToDb()
	return &PackageRepository{db: db}
}

func (r *PackageRepository) Create(dao models.Package) (models.Package, error) {
	err := r.db.Create(&dao).Error
	if err != nil {
		return *new(models.Package), err
	}

	return r.FindOneById(dao.ID)
}

// FindOneById Get one package by ID from the database
func (r *PackageRepository) FindOneById(id int) (models.Package, error) {
	var packModel models.Package

	err := r.db.Model(models.Package{}).Where("id = ?", id).Take(&packModel).Error
	if err != nil {
		return *new(models.Package), err
	}

	return packModel, nil
}

// DeleteById Delete a package by ID from the database
func (r *PackageRepository) DeleteById(id int) error {

	err := r.db.Delete(&models.Package{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
