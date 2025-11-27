package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/arhief32/emp-be/v1/services"
)

type ReportController struct {
	svc services.ReportService
}

func NewReportController(s services.ReportService) ReportController {
	return ReportController{svc: s}
}

// GET /v1/reports/daily?date=YYYY-MM-DD
func (rc *ReportController) Daily(c *gin.Context) {
	dateStr := c.Query("date")
	var t time.Time
	var err error
	if dateStr == "" {
		t = time.Now()
	} else {
		t, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format (YYYY-MM-DD)"})
			return
		}
	}
	resp, err := rc.svc.DailyReport(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
}
