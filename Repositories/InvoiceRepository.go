package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type (
	IInvoiceRepository interface {
		Create(request *Dto.InvoiceRequestDTO, company_id int) (invoice *Models.Invoice, err error)
		UpdateStatus(request *Dto.InvoiceStatusRequestDTO, invoice_id int) (invoice *Models.Invoice, err error)
		UpdateMoneyReceived(request *Dto.InvoiceMoneyReceivedRequestDTO, invoice_id int) (invoice *Response.InvoiceResponse, err error)
		GetAll(company_id int) (invoices *[]Models.Invoice, err error)
	}

	InvoiceRepository struct {
		DB *gorm.DB
	}
)

func InvoiceRepositoryProvider(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{DB: db}
}

func (r *InvoiceRepository) Create(request *Dto.InvoiceRequestDTO, company_id int) (invoice *Models.Invoice, err error) {

	invoice = &Models.Invoice{
		InvoiceNumber:   request.InvoiceNumber,
		InvoiceCustomer: request.InvoiceCustomer,
		CompanyID:       company_id,
		PaymentMethodID: request.PaymentMethodId,
		InvoiceDate:     request.InvoiceDate,
		TotalAmount:     float64(request.TotalAmount),
		MoneyReceived:   float64(request.MoneyReceived),
		StatusInvoice:   Models.WAITING,
		CreatedAt:       time.Now(),
	}

	if err := r.DB.Create(invoice).Preload("PaymentMethod").First(invoice).Error; err != nil {
		return nil, err
	}

	return invoice, nil
}

func (r *InvoiceRepository) UpdateStatus(request *Dto.InvoiceStatusRequestDTO, invoice_id int) (invoice *Models.Invoice, err error) {
	if err := r.DB.First(&invoice, invoice_id).Error; err != nil {
		return nil, fmt.Errorf("invoice not found")
	}

	invoice.StatusInvoice = Models.StatusInvoice(request.StatusInvoice)

	if err := r.DB.Model(&invoice).
		Where("id = ?", invoice_id).
		Updates(map[string]interface{}{"status_invoice": invoice.StatusInvoice}).Error; err != nil {
		return nil, fmt.Errorf("failed to update invoice status")
	}

	return invoice, nil
}
func (r *InvoiceRepository) UpdateMoneyReceived(request *Dto.InvoiceMoneyReceivedRequestDTO, invoice_id int) (invoiceRes *Response.InvoiceResponse, err error) {
	var invoice Models.Invoice

	if err := r.DB.First(&invoice, invoice_id).Error; err != nil {
		return nil, fmt.Errorf("invoice not found")
	}

	invoice.MoneyReceived = request.MoneyReceived

	if err := r.DB.Model(&invoice).
		Where("id = ?", invoice_id).
		Updates(map[string]interface{}{
			"money_received": invoice.MoneyReceived,
		}).Error; err != nil {
		return nil, fmt.Errorf("failed to update invoice money received")
	}

	moneyBack := invoice.MoneyReceived - invoice.TotalAmount

	invoiceRes = &Response.InvoiceResponse{
		ID:            invoice.ID,
		CompanyID:     invoice.CompanyID,
		InvoiceNumber: invoice.InvoiceNumber,
		TotalAmount:   invoice.TotalAmount,
		MoneyReceived: invoice.MoneyReceived,
		StatusInvoice: invoice.StatusInvoice,
		MoneyBack:     moneyBack,
	}

	return invoiceRes, nil
}

func (r *InvoiceRepository) GetAll(company_id int) (invoices *[]Models.Invoice, err error) {
	invoices = &[]Models.Invoice{}

	if err := r.DB.Where("company_id = ?", company_id).
		Preload("PaymentMethod").
		Order("created_at DESC").
		Find(invoices).Error; err != nil {
		return nil, err
	}

	return invoices, nil
}
