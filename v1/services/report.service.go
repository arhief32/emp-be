package services

import (
	"time"

	"github.com/arhief32/emp-be/v1/models"
	"github.com/arhief32/emp-be/v1/repositories"
)

type ReportService interface {
	DailyReport(date time.Time) (models.DailyReport, error)
}

type reportService struct {
	reportRepo repositories.ReportRepository
	empRepo    repositories.EmployeeRepository
}

func NewReportService(rr repositories.ReportRepository, er repositories.EmployeeRepository) ReportService {
	return &reportService{reportRepo: rr, empRepo: er}
}

// DailyReport returns count and list of employees created on that date
func (s *reportService) DailyReport(date time.Time) (models.DailyReport, error) {
	// normalize date start/end
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	end := start.Add(24 * time.Hour).Add(-time.Nanosecond)

	count, err := s.empRepo.CountByCreatedAtRange(start, end)
	if err != nil {
		return models.DailyReport{}, err
	}
	records, err := s.empRepo.FindByCreatedAtRange(start, end)
	if err != nil {
		return models.DailyReport{}, err
	}
	dr := models.DailyReport{
		Date:  start,
		Count: count,
		Data:  records,
	}
	// optional: persist report record if desired via reportRepo (not required)
	_ = s.reportRepo // kept in service for possible future use
	return dr, nil
}
