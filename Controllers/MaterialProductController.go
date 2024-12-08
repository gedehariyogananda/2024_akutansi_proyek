package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	IMaterialProductController interface {
		Create(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		Delete(ctx *gin.Context)
		Update(ctx *gin.Context)
		FindAll(ctx *gin.Context)
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

func (controller *MaterialProductController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	res, statusCode, err := controller.MaterialProductService.FindById(id)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Success get material product", res, 200)
}

func (controller *MaterialProductController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	statusCode, err := controller.MaterialProductService.Delete(id)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Success delete material product", nil, 200)
}

func (controller *MaterialProductController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var request Dto.UpdateMaterialProductDto

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetValidationErrorResponse(ctx, err.Error())
		return
	}

	res, statusCode, err := controller.MaterialProductService.Update(&request, id)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Success update material product", res, 200)
}

func (controller *MaterialProductController) FindAll(ctx *gin.Context) {
	search := ctx.Query("search")

	limit, page := Utils.GetPaginationParams(ctx, Common.DEFAULTLIMIT, Common.DEFAULTPAGE)

	var query Common.Query

	query.Search = &search
	query.Limit = limit
	query.Page = page
	query.Limit = limit

	status := ctx.Query("status")

	if status != "" {
		statusBool, err := strconv.ParseBool(status)
		if err != nil {
			Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		}

		query.Status = statusBool
	}

	materialProducts, meta, err := controller.MaterialProductService.FindAll(ctx.GetString("company_id"), &query)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	meta = Common.PaginateMetadata(ctx, meta.TotalData, meta.Limit, meta.Page)

	Helper.SetPaginationResponse(ctx, "Success get material products", materialProducts, meta, http.StatusOK)
}
