package Routes

import (
	"2024_akutansi_project/Dependencies"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.Engine, deps *Dependencies.Dependency) {

	apiPrefix := c.Group("/api/v1/")

	AuthRoute(apiPrefix, deps.DB, deps.Redis)
	InvoiceRoute(apiPrefix, deps.DB, deps.Redis)
	CategoryRoute(apiPrefix, deps.DB, deps.Redis)
	ProfileRoute(apiPrefix, deps.DB, deps.Mongo)
	Unit(apiPrefix, deps.DB, deps.Redis)
	TaxRoute(apiPrefix, deps.DB, deps.Redis)
	SellableProductRoutes(apiPrefix, deps.DB, deps.Redis)
}
