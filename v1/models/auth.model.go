package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:50" json:"username"`
	Password  string    `json:"-"` // jangan expose password
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"created_at"`
}
