package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func InvoiceRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/invoice")
	m := Di.DICommonMiddleware(db, redis)

	route.Use(m.RolesAll)
	InvoiceController := Di.DIInvoice(db)

	route.GET("/checked", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "checked healt",
		})
	})

	route.POST("/create", InvoiceController.CreateInvoicePurchased)
	route.GET("/sales/history", InvoiceController.GetSalesHistory)
	route.GET("/sales/history/:invoiceID", InvoiceController.GetSpesifySalesHistory)
	route.PUT("/refund/:id", InvoiceController.UpdateRefund)
	route.PUT("/cashier/update/:invoiceID", InvoiceController.UpdateCashier)
	route.GET("/statistic/sales", InvoiceController.StatisticSales)
}
