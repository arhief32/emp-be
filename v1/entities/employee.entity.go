package entities

type EmployeeCreateRequest struct {
	Nama    string `json:"nama" binding:"required"`
	Nip     string `json:"nip" binding:"required"`
	Jabatan string `json:"jabatan"`
}

type EmployeeUpdateRequest struct {
	Nama    string `json:"nama"`
	Nip     string `json:"nip"`
	Jabatan string `json:"jabatan"`
}
