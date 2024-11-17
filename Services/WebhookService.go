package Services

import (
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"context"
	"net/http"
)

type (
	IWebhookService interface {
		ShopeeWebhook(ctx context.Context, request *Dto.ShopeeCallbackRequest) (statusCode int, err error)
		GetWebhookByID(ctx context.Context, id string) (resp *Dto.CallbackProfile, statusCode int, err error)
	}

	WebhookService struct {
		WebhookRepository Repositories.IWebhookRepository
	}
)

func WebhookServiceProvider(paymentMethodRepository Repositories.IWebhookRepository) *WebhookService {
	return &WebhookService{WebhookRepository: paymentMethodRepository}
}

func (service *WebhookService) ShopeeWebhook(ctx context.Context, request *Dto.ShopeeCallbackRequest) (statusCode int, err error) {
	err = service.WebhookRepository.CreateWebhookProduct(ctx, Dto.CallbackProfile{
		ID:       request.ID,
		Price:    request.Price,
		Provider: "shopee",
	})
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

func (service *WebhookService) GetWebhookByID(ctx context.Context, id string) (resp *Dto.CallbackProfile, statusCode int, err error) {
	resp, err = service.WebhookRepository.GetWebhookByID(ctx, id)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	return resp, http.StatusOK, nil
}
