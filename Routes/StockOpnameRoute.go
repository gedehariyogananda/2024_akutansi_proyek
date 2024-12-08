package Routes

import (
	"2024_akutansi_project/Routes/Di"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func StockOpnameRoutes(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/stock-opname")

	m := Di.DICommonMiddleware(db, redis)

	route.Use(m.IsAuthenticate)

	controller := Di.DIStockOpname(db)

	route.GET("/", controller.GetAll)
	route.POST("/", controller.Create)
	route.GET("/available-stock", controller.GetAvailableStock)
}
