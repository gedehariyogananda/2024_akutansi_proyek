package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"time"

	"gorm.io/gorm"
)

type (
	IInvoiceRepository interface {
		Create(request *Dto.InvoiceRequestDTO) (invoice *Models.Invoice, err error)
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
		MoneyReceived:   0,
		StatusInvoice:   Models.WAITING,
		CreatedAt:       time.Now(),
	}

	if err := r.DB.Create(invoice).Error; err != nil {
		return nil, err
	}

	return invoice, nil
}
