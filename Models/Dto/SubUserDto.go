package Dto

type CreateSubUserDto struct {
	Name        string `json:"name" binding:"required"`
	EmployeeKey string `json:"employee_key" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Status      bool   `json:"status"`
	CompanyID   string `json:"_"`
}

type UpdateSubUserDto struct {
	Name        string `json:"name"`
	EmployeeKey string `json:"employee_key"`
	Password    string `json:"password"`
	Status      bool   `json:"status"`
	CompanyID   string `json:"_"`
}
