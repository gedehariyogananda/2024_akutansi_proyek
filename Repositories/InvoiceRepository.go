package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type (
	IInvoiceRepository interface {
		Create(request *Dto.InvoiceRequestDTO) (invoice *Models.Invoice, err error)
		GetAll(company_id int) (invoices *[]Models.Invoice, err error)
		FindById(invoice_id int) (invoice *Models.Invoice, err error)
		Update(invoice *Models.Invoice) (err error)
		FindSelectRelasi(invoice_id int) (invoice *Models.Invoice, err error)
	}

	InvoiceRepository struct {
		DB *gorm.DB
	}
)

func InvoiceRepositoryProvider(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{DB: db}
}

func (r *InvoiceRepository) Create(request *Dto.InvoiceRequestDTO) (invoice *Models.Invoice, err error) {

	invoice = &Models.Invoice{
		InvoiceNumber:   request.InvoiceNumber,
		InvoiceCustomer: request.InvoiceCustomer,
		CompanyID:       request.CompanyID,
		PaymentMethodID: request.PaymentMethodId,
		InvoiceDate:     request.InvoiceDate,
		TotalAmount:     float64(request.TotalAmount),
		MoneyReceived:   float64(request.MoneyReceived),
		StatusInvoice:   Models.WAITING,
		CreatedAt:       time.Now(),
	}

	if err := r.DB.Create(invoice).First(invoice).Error; err != nil {
		return nil, err
	}

	return invoice, nil
}

func (r *InvoiceRepository) GetAll(company_id int) (invoices *[]Models.Invoice, err error) {
	invoices = &[]Models.Invoice{}

	if err := r.DB.Where("company_id = ?", company_id).
		Preload("PaymentMethod").
		Preload("Company").
		Order("created_at DESC").
		Find(invoices).Error; err != nil {
		return nil, err
	}

	return invoices, nil
}

func (r *InvoiceRepository) FindById(invoice_id int) (invoice *Models.Invoice, err error) {
	invoice = &Models.Invoice{}

	if err := r.DB.First(invoice, invoice_id).Error; err != nil {
		return nil, fmt.Errorf("invoice not found")
	}

	return invoice, nil
}

func (r *InvoiceRepository) Update(invoice *Models.Invoice) (err error) {
	if err := r.DB.Save(invoice).Error; err != nil {
		return fmt.Errorf("failed to update invoice: %w", err)
	}

	return nil
}

func (r *InvoiceRepository) FindSelectRelasi(invoice_id int) (invoice *Models.Invoice, err error) {
	invoice = &Models.Invoice{}

	if err := r.DB.Preload("PaymentMethod").Preload("Company").First(invoice, invoice_id).Error; err != nil {
		return nil, fmt.Errorf("invoice not found")
	}

	return invoice, nil
}
