package Repositories

import (
	"2024_akutansi_project/Models"
	"fmt"

	"gorm.io/gorm"
)

type (
	IInvoiceItemRepository interface {
		Store(invoiceItem *Models.InvoiceItem) (*Models.InvoiceItem, error)
		StoreTrx(trx *gorm.DB, invoiceItem *Models.InvoiceItem) error
	}

	InvoiceItemRepository struct {
		DB *gorm.DB
	}
)

func InvoiceItemRepositoryProvider(db *gorm.DB) *InvoiceItemRepository {
	return &InvoiceItemRepository{DB: db}
}

func (r *InvoiceItemRepository) Store(invoiceItem *Models.InvoiceItem) (*Models.InvoiceItem, error) {
	if err := r.DB.Create(invoiceItem).Error; err != nil {
		return nil, fmt.Errorf("error when storing invoice item: %w", err)
	}

	return invoiceItem, nil
}

func (r *InvoiceItemRepository) StoreTrx(trx *gorm.DB, invoiceItem *Models.InvoiceItem) error {
	if err := trx.Create(invoiceItem).Error; err != nil {
		return fmt.Errorf("error when storing invoice item: %w", err)
	}

	return nil
}
