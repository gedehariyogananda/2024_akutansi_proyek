package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"fmt"

	"gorm.io/gorm"
)

type (
	IPaymentMethodRepository interface {
		CreateDefaultPaymentMethod(company_id string) (err error)
		FindAll(company_id string) (paymentMethod *[]Models.PaymentMethod, err error)
		FindById(id string) (paymentMethod *Models.PaymentMethod, err error)

		Create(request *Dto.CreatePaymentMethodRequestDTO, company_id string) (paymentMethod *Models.PaymentMethod, err error)
		Update(request *Dto.UpdatePaymentMethodRequestDTO, id string, company_id string) (paymentMethod *Models.PaymentMethod, err error)
		Delete(id string) (err error)
	}

	PaymentMethodRepository struct {
		DB *gorm.DB
	}
)

func PaymentMethodRepositoryProvider(db *gorm.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{DB: db}
}

func (r *PaymentMethodRepository) FindAll(company_id string) (paymentMethod *[]Models.PaymentMethod, err error) {
	paymentMethod = &[]Models.PaymentMethod{}

	if err := r.DB.Where("company_id = ?", company_id).Find(paymentMethod).Error; err != nil {
		return nil, err
	}

	return paymentMethod, nil
}

func (r *PaymentMethodRepository) CreateDefaultPaymentMethod(company_id string) (err error) {
	defaultPaymentMethod := []string{
		"Cash",
		"Credit Card",
		"Debit Card",
		"E-Wallet",
	}

	for _, methodName := range defaultPaymentMethod {
		paymentMethod := &Models.PaymentMethod{
			MethodName: methodName,
			CompanyID:  company_id,
		}

		if err := r.DB.Create(paymentMethod).Error; err != nil {
			return fmt.Errorf("failed to create payment method '%s': %w", methodName, err)
		}
	}

	return nil
}

func (r *PaymentMethodRepository) FindById(id string) (paymentMethod *Models.PaymentMethod, err error) {
	paymentMethod = &Models.PaymentMethod{}

	if err := r.DB.Where("id = ?", id).First(paymentMethod).Error; err != nil {
		return nil, err
	}

	return paymentMethod, nil
}

func (r *PaymentMethodRepository) Create(request *Dto.CreatePaymentMethodRequestDTO, company_id string) (paymentMethod *Models.PaymentMethod, err error) {
	paymentMethod = &Models.PaymentMethod{
		MethodName: request.MethodName,
		CompanyID:  company_id,
	}

	if err := r.DB.Create(paymentMethod).Error; err != nil {
		return nil, err
	}

	return paymentMethod, nil
}

func (r *PaymentMethodRepository) Update(request *Dto.UpdatePaymentMethodRequestDTO, id string, company_id string) (paymentMethod *Models.PaymentMethod, err error) {
	paymentMethod = &Models.PaymentMethod{
		MethodName: request.MethodName,
		CompanyID:  company_id,
	}

	if err := r.DB.Model(paymentMethod).Where("id = ?", id).Updates(paymentMethod).Error; err != nil {
		return nil, err
	}

	return paymentMethod, nil
}

func (r *PaymentMethodRepository) Delete(id string) (err error) {
	if err := r.DB.Delete(&Models.PaymentMethod{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
