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

	if err := c.service.Register(&registerRequest); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"message": "mantap",
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

	user, token, err := c.service.Login(&loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	if user.Roles == "admin" {
		ctx.JSON(http.StatusAccepted, gin.H{
			"success": true,
			"message": "successfully login",
			"data": gin.H{
				"admin": user,
				"token": token,
			},
		})
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{
			"success": true,
			"message": "successfully login",
			"data": gin.H{
				"user":  user,
				"token": token,
			},
		})

	}
}
