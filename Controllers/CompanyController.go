package Controllers

import (
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ICompanyController interface {
		AddCompany(ctx *gin.Context)
		GetAllCompanyUser(ctx *gin.Context)
	}

	CompanyController struct {
		companyService Services.ICompanyService
	}
)

func CompanyControllerProvider(companyService Services.ICompanyService) *CompanyController {
	return &CompanyController{
		companyService: companyService,
	}
}

func (c *CompanyController) AddCompany(ctx *gin.Context) {
	authorizeUserID, _ := ctx.Get("user_id")
	userID := authorizeUserID.(float64)

	var request Dto.MakeCompanyRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	err := c.companyService.AddCompany(&request, int(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success make companies"})
}

func (c *CompanyController) GetAllCompanyUser(ctx *gin.Context) {
	authorizeUserID, _ := ctx.Get("user_id")
	userID := authorizeUserID.(float64)

	companies, err := c.companyService.GetAllCompanyUser(int(userID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":   true,
		"message":   "success get all companies",
		"companies": companies,
	})
}
