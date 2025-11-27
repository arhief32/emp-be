package models

import "time"

type Role struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Role      string    `json:"role"`
	OrgDesc   string    `json:"org_desc"`
	JobTitle  string    `json:"job_title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
