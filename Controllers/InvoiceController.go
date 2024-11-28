package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"

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

func (controller *InvoiceController) CreateInvoicePurchased(ctx *gin.Context) {
	var requestInvoiceDTO Dto.InvoiceRequestDTO

	if err := ctx.ShouldBindJSON(&requestInvoiceDTO); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	if validationErrors := Utils.ValidateRequest(ctx, &requestInvoiceDTO); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	invoice, statusCode, err := controller.InvoiceService.CreateInvoicePurchased(&requestInvoiceDTO, ctx.GetString("company_id"))
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Create Transaction Purchased Success!", gin.H{
		"invoice_number": invoice.InvoiceNumber,
		"customer_name":  invoice.CustomerName,
		"total_price":    invoice.SubTotal,
	}, statusCode)
}
