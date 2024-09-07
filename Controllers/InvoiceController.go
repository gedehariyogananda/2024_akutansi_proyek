package Controllers

import (
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	IInvoiceController interface {
		CreateInvoicePurchased(ctx *gin.Context)
	}

	InvoiceController struct {
		InvoiceService Services.IInvoiceService
	}
)

func InvoiceControllerProvider(invoiceService Services.IInvoiceService) *InvoiceController {
	return &InvoiceController{InvoiceService: invoiceService}
}

func (c *InvoiceController) CreateInvoicePurchased(ctx *gin.Context) {

	companyId := ctx.GetInt("company_id")

	var request Dto.InvoiceRequestClient
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if err := c.InvoiceService.CreateInvoicePurchased(&request, companyId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success create invoice purchased"})
}
