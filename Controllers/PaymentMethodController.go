package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	IPaymentMethodController interface {
		FindAllPaymentMethod(ctx *gin.Context)
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
