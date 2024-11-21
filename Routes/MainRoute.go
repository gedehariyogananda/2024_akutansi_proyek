package Routes

import (
	Dependencies "2024_akutansi_project/Depedencies"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.Engine, deps *Dependencies.Dependency) {

	apiPrefix := c.Group("/api/v1/")

	// Initialize routes
	AuthRoute(apiPrefix, deps.DB, deps.Redis)
	CompanyRoute(apiPrefix, deps.DB, deps.Redis)
	SaleableProductRoute(apiPrefix, deps.DB, deps.Redis)
	InvoiceRoute(apiPrefix, deps.DB, deps.Redis)
	CategoryRoute(apiPrefix, deps.DB, deps.Redis)
	PaymentMethodRoute(apiPrefix, deps.DB, deps.Redis)
}
