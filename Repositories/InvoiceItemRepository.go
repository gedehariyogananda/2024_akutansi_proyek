package Repositories

import (
	"2024_akutansi_project/Models"
	"fmt"

	"gorm.io/gorm"
)

type (
	IInvoiceItemRepository interface {
		Store(trx *gorm.DB, invoiceItem *Models.InvoiceItem) error
	}

	InvoiceItemRepository struct {
		DB *gorm.DB
	}
)

func InvoiceItemRepositoryProvider(db *gorm.DB) *InvoiceItemRepository {
	return &InvoiceItemRepository{DB: db}
}

func (r *InvoiceItemRepository) Store(trx *gorm.DB, invoiceItem *Models.InvoiceItem) error {

	db := trx
	if db == nil {
		db = r.DB
	}

	if err := db.Create(invoiceItem).Error; err != nil {
		return fmt.Errorf("error when storing invoice item: %w", err)
	}

	return nil
}
