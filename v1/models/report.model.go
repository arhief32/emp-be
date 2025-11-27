package models

import "time"

// DailyReport is a lightweight model returned by service (also used for migration if needed)
type DailyReport struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Date      time.Time `gorm:"index:idx_date" json:"date"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"created_at"`
}
