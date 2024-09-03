package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoute(c *gin.RouterGroup, db *gorm.DB) {
	route := c.Group("/auth")

	authController := Di.DIAuth(db)

	route.GET("/checked", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "checked healt",
		})
	})

	route.POST("/register", authController.Register)
	route.POST("/login", authController.Login)
}
