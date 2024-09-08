package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PaymentMethodRoute(c *gin.RouterGroup, db *gorm.DB) {
	route := c.Group("/payment-method")

	m := Di.DICommonMiddleware(db)

	// open use authenticate
	route.Use(m.IsAuthenticate)

	PaymentMethodController := Di.DIPaymentMethod(db)

	route.GET("/", PaymentMethodController.FindAllPaymentMethod)
}
