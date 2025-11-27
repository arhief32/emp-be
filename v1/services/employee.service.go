package services

import (
	"errors"
	"time"

	"github.com/arhief32/emp-be/v1/entities"
	"github.com/arhief32/emp-be/v1/models"
	"github.com/arhief32/emp-be/v1/repositories"
)

type EmployeeService interface {
	GetAll() ([]models.Employee, error)
	Create(req entities.EmployeeCreateRequest) (models.Employee, error)
	GetByID(id uint) (models.Employee, error)
	Update(id uint, req entities.EmployeeUpdateRequest) error
	Delete(id uint) error
}

type employeeService struct {
	repo repositories.EmployeeRepository
}

func NewEmployeeService(r repositories.EmployeeRepository) EmployeeService {
	return &employeeService{repo: r}
}

func (s *employeeService) GetAll() ([]models.Employee, error) {
	return s.repo.FindAll()
}

func (s *employeeService) Create(req entities.EmployeeCreateRequest) (models.Employee, error) {
	// basic validation
	if req.Nip == "" || req.Nama == "" {
		return models.Employee{}, errors.New("nama and nip required")
	}
	now := time.Now()
	e := models.Employee{
		Nama:      req.Nama,
		Nip:       req.Nip,
		Jabatan:   req.Jabatan,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return s.repo.Create(e)
}

func (s *employeeService) GetByID(id uint) (models.Employee, error) {
	return s.repo.FindByID(id)
}

func (s *employeeService) Update(id uint, req entities.EmployeeUpdateRequest) error {
	// fetch existing
	e, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if req.Nama != "" {
		e.Nama = req.Nama
	}
	if req.Nip != "" {
		e.Nip = req.Nip
	}
	if req.Jabatan != "" {
		e.Jabatan = req.Jabatan
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

func (s *employeeService) Delete(id uint) error {
	return s.repo.Delete(id)
}
