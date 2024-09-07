package Controllers

import (
	"2024_akutansi_project/Services"
	"strconv"

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
	idParam := ctx.Param("id")
	companyID, _ := strconv.Atoi(idParam)

	saleableProducts, materialProduct, err := c.saleableProductService.FindAllSaleableProducts(companyID)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return

	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    gin.H{"saleable_products": saleableProducts, "material_products": materialProduct, "company_id": companyId},
	})
}
