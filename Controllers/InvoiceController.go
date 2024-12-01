package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	IInvoiceController interface {
		CreateInvoicePurchased(ctx *gin.Context)
		GetSalesHistory(ctx *gin.Context)
		GetSpesifySalesHistory(ctx *gin.Context)
		UpdateRefund(ctx *gin.Context)
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

	search := ctx.Query("search")
	status := ctx.Query("status")

	limit, page := Utils.GetPaginationParams(ctx, Common.DEFAULTLIMIT, Common.DEFAULTPAGE)

	var query Common.Query

	if status != "" {
		status, err := strconv.ParseBool(status)

		if err != nil {
			Helper.SetErrorResponse(ctx, "Invalid status", http.StatusBadRequest)
			return
		}

		query.Status = status
	}

	query.Search = &search
	query.Limit = limit
	query.Page = page
	query.Limit = limit

	invoices, meta, statusCode, err := controller.InvoiceService.GetAllByCompany(ctx.GetString("company_id"), &query)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetPaginationResponse(
		ctx,
		"Berhasil mendapatkan data riwayat penjualan!",
		invoices,
		Common.PaginateMetadata(ctx, meta.TotalData, meta.Limit, meta.Page),
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
