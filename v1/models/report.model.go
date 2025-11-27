package models

import "time"

// DailyReport is a lightweight model returned by service (also used for migration if needed)
type DailyReport struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Date      time.Time  `gorm:"index:idx_date" json:"date"`
	Count     int        `json:"count"`
	Data      []Employee `gorm:"-" json:"data"` // not persisted, just for service response
	CreatedAt time.Time  `json:"created_at"`
}
