package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SellableProductRoutes(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/sellable")

	m := Di.DICommonMiddleware(db, redis)

	route.Use(m.IsAuthenticate)

	SellableController := Di.DISellableProduct(db)

	route.GET("/", SellableController.GetAllSellableProduct)
	route.PATCH("/:id", SellableController.UpdateSellableProduct)

}
