package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Unit(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/units")

	m := Di.DICommonMiddleware(db, redis)

	// open use authenticate
	route.Use(m.IsAuthenticate)

	UnitController := Di.DIUnit(db)

	route.POST("/", UnitController.Create)
	route.GET("/", UnitController.FindAll)
	route.PATCH("/:id", UnitController.Update)
	route.DELETE("/:id", UnitController.Delete)
}
