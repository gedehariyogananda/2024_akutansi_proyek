package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	IPromoController interface {
		Create(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		Delete(ctx *gin.Context)
		Update(ctx *gin.Context)
	}

	PromoController struct {
		PromoService Services.IPromoService
	}
)

func PromoControllerProvider(promoService Services.IPromoService) *PromoController {
	return &PromoController{PromoService: promoService}
}

func (controller *PromoController) Create(ctx *gin.Context) {
	var request Dto.CreatePromoDto

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetValidationErrorResponse(ctx, err.Error())
		return
	}
	request.CompanyID = ctx.GetString("company_id")
	res, err := controller.PromoService.Create(&request)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}
	Helper.SetSuccessResponse(ctx, "Success create promo", res, http.StatusCreated)
}

func (controller *PromoController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	res, statusCode, err := controller.PromoService.FindByID(id)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}
	Helper.SetSuccessResponse(ctx, "Success get promo", res, http.StatusOK)
}

func (controller *PromoController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	statusCode, err := controller.PromoService.Delete(id)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}
	Helper.SetSuccessResponse(ctx, "Success delete promo", nil, http.StatusOK)
}

func (controller *PromoController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var request Dto.UpdatePromoDto

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetValidationErrorResponse(ctx, err.Error())
		return
	}
	request.CompanyID = ctx.GetString("company_id")
	res, statusCode, err := controller.PromoService.Update(&request, id)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}
	Helper.SetSuccessResponse(ctx, "Success update promo", res, http.StatusOK)
}
