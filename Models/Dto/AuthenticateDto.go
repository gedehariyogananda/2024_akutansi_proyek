package Dto

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Token    string `json:"token"`
	Me       bool   `json:"me" binding:"required"`
}

type TokenCompanyRequest struct {
	CompanyID string `json:"company_id"`
}
