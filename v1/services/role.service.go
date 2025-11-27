package services

import (
	"errors"

	"github.com/arhief32/emp-be/v1/entities"
	"github.com/arhief32/emp-be/v1/models"
	"github.com/arhief32/emp-be/v1/repositories"
)

type RoleService struct {
	Repo *repositories.RoleRepository
}

func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{Repo: repo}
}

func (s *RoleService) Create(req entities.CreateRoleRequest) (*models.Role, error) {
	role := models.Role{
		Role:     req.Role,
		OrgDesc:  req.OrgDesc,
		JobTitle: req.JobTitle,
	}

	err := s.Repo.Create(&role)
	return &role, err
}

func (s *RoleService) GetAll() ([]models.Role, error) {
	return s.Repo.FindAll()
}

func (s *RoleService) GetByID(id uint) (*models.Role, error) {
	role, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, errors.New("role tidak ditemukan")
	}
	return role, nil
}

func (s *RoleService) Update(id uint, req entities.UpdateRoleRequest) (*models.Role, error) {
	role, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, errors.New("role tidak ditemukan")
	}

	role.Role = req.Role
	role.OrgDesc = req.OrgDesc
	role.JobTitle = req.JobTitle

	err = s.Repo.Update(role)
	return role, err
}

func (s *RoleService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
