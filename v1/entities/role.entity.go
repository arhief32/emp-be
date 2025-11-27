package entities

type CreateRoleRequest struct {
	Role     string `json:"role" binding:"required"`
	OrgDesc  string `json:"org_desc" binding:"required"`
	JobTitle string `json:"job_title" binding:"required"`
}

type UpdateRoleRequest struct {
	Role     string `json:"role" binding:"required"`
	OrgDesc  string `json:"org_desc" binding:"required"`
	JobTitle string `json:"job_title" binding:"required"`
}
