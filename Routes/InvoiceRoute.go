package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InvoiceRoute(c *gin.RouterGroup, db *gorm.DB) {
	route := c.Group("/invoice")
	m := Di.DICommonMiddleware(db)

	// open use authenticate
	route.Use(m.IsAuthenticate)
	InvoiceController := Di.DIInvoice(db)

	route.GET("/checked", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "checked healt",
		})
	})

	route.POST("/create", InvoiceController.CreateInvoicePurchased)
}
