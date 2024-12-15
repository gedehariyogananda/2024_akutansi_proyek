package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func AuthRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/auth")

	authController := Di.DIAuth(db, redis)
	m := Di.DICommonMiddleware(db, redis)

	route.POST("/register", authController.Register)
	route.POST("/login/owner", authController.LoginOwner)
	route.POST("/login/employee", authController.LoginEmployee)
	route.POST("/login", authController.LoginMobile)
	route.GET("/profile", m.IsAuthenticate, authController.Profile)

}
