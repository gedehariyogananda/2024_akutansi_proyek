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

	user, err := c.service.Register(&registerRequest)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"message": "mantap",
		"data":    user,
	})
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Param body body Dto.LoginRequest true "Login Request"
// @Success 202 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/login [post]

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

	ctx.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"message": "successfully login",
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	})

}
