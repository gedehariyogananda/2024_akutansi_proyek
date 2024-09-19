package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ICategoryController interface {
		FindAllCategory(ctx *gin.Context)
		CreateCategory(ctx *gin.Context)
		UpdateCategory(ctx *gin.Context)
		DeleteCategory(ctx *gin.Context)
	}

	CategoryController struct {
		CategoryService Services.ICategoryService
	}
)

func CategoryControllerProvider(categoryService Services.ICategoryService) *CategoryController {
	return &CategoryController{CategoryService: categoryService}
}

func (c *CategoryController) FindAllCategory(ctx *gin.Context) {

	companyId := ctx.GetString("company_id")

	company, err := c.CategoryService.FindAllCategory(companyId)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"message": "Success get all category",
		"success": true,
		"data":    company,
	}, http.StatusOK)
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")

	var request Dto.CreateCategoryRequestDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	category, statusCode, err := c.CategoryService.CreateCategory(&request, companyId)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success create category",
		"data":    category,
	}, http.StatusOK)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	paramId := ctx.Param("id")
	companyId := ctx.GetString("company_id")

	var request Dto.UpdateCategoryRequestDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	category, statusCode, err := c.CategoryService.UpdateCategory(&request, paramId, companyId)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	category.ID = paramId

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success update category",
		"data":    category,
	}, http.StatusOK)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	paramId := ctx.Param("id")

	statusCode, err := c.CategoryService.DeleteCategory(paramId)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success delete category",
	}, http.StatusOK)
}
