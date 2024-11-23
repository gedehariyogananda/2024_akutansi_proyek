package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ICategoryController interface {
		FindAllCategory(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		FindById(ctx *gin.Context)
		Delete(ctx *gin.Context)
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

func (c *CategoryController) Create(ctx *gin.Context) {
	var createCategoryDto Dto.CreateCategory

	if err := ctx.ShouldBindJSON(&createCategoryDto); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	companyId := ctx.GetString("company_id")
	createCategoryDto.CompanyID = companyId

	category, statusCode, err := c.CategoryService.Create(&createCategoryDto, companyId)

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

func (c *CategoryController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateCategoryDto Dto.UpdateCategory

	if err := ctx.ShouldBindJSON(&updateCategoryDto); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	category, statusCode, err := c.CategoryService.Update(&updateCategoryDto, id)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success update category",
		"data":    category,
	}, statusCode)

	return
}

func (c *CategoryController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, statusCode, err := c.CategoryService.FindByID(id)
	fmt.Println(err)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": "Category not found",
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success get single category",
		"data":    res,
	}, statusCode)
	return
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	paramId := ctx.Param("id")

	statusCode, err := c.CategoryService.Delete(paramId)

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
