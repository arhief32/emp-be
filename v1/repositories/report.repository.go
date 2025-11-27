package repositories

import (
	"time"

	"gorm.io/gorm"

	"github.com/yourusername/pegawai-api/v1/models"
)

type ReportRepository interface {
	SaveDailyReport(models.DailyReport) error
	GetDailyReportByDate(date time.Time) (models.DailyReport, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) SaveDailyReport(dr models.DailyReport) error {
	dr.CreatedAt = time.Now()
	return r.db.Create(&dr).Error
}

func (r *reportRepository) GetDailyReportByDate(date time.Time) (models.DailyReport, error) {
	var dr models.DailyReport
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	if err := r.db.Where("date = ?", start).First(&dr).Error; err != nil {
		return models.DailyReport{}, err
	}
	return dr, nil
}
