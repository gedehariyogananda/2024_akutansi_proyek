package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CompanyRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/company")

	m := Di.DICommonMiddleware(db, redis)

	route.Use(m.IsAuthenticate)

	company := Di.DiCompany(db)

	route.GET("/", company.GetDetailCompany)

}
