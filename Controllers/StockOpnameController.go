package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IStockOpnameController interface {
		GetAll(ctx *gin.Context)
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
