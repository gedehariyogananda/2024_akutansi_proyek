package Controllers

import (
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"net/http"
	"strconv"

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

// GetAllCompanyUser godoc
// @Summary Get all companies
// @Description Get all companies users init
// @Tags Company
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} gin.H
// @Router /company [get]

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

func (c *CompanyController) UpdateCompany(ctx *gin.Context) {
	authorizeUserID, _ := ctx.Get("user_id")
	userID := authorizeUserID.(float64)
	idParam := ctx.Param("id")
	companyID, err := strconv.Atoi(idParam)

	var request Dto.EditCompanyRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	company, statusCode, err := c.companyService.UpdateCompany(&request, companyID, int(userID))
	if err != nil {
		ctx.JSON(statusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	company.ID = companyID

	ctx.JSON(statusCode, gin.H{
		"success": true,
		"message": "success update company",
		"data":    company,
	})
}

func (c *CompanyController) DeleteCompany(ctx *gin.Context) {
	authorizeUserID, _ := ctx.Get("user_id")
	userID := authorizeUserID.(float64)
	idParam := ctx.Param("id")
	companyID, err := strconv.Atoi(idParam)

	statusCode, err := c.companyService.DeleteCompany(companyID, int(userID))
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
