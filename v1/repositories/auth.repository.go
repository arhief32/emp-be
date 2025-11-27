package repositories

import (
	"github.com/arhief32/emp-be/v1/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByUsername(username string) (*models.User, error)
	FindByID(id int) (*models.User, error)
	Create(user *models.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) FindByID(id int) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}
