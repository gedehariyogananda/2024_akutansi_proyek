package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SubUserRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/sub-users")

	m := Di.DICommonMiddleware(db, redis)

	route.Use(m.RolesAll)

	SubUserController := Di.DISubUser(db)

	route.POST("/", SubUserController.Create)
	route.GET("/:id", SubUserController.FindByID)
	route.GET("/", SubUserController.FindAll)
	route.PUT("/:id", SubUserController.Update)
	route.DELETE("/:id", SubUserController.Delete)

}
