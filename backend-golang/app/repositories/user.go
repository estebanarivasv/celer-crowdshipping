package repositories

import (
	"github.com/estebanarivasv/Celer/backend-golang/api/app/config"
	"github.com/estebanarivasv/Celer/backend-golang/api/app/models"
	"golang.org/x/crypto/bcrypt"
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

func (r *UserRepository) Create(userDao models.User) (models.User, error) {

	// Generar el hash de la contraseña antes de almacenarla
	hashedPassword, err := r.GeneratePwdHash(userDao.Password)
	if err != nil {
		return models.User{}, err
	}

	// Almacenar el hash de la contraseña en lugar de la contraseña en sí misma
	userDao.Password = string(hashedPassword)

	err = r.db.Create(&userDao).Error
	if err != nil {
		return *new(models.User), err
	}

	user, err := r.FindOneById(userDao.ID)
	if err != nil {
		return models.User{}, err
	}

	user.Password = ""

	return user, err
}

// FindOneById Get one user by ID from the database
func (r *UserRepository) FindOneById(id int) (models.User, error) {
	var user models.User

	err := r.db.Model(models.User{}).Where("id = ?", id).Take(&user).Error
	if err != nil {
		return *new(models.User), err
	}

	return user, nil
}

func (r *UserRepository) GeneratePwdHash(pwd string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hashedPassword, err
}

// DeleteById Delete a user by ID from the database
func (r *UserRepository) DeleteById(id int) error {

	err := r.db.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

// FindOneByUsername Get one user by ID from the database
func (r *UserRepository) FindOneByUsername(username string) (models.User, error) {
	var user models.User

	err := r.db.Model(models.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return *new(models.User), err
	}

	return user, nil
}
