package Routes

import (
	"2024_akutansi_project/Routes/Di"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func WebhookRoute(c *gin.RouterGroup, db *gorm.DB, mongo *mongo.Client) {
	route := c.Group("/webhook")

	Controller := Di.DIWebhook(db, mongo)

	route.POST("/shopee", Controller.ShopeeWebhook)
	route.GET("/detail/:id", Controller.GetByID)
}
