package Routes

import (
	Dependencies "2024_akutansi_project/Depedencies"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.Engine, deps *Dependencies.Dependency) {

	apiPrefix := c.Group("/api/v1/")

	AuthRoute(apiPrefix, deps.DB)
	CompanyRoute(apiPrefix, deps.DB)
	SaleableProductRoute(apiPrefix, deps.DB)
	InvoiceRoute(apiPrefix, deps.DB)
	CategoryRoute(apiPrefix, deps.DB)
	PaymentMethodRoute(apiPrefix, deps.DB)
	ProfileRoute(apiPrefix, deps.DB, deps.Mongo)
	Unit(apiPrefix, deps.DB, deps.Redis)
}
