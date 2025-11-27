package models

import "time"

type Employee struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"size:255;not null" json:"nama"`
	NIP       string    `gorm:"size:100;uniqueIndex;not null" json:"nip"`
	Jabatan   string    `gorm:"size:255" json:"jabatan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
