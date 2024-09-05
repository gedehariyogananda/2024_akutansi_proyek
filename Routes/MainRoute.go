package Routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(c *gin.Engine, db *gorm.DB) {

	apiPrefix := c.Group("/api/v1/")

	// Initialize routes
	AuthRoute(apiPrefix, db)
	CompanyRoute(apiPrefix, db)
	SaleableProductRoute(apiPrefix, db)
	InvoiceRoute(apiPrefix, db)
}
