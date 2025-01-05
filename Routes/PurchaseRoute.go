package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func PurchaseRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/purchase")

	m := Di.DICommonMiddleware(db, redis)

	route.Use(m.IsAuthenticate)

	PurchaseController := Di.DiPurchase(db)

	route.GET("/dropdown", PurchaseController.GetDropDown)
	route.POST("/", PurchaseController.Purchase)
	route.GET("/", PurchaseController.GetAllWithStatistic)
	route.DELETE("/:id", PurchaseController.Delete)
}
