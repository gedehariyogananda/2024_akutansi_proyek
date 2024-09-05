package Dto

type MakeCompanyRequest struct {
	Name         string `json:"name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	ImageCompany string `json:"image_company" binding:"required"`
}

type EditCompanyRequest struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	ImageCompany string `json:"image_company"`
}
