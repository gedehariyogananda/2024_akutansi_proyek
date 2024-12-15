package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	IInvoiceController interface {
		CreateInvoicePurchased(ctx *gin.Context)
		GetSalesHistory(ctx *gin.Context)
		GetSpesifySalesHistory(ctx *gin.Context)
		UpdateRefund(ctx *gin.Context)
		StatisticSales(ctx *gin.Context)
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

func (controller *InvoiceController) GetSalesHistory(ctx *gin.Context) {
	query := Utils.InsertParams(ctx)

	invoices, meta, statusCode, err := controller.InvoiceService.GetAllByCompany(ctx.GetString("company_id"), &query)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetPaginationResponse(
		ctx,
		"Berhasil mendapatkan data riwayat penjualan!",
		invoices,
		meta,
		statusCode,
	)

}

func (controller *InvoiceController) GetSpesifySalesHistory(ctx *gin.Context) {
	invoice, statusCode, err := controller.InvoiceService.GetSpesifySalesHistory(ctx.GetString("company_id"), ctx.Param("invoiceID"))
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil mendapatkan data penjualan!", invoice, statusCode)
}

func (controller *InvoiceController) UpdateRefund(ctx *gin.Context) {
	statusCode, err := controller.InvoiceService.UpdateRefund(ctx.GetString("company_id"), ctx.Param("id"))
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil melakukan pengembalian dana!", nil, statusCode)
}

func (controller *InvoiceController) StatisticSales(ctx *gin.Context) {
	dateNow := time.Now().Format("2006-01-02")

	statistic, statusCode, err := controller.InvoiceService.StatisticSales(ctx.GetString("company_id"), ctx.DefaultQuery("date", dateNow))
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil mendapatkan data statistik penjualan!", statistic, statusCode)

}
