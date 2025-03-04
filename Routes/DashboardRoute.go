package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func DashboardRoute(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/dashboard")
	m := Di.DICommonMiddleware(db, redis)

	route.Use(m.RolesAll)

	DashboardController := Di.DiDashboard(db)

	route.GET("/sales-resume", DashboardController.GetSalesResume)
	route.GET("/best-sales", DashboardController.GetBestSalesProduct)
	route.GET("/revenue-chart", DashboardController.GetRevenue)
	route.GET("/expense-chart", DashboardController.GetExpense)
}
