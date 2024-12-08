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
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success get unit",
		"data":    res,
	}, statusCode)
}

func (c *UnitController) Create(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")
	var request Dto.CreateUnitDto

	request.CompanyID = companyId

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	res, err := c.unitService.Create(&request)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"message": "Success create unit",
		"success": true,
		"data":    res,
	}, http.StatusOK)

}

func (c *UnitController) Update(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")
	id := ctx.Param("id")

	var request Dto.UpdateUnitDto
	request.CompanyID = companyId

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	res, statusCode, err := c.unitService.Update(&request, id)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success update unit",
		"data":    res,
	}, statusCode)
}

func (c *UnitController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err, statusCode := c.unitService.Delete(id)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success delete unit",
	}, statusCode)
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
			Helper.SetResponse(ctx, gin.H{
				"success": false,
				"message": "Invalid status",
			}, http.StatusBadRequest)
			return
		}

		query.Status = status
	}

	query.Search = &search
	query.Limit = limit
	query.Page = page
	query.Limit = limit

	res, meta, err := c.unitService.FindAll(companyId, &query)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	meta = Common.PaginateMetadata(ctx, meta.TotalData, limit, page)

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success get all unit",
		"data":    res,
		"meta":    meta,
	}, http.StatusOK)
}
