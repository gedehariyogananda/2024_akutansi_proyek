package Routes

import (
	"2024_akutansi_project/Routes/Di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CompanyRoute(c *gin.RouterGroup, db *gorm.DB) {
	route := c.Group("/company")

	m := Di.DICommonMiddleware(db)

	// open use authenticate
	route.Use(m.IsAuthenticate)

	CompanyController := Di.DICompany(db)

	route.POST("/", CompanyController.AddCompany)
	route.GET("/", CompanyController.GetAllCompanyUser)
	route.PATCH("/", CompanyController.UpdateCompany)
	route.DELETE("/", CompanyController.DeleteCompany)
}
