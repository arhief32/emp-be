package repositories

import (
	"gorm.io/gorm"

	"github.com/arhief32/emp-be/v1/models"
)

type MerchantSubmissionRepository interface {
	Create(sub *models.MerchantSubmission) error
	Update(sub *models.MerchantSubmission) error
	FindByID(id uint) (*models.MerchantSubmission, error)
	FindByMaker(makerID uint) ([]models.MerchantSubmission, error)
	FindPendingForChecker() ([]models.MerchantSubmission, error)
	FindPendingForSigner() ([]models.MerchantSubmission, error)

	CreateApproved(m *models.Merchant) error
	FindByIDApproved(id uint) (*models.Merchant, error)
}

type merchantSubmissionRepo struct {
	db *gorm.DB
}

func NewMerchantSubmissionRepository(db *gorm.DB) MerchantSubmissionRepository {
	return &merchantSubmissionRepo{db: db}
}

func (r *merchantSubmissionRepo) Create(sub *models.MerchantSubmission) error {
	return r.db.Create(sub).Error
}

func (r *merchantSubmissionRepo) Update(sub *models.MerchantSubmission) error {
	return r.db.Save(sub).Error
}

func (r *merchantSubmissionRepo) FindByID(id uint) (*models.MerchantSubmission, error) {
	var sub models.MerchantSubmission
	if err := r.db.First(&sub, id).Error; err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *merchantSubmissionRepo) FindByMaker(makerID uint) ([]models.MerchantSubmission, error) {
	var list []models.MerchantSubmission
	if err := r.db.Where("maker_id = ?", makerID).Order("created_at desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *merchantSubmissionRepo) FindPendingForChecker() ([]models.MerchantSubmission, error) {
	var list []models.MerchantSubmission
	if err := r.db.Where("status = ?", "submitted").Order("created_at asc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *merchantSubmissionRepo) FindPendingForSigner() ([]models.MerchantSubmission, error) {
	var list []models.MerchantSubmission
	if err := r.db.Where("status = ?", "checked").Order("created_at asc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *merchantSubmissionRepo) CreateApproved(m *models.Merchant) error {
	return r.db.Create(m).Error
}

func (r *merchantSubmissionRepo) FindByIDApproved(id uint) (*models.Merchant, error) {
	var m models.Merchant
	if err := r.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
