package services

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/arhief32/emp-be/v1/entities"
	"github.com/arhief32/emp-be/v1/models"
	"github.com/arhief32/emp-be/v1/repositories"
)

type MerchantSubmissionService interface {
	Create(req map[string]interface{}, makerID uint) (*models.MerchantSubmission, error)
	Update(id uint, req map[string]interface{}, makerID uint) (*models.MerchantSubmission, error)
	Submit(id uint, makerID uint) error
	GetByID(id uint) (*models.MerchantSubmission, error)
	GetByMaker(makerID uint) ([]models.MerchantSubmission, error)
	ListPendingForChecker() ([]entities.MerchantSubmissionResponse, error)
	CheckerApprove(id uint, checkerID uint, notes string) error
	CheckerReject(id uint, checkerID uint, notes string) error
	SignerApprove(id uint, signerID uint) error
	SignerReject(id uint, signerID uint, notes string) error
}

type merchantSubmissionService struct {
	repo repositories.MerchantSubmissionRepository
}

func NewMerchantSubmissionService(r repositories.MerchantSubmissionRepository) MerchantSubmissionService {
	return &merchantSubmissionService{repo: r}
}

func (s *merchantSubmissionService) Create(req map[string]interface{}, makerID uint) (*models.MerchantSubmission, error) {
	// Map request -> model
	docJSON := []byte("{}")
	if d, ok := req["documents"]; ok {
		b, _ := json.Marshal(d)
		docJSON = b
	}

	sub := &models.MerchantSubmission{
		MerchantName:    toString(req["merchant_name"]),
		OwnerName:       toString(req["owner_name"]),
		Phone:           toString(req["phone"]),
		Email:           toString(req["email"]),
		Address:         toString(req["address"]),
		Category:        toString(req["category"]),
		NIB:             toString(req["nib"]),
		NPWP:            toString(req["npwp"]),
		YearEstablished: toInt(req["year_established"]),
		Employees:       toInt(req["employees"]),
		Documents:       docJSON,
		Status:          "draft",
		MakerID:         makerID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	if err := s.repo.Create(sub); err != nil {
		return nil, err
	}
	return sub, nil
}

func (s *merchantSubmissionService) Update(id uint, req map[string]interface{}, makerID uint) (*models.MerchantSubmission, error) {
	sub, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if sub.MakerID != makerID {
		return nil, errors.New("tidak punya akses untuk mengubah pengajuan ini")
	}
	if sub.Status != "draft" && sub.Status != "rejected" {
		return nil, errors.New("hanya pengajuan dengan status draft atau rejected yang bisa diubah")
	}
	// patch fields if present
	if v, ok := req["merchant_name"]; ok {
		sub.MerchantName = toString(v)
	}
	if v, ok := req["owner_name"]; ok {
		sub.OwnerName = toString(v)
	}
	if v, ok := req["phone"]; ok {
		sub.Phone = toString(v)
	}
	if v, ok := req["email"]; ok {
		sub.Email = toString(v)
	}
	if v, ok := req["address"]; ok {
		sub.Address = toString(v)
	}
	if v, ok := req["category"]; ok {
		sub.Category = toString(v)
	}
	if v, ok := req["nib"]; ok {
		sub.NIB = toString(v)
	}
	if v, ok := req["npwp"]; ok {
		sub.NPWP = toString(v)
	}
	if v, ok := req["year_established"]; ok {
		sub.YearEstablished = toInt(v)
	}
	if v, ok := req["employees"]; ok {
		sub.Employees = toInt(v)
	}
	if v, ok := req["documents"]; ok {
		b, _ := json.Marshal(v)
		sub.Documents = b
	}
	sub.UpdatedAt = time.Now()
	if err := s.repo.Update(sub); err != nil {
		return nil, err
	}
	return sub, nil
}

func (s *merchantSubmissionService) Submit(id uint, makerID uint) error {
	sub, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if sub.MakerID != makerID {
		return errors.New("tidak punya akses submit pengajuan ini")
	}
	if sub.Status != "draft" && sub.Status != "rejected" {
		return errors.New("hanya draft atau rejected yang bisa di-submit")
	}
	sub.Status = "submitted"
	sub.UpdatedAt = time.Now()
	return s.repo.Update(sub)
}

func (s *merchantSubmissionService) GetByID(id uint) (*models.MerchantSubmission, error) {
	return s.repo.FindByID(id)
}

func (s *merchantSubmissionService) GetByMaker(makerID uint) ([]models.MerchantSubmission, error) {
	return s.repo.FindByMaker(makerID)
}

func (s *merchantSubmissionService) ListPendingForChecker() ([]entities.MerchantSubmissionResponse, error) {
	data, err := s.repo.FindPendingForChecker()
	if err != nil {
		return nil, err
	}

	// Convert model to response entity
	resp := []entities.MerchantSubmissionResponse{}
	for _, d := range data {
		resp = append(resp, entities.MerchantSubmissionResponse{
			ID:           d.ID,
			MerchantName: d.MerchantName,
			OwnerName:    d.OwnerName,
			Address:      d.Address,
			Status:       d.Status,
			CreatedAt:    d.CreatedAt,
		})
	}

	return resp, nil
}

func (s *merchantSubmissionService) CheckerApprove(id uint, checkerID uint, notes string) error {
	sub, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if sub.Status != "submitted" {
		return errors.New("hanya pengajuan submitted yang dapat dicek")
	}
	sub.Status = "checked"
	sub.CheckerID = &checkerID
	sub.Notes = notes
	sub.UpdatedAt = time.Now()
	return s.repo.Update(sub)
}

func (s *merchantSubmissionService) CheckerReject(id uint, checkerID uint, notes string) error {
	sub, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if sub.Status != "submitted" {
		return errors.New("hanya pengajuan submitted yang dapat ditolak oleh checker")
	}
	sub.Status = "rejected"
	sub.CheckerID = &checkerID
	sub.Notes = notes
	sub.UpdatedAt = time.Now()
	return s.repo.Update(sub)
}

func (s *merchantSubmissionService) SignerApprove(id uint, signerID uint) error {
	sub, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if sub.Status != "checked" {
		return errors.New("hanya pengajuan checked yang bisa disetujui signer")
	}
	// mark approved
	sub.Status = "approved"
	sub.SignerID = &signerID
	sub.UpdatedAt = time.Now()
	if err := s.repo.Update(sub); err != nil {
		return err
	}

	// create final merchant record
	var docs map[string]interface{}
	_ = json.Unmarshal(sub.Documents, &docs)

	merchant := &models.Merchant{
		Name:      sub.MerchantName,
		Owner:     sub.OwnerName,
		Phone:     sub.Phone,
		Email:     sub.Email,
		Address:   sub.Address,
		Category:  sub.Category,
		NIB:       sub.NIB,
		NPWP:      sub.NPWP,
		LegalDocs: sub.Documents,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.repo.CreateApproved(merchant); err != nil {
		return err
	}
	return nil
}

func (s *merchantSubmissionService) SignerReject(id uint, signerID uint, notes string) error {
	sub, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if sub.Status != "checked" {
		return errors.New("hanya pengajuan checked yang bisa ditolak signer")
	}
	sub.Status = "rejected"
	sub.SignerID = &signerID
	sub.Notes = notes
	sub.UpdatedAt = time.Now()
	return s.repo.Update(sub)
}

// helper converters
func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch t := v.(type) {
	case string:
		return t
	default:
		b, _ := json.Marshal(v)
		return string(b)
	}
}

func toInt(v interface{}) int {
	if v == nil {
		return 0
	}
	switch t := v.(type) {
	case int:
		return t
	case float64:
		return int(t)
	case int64:
		return int(t)
	default:
		return 0
	}
}
