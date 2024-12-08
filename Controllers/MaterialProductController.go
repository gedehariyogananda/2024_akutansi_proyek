package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"

	"github.com/gin-gonic/gin"
)

type (
	IMaterialProductController interface {
		Create(ctx *gin.Context)
	}
	MaterialProductController struct {
		MaterialProductService Services.IMaterialProductService
	}
)

func MaterialProductControllerProvider(materialProductService Services.IMaterialProductService) *MaterialProductController {
	return &MaterialProductController{MaterialProductService: materialProductService}
}

func (controller *MaterialProductController) Create(ctx *gin.Context) {
	companyID := ctx.GetString("company_id")
	var request Dto.CreateMaterialProductDto

	request.CompanyID = companyID

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetValidationErrorResponse(ctx, err.Error())
		return
	}

	res, err := controller.MaterialProductService.Create(&request)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), 400)
		return
	}

	Helper.SetSuccessResponse(ctx, "Success create material product", res, 201)
}
