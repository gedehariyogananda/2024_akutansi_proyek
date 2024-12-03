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
		FindByID(id string, companyID string) (invoice *Models.Invoice, err error)
		Update(id string, invoice *Models.Invoice) (err error)
		GetAllByCompany(companyID string, query *Common.Query) (invoices []*Models.Invoice, totalData int64, err error)
		GetByInvoiceID(companyID string, invoiceID string) (invoice *Models.Invoice, err error)
		SumSalesByDate(companyID string, date string) (totalSales float64, err error)
		SumSalesByYearMonth(companyID string, year int, month int) (totalSales float64, err error)
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

func (r *InvoiceRepository) Update(id string, invoice *Models.Invoice) (err error) {
	if err := r.DB.
		Model(&Models.Invoice{}).
		Where("id = ?", id).
		Updates(invoice).Error; err != nil {
		return err
	}

	return nil
}

func (r *InvoiceRepository) GetAllByCompany(companyID string, query *Common.Query) (invoices []*Models.Invoice, totalData int64, err error) {
	if err := r.DB.Model(&Models.Invoice{}).Scopes(Helper.FilterSearchRiwayatTransaction(*query.Search)).Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.
		Select("id", "customer_name", "invoice_number", "status", "sub_total", "created_at").
		Where("company_id = ?", companyID).
		Preload("InvoiceItems", func(invItemPayload *gorm.DB) *gorm.DB {
			return invItemPayload.Select("invoice_id", "quantity")
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

func (r *InvoiceRepository) GetByInvoiceID(companyID string, invoiceID string) (invoice *Models.Invoice, err error) {
	if err := r.DB.
		Where("company_id = ?", companyID).
		Preload("InvoiceItems", func(invItemPayload *gorm.DB) *gorm.DB {
			return invItemPayload.Where("invoice_id = ?", invoiceID).
				Select("invoice_id", "sellable_product_id", "quantity").
				Preload("SellableProduct", func(spPayload *gorm.DB) *gorm.DB {
					return spPayload.Select("id", "name", "price")
				})
		}).
		First(&invoice).Error; err != nil {
		return nil, err
	}

	return invoice, nil
}

func (r *InvoiceRepository) FindByID(id string, companyID string) (invoice *Models.Invoice, err error) {
	if err := r.DB.
		Where("company_id = ?", companyID).
		Where("id = ?", id).
		First(&invoice).Error; err != nil {
		return nil, err
	}

	return invoice, nil
}

func (r *InvoiceRepository) SumSalesByDate(companyID string, date string) (totalSales float64, err error) {
	if err := r.DB.
		Model(&Models.Invoice{}).
		Where("company_id = ?", companyID).
		Where("date(created_at) = ?", date).
		Select("sum(sub_total)").
		Scan(&totalSales).Error; err != nil {
		return 0, err
	}

	return totalSales, nil
}

func (r *InvoiceRepository) SumSalesByYearMonth(companyID string, year int, month int) (totalSales float64, err error) {

	if err := r.DB.
		Model(&Models.Invoice{}).
		Where("company_id = ?", companyID).
		Where("EXTRACT(YEAR FROM created_at) = ?", year).
		Where("EXTRACT(MONTH FROM created_at) = ?", month).
		Select("sum(sub_total)").
		Scan(&totalSales).Error; err != nil {
		return 0, err
	}

	return totalSales, nil
}
