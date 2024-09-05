package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"

	"gorm.io/gorm"
)

type (
	IInvoiceMaterialRepository interface {
		Create(request *Dto.InvoiceMaterialRequestDTO) (err error)
	}

	InvoiceMaterialRepository struct {
		DB *gorm.DB
	}
)

func InvoiceMaterialRepositoryProvider(db *gorm.DB) *InvoiceMaterialRepository {
	return &InvoiceMaterialRepository{DB: db}
}

func (r *InvoiceMaterialRepository) Create(request *Dto.InvoiceMaterialRequestDTO) (err error) {
	invoiceMaterial := &Models.InvoiceMaterialProduct{
		InvoiceID:         request.InvoiceID,
		MaterialProductID: request.MaterialProductID,
		QuantitySold:      request.QuantitySold,
		CompanyID:         request.CompanyID,
	}

	if err := r.DB.Create(invoiceMaterial).Error; err != nil {
		return err
	}

	return nil
}
