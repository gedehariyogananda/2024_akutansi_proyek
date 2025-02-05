package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Account(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/accounts")

	m := Di.DICommonMiddleware(db, redis)

	// open use authenticate
	route.Use(m.RolesAll)

	AccountController := Di.DIAccount(db)

	route.POST("/", AccountController.Create)
	route.GET("/:id", AccountController.FindByID)
	route.PUT("/:id", AccountController.Update)
	route.DELETE("/:id", AccountController.Delete)
	route.GET("/", AccountController.FindAll)
}
