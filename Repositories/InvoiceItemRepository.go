package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto/Response"
	"fmt"

	"gorm.io/gorm"
)

type (
	IInvoiceItemRepository interface {
		Store(trx *gorm.DB, invoiceItem *Models.InvoiceItem) error
		DeleteByInvoiceID(trx *gorm.DB, invoiceID string) error
		GetMostProductSold(companyID string, date string) (invoiceItems Response.InvoiceItemResponse, err error)
		GetBestSellingProducts(companyID, startDate, endDate string, limit int) (invoiceItems []Response.InvoiceItemResponse, err error)
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

func (r *InvoiceItemRepository) GetMostProductSold(companyID string, date string) (invoiceItems Response.InvoiceItemResponse, err error) {
	if err := r.DB.
		Model(&Models.InvoiceItem{}).
		Select("invoice_items.sellable_product_id, SUM(invoice_items.quantity) as count_sale, sellable_products.name as product_name").
		Joins("JOIN invoices ON invoices.id = invoice_items.invoice_id").
		Joins("JOIN sellable_products ON sellable_products.id = invoice_items.sellable_product_id").
		Where("invoices.company_id = ? AND DATE(invoices.created_at) = ?", companyID, date).
		Group("invoice_items.sellable_product_id, sellable_products.name").
		Order("count_sale DESC").
		Limit(1).
		Scan(&invoiceItems).Error; err != nil {
		return invoiceItems, fmt.Errorf("error saat mencari data penjualan: %w", err)
	}

	return invoiceItems, nil
}

func (r *InvoiceItemRepository) GetBestSellingProducts(companyID, startDate, endDate string, limit int) (invoiceItems []Response.InvoiceItemResponse, err error) {

	fmt.Println("repository", startDate)

	if err := r.DB.
		Scopes(Helper.FilterDateInvoiceDashboard(startDate, endDate)).
		Model(&Models.InvoiceItem{}).
		Select("invoice_items.sellable_product_id, SUM(invoice_items.quantity) as count_sale, "+
			"SUM(invoice_items.quantity * invoice_items.price) as total_revenue, "+
			"sellable_products.name as product_name").
		Joins("JOIN invoices ON invoices.id = invoice_items.invoice_id").
		Joins("JOIN sellable_products ON sellable_products.id = invoice_items.sellable_product_id").
		Group("invoice_items.sellable_product_id, sellable_products.name").
		Where("invoices.company_id = ?", companyID).
		Order("count_sale DESC").
		Limit(limit).
		Scan(&invoiceItems).Error; err != nil {
		return invoiceItems, fmt.Errorf("error saat mencari data penjualan: %w", err)
	}

	return invoiceItems, nil
}

func (r *InvoiceItemRepository) DeleteByInvoiceID(trx *gorm.DB, invoiceID string) error {
	db := trx
	if db == nil {
		db = r.DB
	}

	if err := db.Where("invoice_id = ?", invoiceID).Delete(&Models.InvoiceItem{}).Error; err != nil {
		return fmt.Errorf("error saat menghapus invoice item: %w", err)
	}

	return nil
}