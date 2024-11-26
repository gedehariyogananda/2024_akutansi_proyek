package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	IAccountController interface {
		Create(c *gin.Context)
		FindByID(c *gin.Context)
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
