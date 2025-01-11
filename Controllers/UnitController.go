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
	IUnitController interface {
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
		FindAll(ctx *gin.Context)
		FindByID(ctx *gin.Context)
	}

	UnitController struct {
		unitService Services.IUnitService
	}
)

func UnitProvider(unitService Services.IUnitService) *UnitController {
	return &UnitController{unitService: unitService}
}

func (c *UnitController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, statusCode, err := c.unitService.FindByID(id)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Success het single satuan", res, http.StatusOK)
}

func (c *UnitController) Create(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")
	var request Dto.CreateUnitDto

	request.CompanyID = companyId

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan input data", http.StatusBadRequest)
		return
	}

	if validationErrors := Utils.ValidateRequest(ctx, request); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	res, err := c.unitService.Create(&request)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	Helper.SetSuccessResponse(ctx, "Success create satuan", res, http.StatusOK)
}

func (c *UnitController) Update(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")
	id := ctx.Param("id")

	var request Dto.UpdateUnitDto
	request.CompanyID = companyId

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan input data", http.StatusBadRequest)
		return
	}

	if validatioErrors := Utils.ValidateRequest(ctx, request); validatioErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validatioErrors)
		return
	}

	res, statusCode, err := c.unitService.Update(&request, id)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Success update satuan", res, http.StatusOK)
}

func (c *UnitController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err, statusCode := c.unitService.Delete(id)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Success delete satuan", nil, http.StatusOK)
}

func (c *UnitController) FindAll(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")
	search := ctx.Query("search")
	status := ctx.Query("status")

	limit, page := Utils.GetPaginationParams(ctx, Common.DEFAULTLIMIT, Common.DEFAULTPAGE)

	var query Common.Query

	if status != "" {
		status, err := strconv.ParseBool(status)

		if err != nil {
			Helper.SetErrorResponse(ctx, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		query.Status = &status
	}

	query.Search = &search
	query.Limit = limit
	query.Page = page
	query.Limit = limit

	res, meta, err := c.unitService.FindAll(companyId, &query)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	meta = Common.PaginateMetadata(ctx, meta.TotalData, limit, page)

	Helper.SetPaginationResponse(ctx, "Success get all data satuan", res, meta, http.StatusOK)
}
