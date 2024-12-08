package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IStockOpnameController interface {
		GetAll(ctx *gin.Context)
		Create(ctx *gin.Context)
	}

	StockOpnameController struct {
		StockOpnameService Services.IStockOpnameService
	}
)

func StockOpnameControllerProvider(stockOpnameService Services.IStockOpnameService) *StockOpnameController {
	return &StockOpnameController{StockOpnameService: stockOpnameService}
}

func (controller *StockOpnameController) GetAll(ctx *gin.Context) {
	query := Utils.InsertParams(ctx)

	data, meta, err := controller.StockOpnameService.GetAll(&query)

	if err != nil {
		statusCode := Utils.HandleStatusCode(err)
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetPaginationResponse(ctx,
		"Berhasil mendapatkan data stock opname",
		data,
		meta,
		http.StatusOK,
	)
}

func (controller *StockOpnameController) Create(ctx *gin.Context) {
	var dto Dto.CreateStockOpnameDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	dto.CompanyID = ctx.GetString("company_id")
	dto.ChangerName = ctx.GetString("name")

	err := controller.StockOpnameService.Create(&dto)
	if err != nil {
		statusCode := Utils.HandleStatusCode(err)
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(
		ctx,
		"Berhasil membuat stock opname",
		nil,
		http.StatusCreated)
}
