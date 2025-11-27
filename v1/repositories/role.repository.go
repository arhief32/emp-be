package repositories

import (
	"github.com/arhief32/emp-be/v1/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (r *RoleRepository) Create(role *models.Role) error {
	return r.DB.Create(role).Error
}

func (r *RoleRepository) FindAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.DB.Find(&roles).Error
	return roles, err
}

func (r *RoleRepository) FindByID(id uint) (*models.Role, error) {
	var role models.Role
	err := r.DB.First(&role, id).Error
	return &role, err
}

func (r *RoleRepository) Update(role *models.Role) error {
	return r.DB.Save(role).Error
}

func (r *RoleRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Role{}, id).Error
}
