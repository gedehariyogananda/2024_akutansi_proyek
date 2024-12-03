package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Services"

	"github.com/gin-gonic/gin"
)

type (
	ITaxController interface {
		GetAllTax(ctx *gin.Context)
	}

	TaxController struct {
		TaxService Services.ITaxService
	}
)

func TaxControllerProvider(TaxService Services.ITaxService) *TaxController {
	return &TaxController{
		TaxService: TaxService,
	}
}

func (controller *TaxController) GetAllTax(ctx *gin.Context) {
	taxs, statusCode, err := controller.TaxService.GetAll()
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Success Menampilkan Semua Data Pajak", taxs, statusCode)
}
