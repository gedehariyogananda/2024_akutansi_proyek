package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type (
	IInvoiceRepository interface {
		Create(request *Dto.InvoiceRequestDTO, codeCompany string, company_id string) (invoice *Models.Invoice, err error)
		GetAll(company_id string, date string) (invoices *[]Models.Invoice, err error)
		FindById(invoice_id string) (invoice *Models.Invoice, err error)
		Update(invoice *Models.Invoice) (err error)
		FindSelectRelasi(invoice_id string) (invoice *Models.Invoice, err error)
		Delete(invoice_id string) (err error)
		UpdateByInvoiceId(invoice_id string, company_id string, request *Dto.InvoiceRequestDTO) (err error)
	}

	InvoiceRepository struct {
		DB *gorm.DB
	}
)

func InvoiceRepositoryProvider(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{DB: db}
}

func (r *InvoiceRepository) Create(request *Dto.InvoiceRequestDTO, codeCompany string, company_id string) (invoice *Models.Invoice, err error) {

	// generate invoice number nan
	invoiceNumber, err := Utils.GenerateInvoiceNumber(r.DB, codeCompany, company_id)

	if err != nil {
		return nil, fmt.Errorf("failed to generate invoice number: %w", err)
	}

	invoice = &Models.Invoice{
		InvoiceNumber:   invoiceNumber,
		InvoiceCustomer: request.InvoiceCustomer,
		CompanyID:       request.CompanyID,
		PaymentMethodID: request.PaymentMethodId,
		InvoiceDate:     request.InvoiceDate,
		TotalAmount:     float64(request.TotalAmount),
		MoneyReceived:   float64(request.MoneyReceived),
		StatusInvoice:   Models.StatusInvoice(request.StatusInvoice),
		CreatedAt:       time.Now(),
	}

	if err := r.DB.Create(invoice).First(invoice).Error; err != nil {
		return nil, err
	}

	return invoice, nil
}

func (r *InvoiceRepository) GetAll(company_id string, date string) (invoices *[]Models.Invoice, err error) {
	invoices = &[]Models.Invoice{}

	if err := r.DB.Where("company_id = ? AND DATE(invoice_date) = ?", company_id, date).
		Preload("PaymentMethod").
		Preload("Company").
		Order("created_at DESC").
		Find(invoices).Error; err != nil {
		return nil, err
	}

	return invoices, nil
}

func (r *InvoiceRepository) FindById(invoice_id string) (invoice *Models.Invoice, err error) {
	invoice = &Models.Invoice{}

	if err := r.DB.First(invoice, "id = ?", invoice_id).Error; err != nil {
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

func (r *InvoiceRepository) FindSelectRelasi(invoice_id string) (invoice *Models.Invoice, err error) {
	invoice = &Models.Invoice{}

	if err := r.DB.Preload("PaymentMethod").First(invoice, "id = ?", invoice_id).Error; err != nil {
		return nil, fmt.Errorf("invoice not found")
	}

	return invoice, nil
}

func (r *InvoiceRepository) Delete(invoice_id string) (err error) {
	if err := r.DB.Delete(&Models.Invoice{}, "id = ?", invoice_id).Error; err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}

func (r *InvoiceRepository) UpdateByInvoiceId(invoice_id string, company_id string, request *Dto.InvoiceRequestDTO) (err error) {
	var invoice Models.Invoice
	if err := r.DB.First(&invoice, "id = ? AND company_id = ?", invoice_id, company_id).Error; err != nil {
		return fmt.Errorf("invoice not found")
	}

	invoice.InvoiceCustomer = request.InvoiceCustomer
	invoice.CompanyID = request.CompanyID
	invoice.PaymentMethodID = request.PaymentMethodId
	invoice.TotalAmount = float64(request.TotalAmount)
	invoice.MoneyReceived = float64(request.MoneyReceived)
	invoice.StatusInvoice = Models.StatusInvoice(request.StatusInvoice)
	invoice.UpdatedAt = time.Now()

	if err := r.DB.Save(&invoice).Error; err != nil {
		return fmt.Errorf("failed to update invoice: %w", err)
	}

	return nil
}
