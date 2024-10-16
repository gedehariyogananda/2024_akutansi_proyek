package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"

	"gorm.io/gorm"
)

type (
	IInvoiceSaleableTopingRepository interface {
		Create(request *Dto.TopingsItem, companyID string, invSaleableProductID string) (err error)
	}

	InvoiceSaleableTopingRepository struct {
		DB *gorm.DB
	}
)

func InvoiceSaleableTopingRepositoryProvider(db *gorm.DB) *InvoiceSaleableTopingRepository {
	return &InvoiceSaleableTopingRepository{DB: db}
}

func (repository *InvoiceSaleableTopingRepository) Create(request *Dto.TopingsItem, companyID string, invSaleableProductID string) (err error) {
	productToping := &Models.InvoiceSaleableToping{
		TopingID:                 request.TopingID,
		InvoiceSaleableProductID: invSaleableProductID,
		CompanyID:                companyID,
	}

	if err := repository.DB.Create(productToping).Error; err != nil {
		return err
	}

	return nil
}
