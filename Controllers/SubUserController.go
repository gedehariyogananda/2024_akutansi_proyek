package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ISubUserController interface {
		Create(c *gin.Context)
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
