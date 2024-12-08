package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	ICategoryController interface {
		FindAll(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	CategoryController struct {
		CategoryService Services.ICategoryService
	}
)

func CategoryControllerProvider(categoryService Services.ICategoryService) *CategoryController {
	return &CategoryController{CategoryService: categoryService}
}

func (c *CategoryController) FindAll(ctx *gin.Context) {

	companyId := ctx.GetString("company_id")

	var query Common.Query

	limit, page := Utils.GetPaginationParams(ctx, Common.DEFAULTLIMIT, Common.DEFAULTPAGE)

	search := ctx.Query("search")

	query.Limit = limit
	query.Page = page
	query.Search = &search

	status := ctx.Query("status")

	if status != "" {
		status, err := strconv.ParseBool(status)

		if err != nil {
			Helper.SetResponse(ctx, gin.H{
				"success": false,
				"message": "Invalid status",
			}, http.StatusBadRequest)
			return
		}

		query.Status = status
	}

	res, meta, err := c.CategoryService.FindAll(companyId, &query)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	meta = Common.PaginateMetadata(ctx, meta.TotalData, limit, page)

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success get all category",
		"data":    res,
		"meta":    meta,
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
}

func (c *CategoryController) Delete(ctx *gin.Context) {
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
