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
		UpdateCompany(ctx *gin.Context)
		DeleteCompany(ctx *gin.Context)
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
	userID := ctx.GetInt("user_id")

	var request Dto.MakeCompanyRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	err := c.companyService.AddCompany(&request, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success make companies"})
}

func (c *CompanyController) GetAllCompanyUser(ctx *gin.Context) {

	userId := ctx.GetInt("user_id")

	companies, err := c.companyService.GetAllCompanyUser(userId)

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

func (c *CompanyController) UpdateCompany(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	companyId := ctx.GetInt("company_id")

	var request Dto.EditCompanyRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	company, statusCode, err := c.companyService.UpdateCompany(&request, companyId, userId)
	if err != nil {
		ctx.JSON(statusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	company.ID = companyId

	ctx.JSON(statusCode, gin.H{
		"success": true,
		"message": "success update company",
		"data":    company,
	})
}

func (c *CompanyController) DeleteCompany(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	companyId := ctx.GetInt("company_id")

	statusCode, err := c.companyService.DeleteCompany(companyId, userId)
	if err != nil {
		ctx.JSON(statusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(statusCode, gin.H{
		"success": true,
		"message": "success delete company",
	})
}
