package entities

import "time"

// Request/response shapes for submission

type MerchantSubmissionCreateRequest struct {
	MerchantName    string                 `json:"merchant_name" binding:"required"`
	OwnerName       string                 `json:"owner_name" binding:"required"`
	Phone           string                 `json:"phone"`
	Email           string                 `json:"email"`
	Address         string                 `json:"address"`
	Category        string                 `json:"category"`
	NIB             string                 `json:"nib"`
	NPWP            string                 `json:"npwp"`
	YearEstablished int                    `json:"year_established"`
	Employees       int                    `json:"employees"`
	Documents       map[string]interface{} `json:"documents"` // flexible JSON map
}

type MerchantSubmissionUpdateRequest struct {
	MerchantName    *string                `json:"merchant_name"`
	OwnerName       *string                `json:"owner_name"`
	Phone           *string                `json:"phone"`
	Email           *string                `json:"email"`
	Address         *string                `json:"address"`
	Category        *string                `json:"category"`
	NIB             *string                `json:"nib"`
	NPWP            *string                `json:"npwp"`
	YearEstablished *int                   `json:"year_established"`
	Employees       *int                   `json:"employees"`
	Documents       map[string]interface{} `json:"documents"`
}

type MerchantSubmissionResponse struct {
	ID           uint                   `json:"id"`
	MerchantName string                 `json:"merchant_name"`
	OwnerName    string                 `json:"owner_name"`
	Phone        string                 `json:"phone"`
	Email        string                 `json:"email"`
	Address      string                 `json:"address"`
	Category     string                 `json:"category"`
	NIB          string                 `json:"nib"`
	NPWP         string                 `json:"npwp"`
	Documents    map[string]interface{} `json:"documents"`
	Status       string                 `json:"status"`
	MakerID      uint                   `json:"maker_id"`
	CheckerID    *uint                  `json:"checker_id,omitempty"`
	SignerID     *uint                  `json:"signer_id,omitempty"`
	Notes        string                 `json:"notes,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
}
