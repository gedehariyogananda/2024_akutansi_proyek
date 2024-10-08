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
		GetInvoiceDetail(ctx *gin.Context)
		DeleteInvoice(ctx *gin.Context)
		UpdateInvoiceDetail(ctx *gin.Context)
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
	filterDate := ctx.DefaultQuery("date", "")
	formattedDateClient := Helper.FormatDateClient(filterDate)

	companyId := ctx.GetString("company_id")

	invoices, err, statusCode := c.InvoiceService.GetAllInvoices(companyId, formattedDateClient)
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

func (c *InvoiceController) GetInvoiceDetail(ctx *gin.Context) {
	invoiceParams := ctx.Param("invoice_id")

	invoice, purchaseDetail, err, statusCode := c.InvoiceService.GetInvoice(invoiceParams)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success get invoice detail",
		"data": gin.H{
			"purchase_detail": purchaseDetail,
			"invoice":         invoice,
		},
	}, statusCode)
}

func (c *InvoiceController) DeleteInvoice(ctx *gin.Context) {
	invoiceParams := ctx.Param("invoice_id")
	companyId := ctx.GetString("company_id")

	statusCode, err := c.InvoiceService.DeleteInvoice(invoiceParams, companyId)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success delete invoice",
	}, statusCode)
}

func (c *InvoiceController) UpdateInvoiceDetail(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")

	invoiceParams := ctx.Param("invoice_id")
	var request Dto.InvoiceRequestClient
	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	invoice, err, statusCode := c.InvoiceService.UpdateInvoiceDetail(companyId, invoiceParams, &request)
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
			"message": "Success update invoice detail",
			"data": gin.H{
				"invoice":     invoice,
				"is_cashless": true,
			},
		}, statusCode)
	} else {
		Helper.SetResponse(ctx, gin.H{
			"success": true,
			"message": "Success update invoice detail",
			"data": gin.H{
				"invoice":     invoice,
				"is_cashless": false,
			},
		}, statusCode)
	}
}
