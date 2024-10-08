package Controllers

import (
	"2024_akutansi_project/Config"
	"2024_akutansi_project/Helper"
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
	fileHeader, err := ctx.FormFile("image_company")
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": "Image company is required",
		}, http.StatusBadRequest)
		return
	}

	if fileHeader.Size > int64(Config.MaxMultipartMemory) {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": "Image size too large",
		}, http.StatusBadRequest)
		return
	}

	userID := ctx.GetString("user_id")

	var request Dto.MakeCompanyRequest
	if err := ctx.ShouldBind(&request); err != nil { // shouldBind should, not shouldBindJSON
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	company, err, filePath, statusCode := c.companyService.AddCompany(&request, userID, fileHeader.Filename)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	successUpload := ctx.SaveUploadedFile(fileHeader, filePath)

	if successUpload != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": "Failed to upload image",
		}, http.StatusInternalServerError)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success make companies",
		"data":    company,
	}, http.StatusOK)
}

func (c *CompanyController) GetAllCompanyUser(ctx *gin.Context) {
	userId := ctx.GetString("user_id")

	companies, err, statusCode := c.companyService.GetAllCompanyUser(userId)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success":   true,
		"message":   "Success get all companies",
		"companies": companies,
	}, http.StatusOK)
}
func (c *CompanyController) UpdateCompany(ctx *gin.Context) {
	userId := ctx.GetString("user_id")
	companyId := ctx.GetString("company_id")
	var request Dto.EditCompanyRequest
	if err := ctx.ShouldBind(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	formHandling, err := ctx.FormFile("image_company")
	var fileName string
	if err == nil {
		if formHandling.Size > int64(Config.MaxMultipartMemory) {
			Helper.SetResponse(ctx, gin.H{
				"success": false,
				"message": "Image size too large",
			}, http.StatusBadRequest)
			return
		}
		fileName = formHandling.Filename
	}

	company, statusCode, filePath, err := c.companyService.UpdateCompany(&request, companyId, userId, fileName)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	if formHandling != nil {
		if err := ctx.SaveUploadedFile(formHandling, filePath); err != nil {
			Helper.SetResponse(ctx, gin.H{
				"success": false,
				"message": "Failed to upload image",
			}, http.StatusInternalServerError)
			return
		}
	}

	company.ID = companyId

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success update company",
		"data":    company,
	}, statusCode)
}

func (c *CompanyController) DeleteCompany(ctx *gin.Context) {
	userId := ctx.GetString("user_id")
	companyId := ctx.GetString("company_id")

	statusCode, err := c.companyService.DeleteCompany(companyId, userId)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success delete company",
	}, statusCode)
}
