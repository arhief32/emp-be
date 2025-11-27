package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:50" json:"username"`
	Password  string    `json:"-"` // jangan expose password
	Name      string    `json:"name"`
	OrgDesc   string    `json:"org_desc"`
	BranchID  int       `json:"branch_id"`
	JobTitle  string    `json:"job_title"`
	CreatedAt time.Time `json:"created_at"`
}
