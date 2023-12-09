package repositories

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository Returns a new instance
func NewUserRepository() *UserRepository {
	var db = config.ConnectToDb()
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(dao models.User) (models.User, error) {
	err := r.db.Create(&dao).Error
	if err != nil {
		return *new(models.User), err
	}

	return r.FindOneById(dao.ID)
}

// FindOneById Get one user by ID from the database
func (r *UserRepository) FindOneById(id int) (models.User, error) {
	var packModel models.User

	err := r.db.Model(models.User{}).Where("id = ?", id).Take(&packModel).Error
	if err != nil {
		return *new(models.User), err
	}

	return packModel, nil
}

// DeleteById Delete a user by ID from the database
func (r *UserRepository) DeleteById(id int) error {

	err := r.db.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
