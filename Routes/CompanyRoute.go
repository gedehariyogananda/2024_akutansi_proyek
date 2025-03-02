package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CompanyRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client, minio *minio.Client) {
	route := c.Group("/company")

	m := Di.DICommonMiddleware(db, redis)

	route.Use(m.RolesAll)

	company := Di.DiCompany(db, minio)

	route.GET("/", company.GetDetailCompany)

}
