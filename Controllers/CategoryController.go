package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ICategoryController interface {
		FindAllCategory(ctx *gin.Context)
	}

	CategoryController struct {
		CategoryService Services.ICategoryService
	}
)

func CategoryControllerProvider(categoryService Services.ICategoryService) *CategoryController {
	return &CategoryController{CategoryService: categoryService}
}

func (c *CategoryController) FindAllCategory(ctx *gin.Context) {

	companyId := ctx.GetInt("company_id")

	company, err := c.CategoryService.FindAllCategory(companyId)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
	}

	Helper.SetResponse(ctx, gin.H{
		"message": "Success get all category",
		"success": true,
		"data":    company,
	}, http.StatusOK)
}
