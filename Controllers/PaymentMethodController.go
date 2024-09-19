package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	IPaymentMethodController interface {
		FindAllPaymentMethod(ctx *gin.Context)
		CreatePaymentMethod(ctx *gin.Context)
		UpdatePaymentMethod(ctx *gin.Context)
		DeletePaymentMethod(ctx *gin.Context)
	}

	PaymentMethodController struct {
		PaymentMethodService Services.IPaymentMethodService
	}
)

func PaymentMethodControllerProvider(PaymentMethodService Services.IPaymentMethodService) *PaymentMethodController {
	return &PaymentMethodController{
		PaymentMethodService: PaymentMethodService,
	}
}

func (controller *PaymentMethodController) FindAllPaymentMethod(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")

	paymentMethod, err := controller.PaymentMethodService.FindAllPaymentMethod(companyId)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success get payment methods",
		"data":    paymentMethod,
	}, http.StatusOK)

}

func (controller *PaymentMethodController) CreatePaymentMethod(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")

	var request Dto.CreatePaymentMethodRequestDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	paymentMethod, statusCode, err := controller.PaymentMethodService.CreatePaymentMethod(&request, companyId)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success create payment method",
		"data":    paymentMethod,
	}, statusCode)
}

func (controller *PaymentMethodController) UpdatePaymentMethod(ctx *gin.Context) {
	companyId := ctx.GetString("company_id")
	paymentMethodId := ctx.Param("id")

	var request Dto.UpdatePaymentMethodRequestDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	paymentMethod, statusCode, err := controller.PaymentMethodService.UpdatePaymentMethod(&request, paymentMethodId, companyId)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	paymentMethod.ID = paymentMethodId

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success update payment method",
		"data":    paymentMethod,
	}, statusCode)
}

func (controller *PaymentMethodController) DeletePaymentMethod(ctx *gin.Context) {
	paymentMethodId := ctx.Param("id")

	statusCode, err := controller.PaymentMethodService.DeletePaymentMethod(paymentMethodId)

	if err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, statusCode)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Success delete payment method",
	}, statusCode)
}
