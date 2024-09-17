package Repositories

import (
	"2024_akutansi_project/Models"
	"fmt"

	"gorm.io/gorm"
)

type (
	IPaymentMethodRepository interface {
		CreateDefaultPaymentMethod(company_id int) (err error)
		FindAll(company_id int) (paymentMethod *[]Models.PaymentMethod, err error)
	}

	PaymentMethodRepository struct {
		DB *gorm.DB
	}
)

func PaymentMethodRepositoryProvider(db *gorm.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{DB: db}
}

func (r *PaymentMethodRepository) FindAll(company_id int) (paymentMethod *[]Models.PaymentMethod, err error) {
	paymentMethod = &[]Models.PaymentMethod{}

	if err := r.DB.Where("company_id = ?", company_id).Find(paymentMethod).Error; err != nil {
		return nil, err
	}

	return paymentMethod, nil
}

func (r *PaymentMethodRepository) CreateDefaultPaymentMethod(company_id int) (err error) {
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
