package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func MaterialProduct(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/material-products")

	m := Di.DICommonMiddleware(db, redis)

	// open use authenticate
	route.Use(m.IsAuthenticate)

	MaterialProductController := Di.DIMaterialProduct(db)

	route.POST("/", MaterialProductController.Create)
}
