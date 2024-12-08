package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CategoryRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/category")

	m := Di.DICommonMiddleware(db, redis)

	// open use authenticate
	route.Use(m.IsAuthenticate)

	CategoryController := Di.DICategory(db)

	route.GET("/", CategoryController.FindAll)
	route.GET("/:id", CategoryController.FindByID)

	route.POST("/", CategoryController.Create)
	route.PATCH("/:id", CategoryController.Update)
	route.DELETE("/:id", CategoryController.Delete)
}
