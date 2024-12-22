package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"

	"github.com/gin-gonic/gin"
)

type (
	IPurchaseController interface {
		GetDropDown(ctx *gin.Context)
		Purchase(ctx *gin.Context)
	}

	PurchaseController struct {
		IPurchaseService Services.IPurchaseService
	}
)

func PurchaseControllerProvider(purchaseService Services.IPurchaseService) *PurchaseController {
	return &PurchaseController{IPurchaseService: purchaseService}
}

func (p *PurchaseController) GetDropDown(ctx *gin.Context) {
	companyID := ctx.GetString("company_id")

	res, err := p.IPurchaseService.GetDropdown(companyID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

func (p *PurchaseController) Purchase(ctx *gin.Context) {
	dto := Dto.CreatePurchasesDto{}
	if err := ctx.BindJSON(&dto); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	if validationErrors := Utils.ValidateRequest(ctx, &dto); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	dto.CompanyID = ctx.GetString("company_id")

	err := p.IPurchaseService.Create(dto)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), 500)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil membuat pembelian", nil, 200)
}
