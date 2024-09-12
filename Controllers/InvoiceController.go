package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	IInvoiceController interface {
		CreateInvoicePurchased(ctx *gin.Context)
		UpdateInvoiceStatus(ctx *gin.Context)
		UpdateMoneyReceived(ctx *gin.Context)
		GetAllInvoices(ctx *gin.Context)
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

func (c *InvoiceController) UpdateInvoiceStatus(ctx *gin.Context) {
	companyId := ctx.GetInt("company_id")

	invoiceParams := ctx.Param("invoice_id") // status "PROCESS", "CANCLE" in body Request
	invoiceId, _ := strconv.Atoi(invoiceParams)
	var request Dto.InvoiceStatusRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": "Invalid request body",
		}, http.StatusBadRequest)
		return
	}

	invoice, err, statusCode := c.InvoiceService.UpdateStatusInvoice(&request, invoiceId, companyId)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	if request.StatusInvoice == "PROCESS" {
		Helper.SetResponse(ctx, gin.H{
			"success": true,
			"message": "Success update invoice status to PROCESS",
			"data": gin.H{
				"status_invoice": invoice.StatusInvoice,
			},
		}, statusCode)
		return
	}

	if request.StatusInvoice == "CANCEL" {
		Helper.SetResponse(ctx, gin.H{
			"success": true,
			"message": "Success update invoice status to CANCEL",
			"data": gin.H{
				"status_invoice": invoice.StatusInvoice,
			},
		}, statusCode)
		return
	}
}

func (c *InvoiceController) UpdateMoneyReceived(ctx *gin.Context) {
	companyId := ctx.GetInt("company_id")

	invoiceParams := ctx.Param("invoice_id")
	invoiceId, _ := strconv.Atoi(invoiceParams)
	var request Dto.InvoiceMoneyReceivedRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	invoice, err, statusCode := c.InvoiceService.UpdateMoneyReveived(&request, invoiceId, companyId)
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
		"data":    invoice,
	}, statusCode)
}

func (c *InvoiceController) GetAllInvoices(ctx *gin.Context) {
	companyId := ctx.GetInt("company_id")

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
