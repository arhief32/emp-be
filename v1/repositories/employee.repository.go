package repositories

import (
	"time"

	"gorm.io/gorm"

	"github.com/yourusername/pegawai-api/v1/models"
)

type EmployeeRepository interface {
	FindAll() ([]models.Employee, error)
	Create(models.Employee) (models.Employee, error)
	FindByID(uint) (models.Employee, error)
	Update(models.Employee) error
	Delete(uint) error
	CountByCreatedAtRange(start, end time.Time) (int, error)
	FindByCreatedAtRange(start, end time.Time) ([]models.Employee, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) FindAll() ([]models.Employee, error) {
	var list []models.Employee
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *employeeRepository) Create(e models.Employee) (models.Employee, error) {
	if err := r.db.Create(&e).Error; err != nil {
		return models.Employee{}, err
	}
	return e, nil
}

func (r *employeeRepository) FindByID(id uint) (models.Employee, error) {
	var e models.Employee
	if err := r.db.First(&e, id).Error; err != nil {
		return models.Employee{}, err
	}
	return e, nil
}

func (r *employeeRepository) Update(e models.Employee) error {
	return r.db.Save(&e).Error
}

func (r *employeeRepository) Delete(id uint) error {
	return r.db.Delete(&models.Employee{}, id).Error
}

func (r *employeeRepository) CountByCreatedAtRange(start, end time.Time) (int, error) {
	var count int64
	if err := r.db.Model(&models.Employee{}).Where("created_at BETWEEN ? AND ?", start, end).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *employeeRepository) FindByCreatedAtRange(start, end time.Time) ([]models.Employee, error) {
	var list []models.Employee
	if err := r.db.Where("created_at BETWEEN ? AND ?", start, end).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
