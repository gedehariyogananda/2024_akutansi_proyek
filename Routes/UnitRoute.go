package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Unit(c *gin.RouterGroup, db *gorm.DB) {
	route := c.Group("/unit")

	m := Di.DICommonMiddleware(db)

	// open use authenticate
	route.Use(m.IsAuthenticate)

	UnitController := Di.DIUnit(db)

	route.POST("/", UnitController.Create)
	// route.GET("/", UnitController.GetAllUnit)
	route.PATCH("/", UnitController.Update)
	route.DELETE("/", UnitController.Delete)
}
