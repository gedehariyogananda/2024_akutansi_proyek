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
	IAccountController interface {
		Create(c *gin.Context)
		FindByID(c *gin.Context)
		Update(c *gin.Context)
		Delete(c *gin.Context)
		FindAll(c *gin.Context)
	}

	AccountController struct {
		accountService Services.IAccountService
	}
)

func AccountProvider(accountService Services.IAccountService) *AccountController {
	return &AccountController{accountService: accountService}
}

func (controller *AccountController) Create(c *gin.Context) {
	var request Dto.CreateAccountDto
	if err := c.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	request.CompanyID = c.GetString("company_id")
	res, err := controller.accountService.Create(&request)
	if err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	Helper.SetResponse(c, gin.H{
		"success": true,
		"message": "Success create account",
		"data":    res,
	}, http.StatusOK)
}

func (controller *AccountController) FindByID(c *gin.Context) {
	id := c.Param("id")
	res, statusCode, err := controller.accountService.FindByID(id)
	if err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}
	Helper.SetResponse(c, gin.H{
		"success": true,
		"message": "Success get account",
		"data":    res,
	}, http.StatusOK)
}

func (controller *AccountController) Update(c *gin.Context) {
	id := c.Param("id")
	var request Dto.UpdateAccountDto
	if err := c.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	res, statusCode, err := controller.accountService.Update(&request, id)
	if err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}
	Helper.SetResponse(c, gin.H{
		"success": true,
		"message": "Success update account",
		"data":    res,
	}, http.StatusOK)
}

func (controller *AccountController) Delete(c *gin.Context) {
	id := c.Param("id")
	statusCode, err := controller.accountService.Delete(id)
	if err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}
	Helper.SetResponse(c, gin.H{
		"success": true,
		"message": "Success delete account",
	}, http.StatusOK)
}

func (cotroller *AccountController) FindAll(c *gin.Context) {
	companyId := c.GetString("company_id")
	var query Common.Query

	limit, page := Utils.GetPaginationParams(c, Common.DEFAULTLIMIT, Common.DEFAULTPAGE)

	query.Limit = limit
	query.Page = page

	search := c.Query("search")

	query.Search = &search

	status := c.Query("status")

	fmt.Println(status)

	if status != "" {
		statusBool, err := strconv.ParseBool(status)
		if err != nil {
			Helper.SetResponse(c, gin.H{
				"success": false,
				"message": err.Error(),
			}, http.StatusBadRequest)
			return
		}

		query.Status = statusBool
	}

	isLocked := c.Query("is_locked")

	if isLocked != "" {
		isLockedBool, err := strconv.ParseBool(isLocked)
		if err != nil {
			Helper.SetResponse(c, gin.H{
				"success": false,
				"message": err.Error(),
			}, http.StatusBadRequest)
			return
		}

		query.IsLocked = isLockedBool
	}

	res, totalData, err := cotroller.accountService.FindAll(companyId, &query)
	if err != nil {
		Helper.SetResponse(c, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	Helper.SetResponse(c, gin.H{
		"success":   true,
		"message":   "Success get all account",
		"data":      res,
		"totalData": totalData,
	}, http.StatusOK)
}
