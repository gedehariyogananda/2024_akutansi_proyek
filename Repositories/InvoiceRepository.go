package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IInvoiceRepository interface {
		Store(trx *gorm.DB, invoice *Models.Invoice) (*Models.Invoice, error)
	}

	InvoiceRepository struct {
		DB *gorm.DB
	}
)

func InvoiceRepositoryProvider(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{DB: db}
}

func (r *InvoiceRepository) Store(trx *gorm.DB, invoice *Models.Invoice) (*Models.Invoice, error) {

	db := trx
	if db == nil {
		db = r.DB
	}

	if err := db.Create(invoice).Error; err != nil {
		return nil, err
	}

	return invoice, nil
}
