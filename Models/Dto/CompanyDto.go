package Dto

type MakeCompanyRequest struct {
	Name         string `form:"name" binding:"required"`
	Address      string `form:"address" binding:"required"`
	CodeCompany  string `form:"code_company"`
	ImageCompany string
}

type EditCompanyRequest struct {
	Name         string `form:"name" binding:"required"`
	Address      string `form:"address" binding:"required"`
	ImageCompany string
}
