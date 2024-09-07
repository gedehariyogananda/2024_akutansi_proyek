package Controllers

import (
	"net/http"

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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user, err, statusCode := c.service.Register(&registerRequest)

	if err != nil {
		ctx.JSON(statusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(statusCode, gin.H{
		"success": true,
		"message": "mantap",
		"data":    user,
	})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginRequest Dto.LoginRequest

	if err := ctx.ShouldBind(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user, token, err, statusCode := c.service.Login(&loginRequest)
	if err != nil {
		ctx.JSON(statusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(statusCode, gin.H{
		"success": true,
		"message": "successfully login",
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	})

}

func (c *AuthController) UpdateTokenCompany(ctx *gin.Context) {

	authorizeUserID := ctx.GetInt("user_id")

	var request Dto.TokenCompanyRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	token, company, err, statusCode := c.service.TokenCompany(&request, authorizeUserID)
	if err != nil {
		ctx.JSON(statusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(statusCode, gin.H{
		"success": true,
		"message": "success update token",
		"data": gin.H{
			"token":        token,
			"company_user": company,
		},
	})
}
