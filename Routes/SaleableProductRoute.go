package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaleableProductRoute(c *gin.RouterGroup, db *gorm.DB) {
	route := c.Group("/saleable-product")

	m := Di.DICommonMiddleware(db)

	// open use authenticate
	route.Use(m.IsAuthenticate)

	SaleableProductController := Di.DISaleableProduct(db)

	route.GET("/", SaleableProductController.FindAllSaleableProduct)

}
