package Dto

type RegisterRequest struct {
	Name        string `json:"name" validate:"required"`
	Phone       string `json:"phone" validate:"required,min=10,max=15"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	CompanyName string `json:"company_name" validate:"required"`
}

type LoginOwnerRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	//Device   string `json:"device" validate:"required"`
}

type LoginEmployeeRequest struct {
	EmployeeKey string `json:"employee_key" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type LoginMobileRequest struct {
	Key      string `json:"key" validate:"required"`
	Password string `json:"password" validate:"required"`
}
