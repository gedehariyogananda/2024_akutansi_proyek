package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"

	"github.com/gin-gonic/gin"
)

type (
	IAuthController interface {
		Register(ctx *gin.Context)
		LoginOwner(ctx *gin.Context)
		LoginEmployee(ctx *gin.Context)
	}

	AuthController struct {
		service Services.IAuthService
	}
)

func AuthControllerProvider(service Services.IAuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var registerRequest Dto.RegisterRequest

	if err := ctx.ShouldBind(&registerRequest); err != nil {
		Helper.SetValidationErrorResponse(ctx, err.Error())
		return
	}

	user, statusCode, err := c.service.Register(&registerRequest)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Register Successful", gin.H{
		"user": user,
	}, statusCode)
}
func (c *AuthController) LoginOwner(ctx *gin.Context) {
	var loginOwnerDTO Dto.LoginOwnerRequest

	if err := ctx.ShouldBind(&loginOwnerDTO); err != nil {
		Helper.SetValidationErrorResponse(ctx, err.Error())
		return
	}

	token, statusCode, err := c.service.LoginOwner(&loginOwnerDTO)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Login Owner Successful", gin.H{
		"token": token,
	}, statusCode)
}

func (c *AuthController) LoginEmployee(ctx *gin.Context) {
	var loginEmployeeDTO Dto.LoginEmployeeRequest

	if err := ctx.ShouldBind(&loginEmployeeDTO); err != nil {
		Helper.SetValidationErrorResponse(ctx, err.Error())
		return
	}

	token, statusCode, err := c.service.LoginEmployee(&loginEmployeeDTO)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Login Employee Successful", gin.H{
		"token": token,
	}, statusCode)
}
