package Controllers

import (
	"2024_akutansi_project/Helper"
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
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": "Invalid request body",
		}, http.StatusBadRequest)
		return
	}

	invoice, err, statusCode := c.InvoiceService.CreateInvoicePurchased(&request, companyId)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	if invoice.PaymentMethod.MethodName == "Cash" {

		Helper.SetResponse(ctx, gin.H{
			"success": true,
			"message": "Success create invoice purchased",
			"data": gin.H{
				"invoice":     invoice,
				"is_cashless": true,
			},
		}, statusCode)
	} else {
		Helper.SetResponse(ctx, gin.H{
			"success": true,
			"message": "Success create invoice purchased",
			"data": gin.H{
				"invoice":     invoice,
				"is_cashless": false,
			},
		}, statusCode)
	}
}
