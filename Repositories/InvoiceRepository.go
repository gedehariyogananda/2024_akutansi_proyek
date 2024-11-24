package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IInvoiceRepository interface {
		Store(userClient *Models.Invoice) (*Models.Invoice, error)
		StoreTrx(trx *gorm.DB, invoice *Models.Invoice) (*Models.Invoice, error)
	}

	InvoiceRepository struct {
		DB *gorm.DB
	}
)

func InvoiceRepositoryProvider(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{DB: db}
}

func (r *InvoiceRepository) Store(invoice *Models.Invoice) (*Models.Invoice, error) {
	if err := r.DB.Create(invoice).Error; err != nil {
		return nil, err
	}

	return invoice, nil
}

func (r *InvoiceRepository) StoreTrx(trx *gorm.DB, invoice *Models.Invoice) (*Models.Invoice, error) {
	if err := trx.Create(invoice).Error; err != nil {
		return nil, err
	}

	return invoice, nil
}
