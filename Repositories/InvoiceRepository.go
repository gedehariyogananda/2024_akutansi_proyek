package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Utils"

	"gorm.io/gorm"
)

type (
	IInvoiceRepository interface {
		Store(trx *gorm.DB, invoice *Models.Invoice) (*Models.Invoice, error)
		GetAllByCompany(companyID string, query *Common.Query) (invoices []*Models.Invoice, totalData int64, err error)
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

func (r *InvoiceRepository) GetAllByCompany(companyID string, query *Common.Query) (invoices []*Models.Invoice, totalData int64, err error) {
	if err := r.DB.Model(&Models.Invoice{}).Scopes(Helper.FilterSearchRiwayatTransaction(*query.Search)).Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.
		Select("id", "customer_name", "invoice_number", "status", "sub_total", "created_at").
		Where("company_id = ?", companyID).
		Preload("InvoiceItems", func(invItemPayload *gorm.DB) *gorm.DB {
			return invItemPayload.Select("invoice_id", "sum(quantity) as total").Group("invoice_id")
		}).
		Scopes(
			Utils.Paginate(query.Page, query.Limit),
			Helper.FilterSearchRiwayatTransaction(*query.Search)).
		Order("created_at desc").
		Find(&invoices).Error; err != nil {
		return nil, 0, err
	}

	return invoices, totalData, nil
}
