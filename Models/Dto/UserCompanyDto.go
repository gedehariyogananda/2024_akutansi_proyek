package Dto 

type MakeUserCompanyRequest struct {
	UserId    int `json:"user_id"`
	CompanyId int `json:"company_id"`
}