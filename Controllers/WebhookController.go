package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IWebhookController interface {
		ShopeeWebhook(ctx *gin.Context)
		GetByID(ctx *gin.Context)
	}

	WebhookController struct {
		webhookService Services.IWebhookService
	}
)

func WebhookControllerProvider(webhookService Services.IWebhookService) *WebhookController {
	return &WebhookController{
		webhookService: webhookService,
	}
}

func (controller *WebhookController) ShopeeWebhook(ctx *gin.Context) {
	var request Dto.ShopeeCallbackRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	statusCode, err := controller.webhookService.ShopeeWebhook(ctx.Request.Context(), &request)
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

func (controller *WebhookController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	data, statusCode, err := controller.webhookService.GetWebhookByID(ctx.Request.Context(), id)
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
