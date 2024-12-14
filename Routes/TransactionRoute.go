package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func TransactionRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/transactions")

	m := Di.DICommonMiddleware(db, redis)

	// open use authenticate
	route.Use(m.IsAuthenticate)

	TransactionController := Di.DITransaction(db)

	route.GET("/", TransactionController.GetAll)
	route.POST("/", TransactionController.Store)
}
