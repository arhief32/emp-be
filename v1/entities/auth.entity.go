package entities

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	OrgDesc  string `json:"org_desc" binding:"required"`
	BranchID int    `json:"branch_id" binding:"required"`
	JobTitle string `json:"job_title" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
