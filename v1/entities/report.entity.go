package entities

import "time"

type DailyReportResponse struct {
	Date  time.Time     `json:"date"`
	Count int           `json:"count"`
	Data  []interface{} `json:"data"`
}
