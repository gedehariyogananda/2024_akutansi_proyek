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
		UpdateInvoiceStatus(ctx *gin.Context)
		UpdateMoneyReceived(ctx *gin.Context)
		GetAllInvoices(ctx *gin.Context)
		UpdateInvoiceCustomer(ctx *gin.Context)
	}

	InvoiceController struct {
		InvoiceService Services.IInvoiceService
	}
)

func InvoiceControllerProvider(invoiceService Services.IInvoiceService) *InvoiceController {
	return &InvoiceController{InvoiceService: invoiceService}
}

func (c *InvoiceController) CreateInvoicePurchased(ctx *gin.Context) {

	companyId := ctx.GetString("company_id")

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

func (c *InvoiceController) UpdateInvoiceStatus(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")

	invoiceParams := ctx.Param("invoice_id") // status "PROCESS", "CANCLE", "DONE" in body Request
	var request Dto.InvoiceUpdateRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": "Invalid request body",
		}, http.StatusBadRequest)
		return
	}

	invoice, err, statusCode := c.InvoiceService.UpdateStatusInvoice(&request, invoiceParams, companyId)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success update invoice status to PROCESS",
		"data":    invoice,
	}, statusCode)
	return

}

func (c *InvoiceController) UpdateMoneyReceived(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")

	invoiceParams := ctx.Param("invoice_id")
	var request Dto.InvoiceMoneyReceivedRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	invoice, moneyBack, err, statusCode := c.InvoiceService.UpdateMoneyReveived(&request, invoiceParams, companyId)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success update invoice money received",
		"data": gin.H{
			"invoice": gin.H{
				"id":             invoice.ID,
				"invoice_number": invoice.InvoiceNumber,
				"total_amount":   invoice.TotalAmount,
				"money_received": invoice.MoneyReceived,
				"money_back":     moneyBack,
				"status_invoice": invoice.StatusInvoice,
			},
		},
	}, statusCode)
}

func (c *InvoiceController) GetAllInvoices(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")

	invoices, err, statusCode := c.InvoiceService.GetAllInvoices(companyId)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success get all invoices",
		"data":    invoices,
	}, statusCode)
}

func (c *InvoiceController) UpdateInvoiceCustomer(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")

	invoiceParams := ctx.Param("invoice_id")
	var request Dto.InvoiceUpdateRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	invoice, err, statusCode := c.InvoiceService.UpdateInvoiceCustomer(companyId, invoiceParams, &request)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success update invoice customer",
		"data":    invoice,
	}, statusCode)
}
