package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Promo(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/promos")

	m := Di.DICommonMiddleware(db, redis)

	// open use authenticate
	route.Use(m.IsAuthenticate)

	PromoController := Di.DIPromo(db)

	route.POST("/", PromoController.Create)
	route.GET("/:id", PromoController.FindByID)
	route.DELETE("/:id", PromoController.Delete)
	route.PUT("/:id", PromoController.Update)
	route.GET("/", PromoController.FindAll)
}
