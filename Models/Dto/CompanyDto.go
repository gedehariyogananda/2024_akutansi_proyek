package Dto

type MakeCompanyRequest struct {
	Name         string `form:"name" binding:"required"`
	Address      string `form:"address" binding:"required"`
	ImageCompany string
}

type EditCompanyRequest struct {
	Name         string `form:"name" binding:"required"`
	Address      string `form:"address" binding:"required"`
	ImageCompany string
}
