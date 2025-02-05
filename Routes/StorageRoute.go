package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func StorageRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client, minio *minio.Client) {
	route := c.Group("/storage")

	m := Di.DICommonMiddleware(db, redis)

	route.Use(m.RolesAll)

	StorageController := Di.DIStorage(minio)

	route.GET("/signed-url", StorageController.GetSignedUrl)

}
