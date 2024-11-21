package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func PaymentMethodRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/payment-method")

	m := Di.DICommonMiddleware(db, redis)

	// open use authenticate
	route.Use(m.IsAuthenticate)

	PaymentMethodController := Di.DIPaymentMethod(db)

	route.GET("/", PaymentMethodController.FindAllPaymentMethod)
	route.POST("/", PaymentMethodController.CreatePaymentMethod)
	route.PUT("/:id", PaymentMethodController.UpdatePaymentMethod)
	route.DELETE("/:id", PaymentMethodController.DeletePaymentMethod)

}
