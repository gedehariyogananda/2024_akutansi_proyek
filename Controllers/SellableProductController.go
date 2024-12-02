package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"

	"github.com/gin-gonic/gin"
)

type (
	ISellableProductController interface {
		GetAllSellableProduct(ctx *gin.Context)
	}

	SellableProductController struct {
		SellableProductService Services.ISellableProductService
	}
)

func SellableProductControllerProvider(SellableProductService Services.ISellableProductService) *SellableProductController {
	return &SellableProductController{
		SellableProductService: SellableProductService,
	}
}

func (controller *SellableProductController) GetAllSellableProduct(ctx *gin.Context) {
	search := ctx.Query("search")

	limit, page := Utils.GetPaginationParams(ctx, Common.DEFAULTLIMIT, Common.DEFAULTPAGE)

	var query Common.Query

	query.Search = &search
	query.Limit = limit
	query.Page = page
	query.Limit = limit

	sellableProducts, meta, statusCode, err := controller.SellableProductService.GetAll(ctx.GetString("company_id"), &query)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetPaginationResponse(ctx,
		"Berhasil mendapatkan data sellable product",
		sellableProducts,
		Common.PaginateMetadata(ctx, meta.TotalData, meta.Limit, meta.Page),
		statusCode)
}
