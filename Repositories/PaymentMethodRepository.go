package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IPaymentMethodRepository interface {
		FindAll(company_id int) (paymentMethod *[]Models.PaymentMethod, err error)
	}

	PaymentMethodRepository struct {
		DB *gorm.DB
	}
)

func PaymentMethodRepositoryProvider(db *gorm.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{DB: db}
}

func (repository *PaymentMethodRepository) FindAll(company_id int) (paymentMethod *[]Models.PaymentMethod, err error) {
	paymentMethod = &[]Models.PaymentMethod{}
	err = repository.DB.Where("company_id = ?", company_id).Find(paymentMethod).Error
	if err != nil {
		return nil, err
	}
	return paymentMethod, nil
}
