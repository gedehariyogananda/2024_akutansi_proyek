package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	IPurchaseController interface {
		GetDropDown(ctx *gin.Context)
		Purchase(ctx *gin.Context)
		GetAllWithStatistic(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	PurchaseController struct {
		PurchaseService Services.IPurchaseService
	}
)

func PurchaseControllerProvider(purchaseService Services.IPurchaseService) *PurchaseController {
	return &PurchaseController{PurchaseService: purchaseService}
}

func (p *PurchaseController) GetDropDown(ctx *gin.Context) {
	companyID := ctx.GetString("company_id")

	res, err := p.PurchaseService.GetDropdown(companyID)
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

	fmt.Println(dto.DueDate)

	if validationErrors := Utils.ValidateRequest(ctx, &dto); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	dto.CompanyID = ctx.GetString("company_id")

	err := p.PurchaseService.Create(dto)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), 500)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil membuat pembelian", nil, 200)
}

func (p *PurchaseController) GetAllWithStatistic(ctx *gin.Context) {
	companyID := ctx.GetString("company_id")

	var dto Common.Query

	perage, page := Utils.GetPaginationParams(ctx, Common.DEFAULTLIMIT, Common.DEFAULTPAGE)

	dto.Page = page
	dto.Limit = perage

	res, meta, err := p.PurchaseService.GetAllPurchaseWithStatistic(companyID, &dto)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), 500)
		return
	}

	Helper.SetPaginationResponse(ctx, "Get all purchases success", res, meta, 200)
}

func (p *PurchaseController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := p.PurchaseService.Delete(id)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), 500)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil menghapus pembelian", nil, 200)
}
