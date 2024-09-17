package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Repositories"
)

type (
	IPaymentMethodService interface {
		FindAllPaymentMethod(company_id string) (paymentMethod *[]Models.PaymentMethod, err error)
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
