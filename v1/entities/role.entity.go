package entities

type CreateRoleRequest struct {
	Role     string `json:"role" binding:"required"`
	OrgDesc  string `json:"org_desc"`
	JobTitle string `json:"job_title"`
}

type UpdateRoleRequest struct {
	Role     string `json:"role" binding:"required"`
	OrgDesc  string `json:"org_desc"`
	JobTitle string `json:"job_title"`
}
