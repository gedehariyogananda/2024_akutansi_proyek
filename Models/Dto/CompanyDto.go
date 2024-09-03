package Dto

type MakeCompanyRequest struct {
	Name         string `json:"name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	ImageCompany string `json:"image_company" binding:"required"`
}
