package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func LogActivityRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/log-activity")

	m := Di.DICommonMiddleware(db, redis)

	route.Use(m.IsAuthenticate)

	LogActivityController := Di.DiLogActivity(db)

	route.GET("/", LogActivityController.FindAll)
}
