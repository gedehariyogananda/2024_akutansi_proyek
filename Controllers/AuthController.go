package Controllers

import (
	"net/http"

	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"

	"github.com/gin-gonic/gin"
)

type (
	IAuthController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		UpdateTokenCompany(ctx *gin.Context)
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
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	user, err, statusCode := c.service.Register(&registerRequest)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Registration successful",
		"data":    user,
	}, statusCode)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginRequest Dto.LoginRequest

	if err := ctx.ShouldBind(&loginRequest); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	user, token, err, statusCode := c.service.Login(&loginRequest)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Login successful",
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	}, statusCode)
}

func (c *AuthController) UpdateTokenCompany(ctx *gin.Context) {
	authorizeUserID := ctx.GetInt("user_id")

	var request Dto.TokenCompanyRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": "Invalid request body",
		}, http.StatusBadRequest)
		return
	}

	token, company, err, statusCode := c.service.TokenCompany(&request, authorizeUserID)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Token updated successfully",
		"data": gin.H{
			"token":        token,
			"company_user": company,
		},
	}, statusCode)
}
