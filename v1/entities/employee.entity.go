package entities

type EmployeeCreateRequest struct {
	Nama    string `json:"nama" binding:"required"`
	NIP     string `json:"nip" binding:"required"`
	Jabatan string `json:"jabatan"`
}

type EmployeeUpdateRequest struct {
	Nama    string `json:"nama"`
	NIP     string `json:"nip"`
	Jabatan string `json:"jabatan"`
}
