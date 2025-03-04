package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Services"

	"github.com/gin-gonic/gin"
)

type (
	ICompanyController interface {
		GetDetailCompany(ctx *gin.Context)
	}

	CompanyController struct {
		companyService Services.ICompanyService
	}
)

func CompanyControllerProvider(companyService Services.ICompanyService) *CompanyController {
	return &CompanyController{companyService: companyService}
}

func (controller *CompanyController) GetDetailCompany(ctx *gin.Context) {
	companyID := ctx.GetString("company_id")

	company, err := controller.companyService.GetDetailCompany(companyID)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), 500)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil mendapatkan data perusahaan", company, 200)
}
