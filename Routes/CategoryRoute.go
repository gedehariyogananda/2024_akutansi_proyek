package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRoute(c *gin.RouterGroup, db *gorm.DB) {
	route := c.Group("/category")

	m := Di.DICommonMiddleware(db)

	// open use authenticate
	route.Use(m.IsAuthenticate)

	CategoryController := Di.DICategory(db)

	route.GET("/", CategoryController.FindAllCategory)

	route.POST("/", CategoryController.CreateCategory)
	route.PUT("/:id", CategoryController.UpdateCategory)
	route.DELETE("/:id", CategoryController.DeleteCategory)
}
