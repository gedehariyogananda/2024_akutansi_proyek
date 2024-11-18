package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IProfileController interface {
		IntegrateShopeeProfile(ctx *gin.Context)
		GetIntegratedProfile(ctx *gin.Context)
	}

	ProfileController struct {
		ProfileService Services.IProfileService
	}
)

func ProfileControllerProvider(ProfileService Services.IProfileService) *ProfileController {
	return &ProfileController{
		ProfileService: ProfileService,
	}
}

func (controller *ProfileController) IntegrateShopeeProfile(ctx *gin.Context) {
	var request Dto.ShopeeIntegrateRequest

	ctx.Set("user_id", "123")
	userID, _ := ctx.Get("user_id")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "user_id", userID))

	statusCode, err := controller.ProfileService.IntegrateShopeeProfile(ctx.Request.Context(), &request)
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success",
	}, statusCode)
}

func (controller *ProfileController) GetIntegratedProfile(ctx *gin.Context) {
	ctx.Set("user_id", "123")
	userID, _ := ctx.Get("user_id")

	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "user_id", userID))
	data, statusCode, err := controller.ProfileService.GetIntegratedProfile(ctx.Request.Context())
	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success",
		"data":    data,
	}, statusCode)
}
