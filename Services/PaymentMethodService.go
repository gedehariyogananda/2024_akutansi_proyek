package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"fmt"
	"net/http"
)

type (
	IPaymentMethodService interface {
		FindAllPaymentMethod(company_id string) (paymentMethod *[]Models.PaymentMethod, err error)
		CreatePaymentMethod(request *Dto.CreatePaymentMethodRequestDTO, company_id string) (paymentMethod *Models.PaymentMethod, statusCode int, err error)
		UpdatePaymentMethod(request *Dto.UpdatePaymentMethodRequestDTO, id string, company_id string) (paymentMethod *Models.PaymentMethod, statusCode int, err error)
		DeletePaymentMethod(id string) (statusCode int, err error)
	}

	PaymentMethodService struct {
		PaymentMethodRepository Repositories.IPaymentMethodRepository
	}
)

func PaymentMethodServiceProvider(paymentMethodRepository Repositories.IPaymentMethodRepository) *PaymentMethodService {
	return &PaymentMethodService{PaymentMethodRepository: paymentMethodRepository}
}

func (service *PaymentMethodService) FindAllPaymentMethod(company_id string) (paymentMethod *[]Models.PaymentMethod, err error) {
	paymentMethod, err = service.PaymentMethodRepository.FindAll(company_id)
	if err != nil {
		return nil, err
	}
	return paymentMethod, nil
}

func (service *PaymentMethodService) CreatePaymentMethod(request *Dto.CreatePaymentMethodRequestDTO, company_id string) (paymentMethod *Models.PaymentMethod, statusCode int, err error) {
	paymentMethod, err = service.PaymentMethodRepository.Create(request, company_id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return paymentMethod, http.StatusOK, nil
}

func (service *PaymentMethodService) UpdatePaymentMethod(request *Dto.UpdatePaymentMethodRequestDTO, id string, company_id string) (paymentMethod *Models.PaymentMethod, statusCode int, err error) {
	paymentMethodInit, err := service.PaymentMethodRepository.FindById(id)

	if err != nil {
		return nil, http.StatusNotFound, fmt.Errorf("payment method not found")
	}

	if paymentMethodInit.MethodName == "Cash" {
		return nil, http.StatusBadRequest, fmt.Errorf("cannot update CASH default payment method")
	}

	paymentMethod, err = service.PaymentMethodRepository.Update(request, id, company_id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return paymentMethod, http.StatusOK, nil
}

func (service *PaymentMethodService) DeletePaymentMethod(id string) (statusCode int, err error) {
	err = service.PaymentMethodRepository.Delete(id)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
