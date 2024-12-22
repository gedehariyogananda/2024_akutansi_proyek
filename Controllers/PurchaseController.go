package Controllers

import (
	"2024_akutansi_project/Services"

	"github.com/gin-gonic/gin"
)

type (
	IPurchaseController interface {
		GetDropDown(ctx *gin.Context)
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
