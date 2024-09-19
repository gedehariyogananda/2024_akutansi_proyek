package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"

	"gorm.io/gorm"
)

type (
	IInvoiceSaleableRepository interface {
		Create(request *Dto.InvoiceSaleableRequestDTO) (err error)
		FindByInvoiceId(invoice_id string) (invoiceSaleable *[]Models.InvoiceSaleableProduct, err error)
	}

	InvoiceSaleableRepository struct {
		DB *gorm.DB
	}
)

func InvoiceSaleableRepositoryProvider(db *gorm.DB) *InvoiceSaleableRepository {
	return &InvoiceSaleableRepository{DB: db}
}

func (r *InvoiceSaleableRepository) Create(request *Dto.InvoiceSaleableRequestDTO) (err error) {
	invoiceSaleable := &Models.InvoiceSaleableProduct{
		InvoiceID:         request.InvoiceID,
		SaleableProductID: request.SaleableProductID,
		QuantitySold:      request.QuantitySold,
		CompanyID:         request.CompanyID,
	}

	if err := r.DB.Create(invoiceSaleable).Error; err != nil {
		return err
	}

	return nil
}

func (r *InvoiceSaleableRepository) FindByInvoiceId(invoice_id string) (invoiceSaleable *[]Models.InvoiceSaleableProduct, err error) {

	invoiceSaleable = &[]Models.InvoiceSaleableProduct{}

	if err := r.DB.Where("invoice_id = ?", invoice_id).
		Preload("SaleableProduct").
		Preload("SaleableProduct.Category").
		Find(&invoiceSaleable).Error; err != nil {
		return nil, err
	}

	return invoiceSaleable, nil
}
