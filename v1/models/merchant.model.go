package models

import (
	"time"

	"gorm.io/datatypes"
)

type MerchantSubmission struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	MerchantName    string         `gorm:"size:255;not null" json:"merchant_name"`
	OwnerName       string         `gorm:"size:255" json:"owner_name"`
	Phone           string         `gorm:"size:50" json:"phone"`
	Email           string         `gorm:"size:100" json:"email"`
	Address         string         `gorm:"type:text" json:"address"`
	Category        string         `gorm:"size:100" json:"category"`
	NIB             string         `gorm:"size:100" json:"nib"`
	NPWP            string         `gorm:"size:100" json:"npwp"`
	YearEstablished int            `json:"year_established"`
	Employees       int            `json:"employees"`
	Documents       datatypes.JSON `gorm:"type:json" json:"documents"` // JSON metadata for uploaded files
	Status          string         `gorm:"size:50;index" json:"status"`
	MakerID         uint           `json:"maker_id"`
	CheckerID       *uint          `json:"checker_id,omitempty"`
	SignerID        *uint          `json:"signer_id,omitempty"`
	Notes           string         `gorm:"type:text" json:"notes,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

type Merchant struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Owner     string         `gorm:"size:255" json:"owner"`
	Phone     string         `gorm:"size:50" json:"phone"`
	Email     string         `gorm:"size:100" json:"email"`
	Address   string         `gorm:"type:text" json:"address"`
	Category  string         `gorm:"size:100" json:"category"`
	NIB       string         `gorm:"size:100" json:"nib"`
	NPWP      string         `gorm:"size:100" json:"npwp"`
	LegalDocs datatypes.JSON `gorm:"type:json" json:"legal_docs"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
