package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoute(c *gin.RouterGroup, db *gorm.DB) {
	route := c.Group("/auth")

	authController := Di.DIAuth(db)

	route.POST("/register", authController.Register)
	route.POST("/login/owner", authController.LoginOwner)
	route.POST("/login/employee", authController.LoginEmployee)

}
