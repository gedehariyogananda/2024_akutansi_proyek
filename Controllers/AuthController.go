package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"

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
	var registerDTO Dto.RegisterRequest

	if err := ctx.ShouldBind(&registerDTO); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	if validationErrors := Utils.ValidateRequest(ctx, &registerDTO); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	user, statusCode, err := c.service.Register(&registerDTO)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Register berhasil!", gin.H{
		"user": user,
	}, statusCode)
}
func (c *AuthController) LoginOwner(ctx *gin.Context) {
	var loginOwnerDTO Dto.LoginOwnerRequest

	if err := ctx.ShouldBind(&loginOwnerDTO); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	if validationErrors := Utils.ValidateRequest(ctx, &loginOwnerDTO); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	token, statusCode, err := c.service.LoginOwner(ctx.Request.Context(), &loginOwnerDTO)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Login Berhasil!", gin.H{
		"token": token,
	}, statusCode)
}

func (c *AuthController) LoginEmployee(ctx *gin.Context) {
	var loginEmployeeDTO Dto.LoginEmployeeRequest

	if err := ctx.ShouldBind(&loginEmployeeDTO); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	if validationErrors := Utils.ValidateRequest(ctx, &loginEmployeeDTO); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	token, statusCode, err := c.service.LoginEmployee(ctx.Request.Context(), &loginEmployeeDTO)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Login Berhasil!", gin.H{
		"token": token,
	}, statusCode)
}
