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

	route.POST("/", SellableController.Create)
	route.DELETE("/", SellableController.UnAssignMAterial)
	route.POST("/assign-material", SellableController.AssignMaterial)
	route.GET("/", SellableController.GetAllSellableProduct)
	route.PUT("/:id/stock", SellableController.UpdateSellableProduct)
	route.GET("/:id", SellableController.FindById)
	route.DELETE("/:id", SellableController.Delete)
	route.PUT("/:id", SellableController.Update)
}
