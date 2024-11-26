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
	ISubUserController interface {
		Create(c *gin.Context)
		Update(c *gin.Context)
		Delete(c *gin.Context)
		FindByID(c *gin.Context)
		FindAll(c *gin.Context)
	}

	SubUserController struct {
		SubUserService Services.ISubUserService
	}
)

func SubUserProvider(subUserService Services.ISubUserService) *SubUserController {
	return &SubUserController{SubUserService: subUserService}
}

func (controller *SubUserController) Create(c *gin.Context) {
	companyId := c.GetString("company_id")
	var request Dto.CreateSubUserDto
	request.CompanyID = companyId
	if err := c.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	res, err := controller.SubUserService.Create(&request)
	if err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	Helper.SetResponse(c, gin.H{
		"success": true,
		"message": "Success create sub user",
		"data":    res,
	}, http.StatusOK)
}

func (controller *SubUserController) Update(c *gin.Context) {
	id := c.Param("id")
	var request Dto.UpdateSubUserDto
	if err := c.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	res, statusCode, err := controller.SubUserService.Update(&request, id)
	if err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}
	Helper.SetResponse(c, gin.H{
		"success": true,
		"message": "Success update sub user",
		"data":    res,
	}, http.StatusOK)
}

func (controller *SubUserController) Delete(c *gin.Context) {
	id := c.Param("id")
	statusCode, err := controller.SubUserService.Delete(id)
	if err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}
	Helper.SetResponse(c, gin.H{
		"success": true,
		"message": "Success delete sub user",
	}, http.StatusOK)
}

func (controller *SubUserController) FindByID(c *gin.Context) {
	id := c.Param("id")
	res, statusCode, err := controller.SubUserService.FindByID(id)
	if err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}
	Helper.SetResponse(c, gin.H{
		"success": true,
		"message": "Success get sub user",
		"data":    res,
	}, http.StatusOK)
}

func (controller *SubUserController) FindAll(c *gin.Context) {
	companyId := c.GetString("company_id")

	var dto Common.Query

	perPage, page := Utils.GetPaginationParams(c, Common.DEFAULTLIMIT, Common.DEFAULTPAGE)

	dto.Limit = perPage
	dto.Page = page

	search := c.Query("search")
	dto.Search = &search

	status := c.Query("status")

	if status != "" {
		statusBool, err := strconv.ParseBool(status)
		if err != nil {
			Helper.SetResponse(c, gin.H{
				"success": false,
				"message": err.Error(),
			}, http.StatusBadRequest)
			return
		}

		dto.Status = statusBool
	}

	res, meta, err := controller.SubUserService.FindAll(companyId, &dto)
	if err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	meta = Common.PaginateMetadata(c, meta.TotalData, meta.Limit, meta.Page)

	Helper.SetResponse(c, gin.H{
		"success": true,
		"message": "Success get all sub user",
		"data":    res,
		"meta":    meta,
	}, http.StatusOK)
}
