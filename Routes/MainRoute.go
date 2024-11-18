package Routes

import (
	"2024_akutansi_project/Dependencies"
	"github.com/gin-gonic/gin"
)

func Init(c *gin.Engine, deps *Dependencies.Dependency) {

	apiPrefix := c.Group("/api/v1/")

	// Initialize routes
	AuthRoute(apiPrefix, deps.DB)
	CompanyRoute(apiPrefix, deps.DB)
	SaleableProductRoute(apiPrefix, deps.DB)
	InvoiceRoute(apiPrefix, deps.DB)
	CategoryRoute(apiPrefix, deps.DB)
	PaymentMethodRoute(apiPrefix, deps.DB)

	ProfileRoute(apiPrefix, deps.DB, deps.Mongo)
}
