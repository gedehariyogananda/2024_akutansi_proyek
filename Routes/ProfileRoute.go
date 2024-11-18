package Routes

import (
	"2024_akutansi_project/Routes/Di"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func ProfileRoute(c *gin.RouterGroup, db *gorm.DB, mongo *mongo.Client) {
	route := c.Group("/profile")

	Controller := Di.DIProfile(db, mongo)

	// todo :: use middleware

	route.POST("/integrate/shopee", Controller.IntegrateShopeeProfile)
	route.GET("/detail", Controller.GetIntegratedProfile)
}
