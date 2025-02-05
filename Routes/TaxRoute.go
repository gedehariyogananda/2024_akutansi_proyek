package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func TaxRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/tax")
	m := Di.DICommonMiddleware(db, redis)

	route.Use(m.RolesAll)

	TaxController := Di.DITax(db)

	route.GET("/", TaxController.GetAllTax)
}
