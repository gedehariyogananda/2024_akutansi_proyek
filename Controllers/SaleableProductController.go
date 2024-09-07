package Controllers

import (
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
	companyId := ctx.GetInt("company_id")

	saleableProducts, materialProducts, err, statusCode := c.saleableProductService.FindAllSaleableProducts(companyId)

	if err != nil {
		ctx.JSON(statusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return

	}

	ctx.JSON(statusCode, gin.H{
		"success": true,
		"message": "success get saleable products",
		"data": gin.H{
			"saleable_products": saleableProducts,
			"material_products": materialProducts,
		},
	})
}
