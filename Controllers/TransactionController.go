package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ITransactionController interface {
		GetAll(ctx *gin.Context)
		Store(ctx *gin.Context)
	}

	TransactionController struct {
		transactionService Services.ITransactionService
	}
)

func TransactionControllerProvider(transactionService Services.ITransactionService) *TransactionController {
	return &TransactionController{transactionService: transactionService}
}

func (controller *TransactionController) GetAll(ctx *gin.Context) {
	query := Utils.InsertParams(ctx)

	sellableProducts, meta, statusCode, err := controller.transactionService.GetAll(&query)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetPaginationResponse(ctx,
		"Berhasil mendapatkan data sellable product active",
		sellableProducts,
		meta,
		statusCode)
}

func (controller *TransactionController) Store(ctx *gin.Context) {
	var dto Dto.CreateTransactionDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		Helper.SetValidationErrorResponse(ctx, err.Error())
		return
	}

	dto.CompanyID = ctx.GetString("company_id")
	dto.Name = ctx.GetString("name")
	companyCode := ctx.GetString("company_code")

	recordCode, statusCode, err := controller.transactionService.Store(&dto, companyCode)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil menambahkan data transaksi", gin.H{
		"transaction_record_code": recordCode,
	}, statusCode)
}
