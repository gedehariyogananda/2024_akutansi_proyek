package Dto

type MakeUserCompanyRequest struct {
	UserId    string `json:"user_id"`
	CompanyId string `json:"company_id"`
}