package Dto

type RegisterRequest struct {
	Name        string `json:"name" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
}

type LoginOwnerRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Me       bool   `json:"me" binding:"required"`
}

type LoginEmployeeRequest struct {
	EmployeeKey string `json:"employee_key" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Me          bool   `json:"me" binding:"required"`
}
