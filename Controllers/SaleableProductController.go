package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Services"

	"github.com/gin-gonic/gin"
)

type (
	ISaleableProductController interface {
		FindAllSaleableProduct(ctx *gin.Context)
	}

	SaleableProductController struct {
		saleableProductService Services.ISaleableProductService
	}
)

func SaleableProductControllerProvider(saleableProductService Services.ISaleableProductService) *SaleableProductController {
	return &SaleableProductController{
		saleableProductService: saleableProductService,
	}
}

func (c *SaleableProductController) FindAllSaleableProduct(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")
	categoryQueries := ctx.QueryArray("category")

	saleableProducts, err, statusCode := c.saleableProductService.FindAllSaleableProducts(companyId, categoryQueries)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success get saleable products",
		"data":    saleableProducts,
	}, statusCode)
}
