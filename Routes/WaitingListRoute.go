package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WaitingListRoute(c *gin.RouterGroup, db *gorm.DB) {
	route := c.Group("/waiting-list")

	waitingListController := Di.DIWaitingList(db)

	route.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "checked health",
		})
	})

	route.POST("/", waitingListController.Insert)
}
