package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func JournalEntriesRoutes(c *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	route := c.Group("/journal-entries")

	m := Di.DICommonMiddleware(db, redis)

	route.Use(m.IsAuthenticate)

	controller := Di.DIJournalEntries(db)

	route.GET("/", controller.FindAll)
	route.GET("/trial-balance", controller.TrialBalanceReport)
	route.GET("/financial-balance", controller.FinancialBalanceReport)
	route.GET("/profit-or-loss", controller.ProfitOrLossReport)
}
