package Services

import (
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"fmt"
	"time"
)

type (
	IInvoiceService interface {
		CreateInvoicePurchased(request *Dto.InvoiceRequestClient) (err error)
	}

	InvoiceService struct {
		InvoiceRepository         Repositories.IInvoiceRepository
		InvoiceMaterialRepository Repositories.IInvoiceMaterialRepository
		InvoiceSaleableRepository Repositories.IInvoiceSaleableRepository
	}
)

func InvoiceServiceProvider(invoiceRepository Repositories.IInvoiceRepository, invoiceMaterialRepository Repositories.IInvoiceMaterialRepository, invoiceSaleableRepository Repositories.IInvoiceSaleableRepository) *InvoiceService {
	return &InvoiceService{
		InvoiceRepository:         invoiceRepository,
		InvoiceMaterialRepository: invoiceMaterialRepository,
		InvoiceSaleableRepository: invoiceSaleableRepository,
	}
}

func (s *InvoiceService) CreateInvoicePurchased(request *Dto.InvoiceRequestClient) (err error) {
	dateInvoice := time.Now().Format("2006/01/02")
	invoiceNumber := fmt.Sprintf("%s-%s", dateInvoice, request.InvoiceCustomer)

	// calculate total_amount
	totalAmount := 0
	for _, purchase := range request.Purchaseds {
		totalAmount += purchase.TotalPrice
	}

	invoiceRequestDTO := &Dto.InvoiceRequestDTO{
		InvoiceNumber:   invoiceNumber,
		InvoiceCustomer: request.InvoiceCustomer,
		InvoiceDate:     time.Now().Format("2006-01-02"),
		TotalAmount:     totalAmount,
		CompanyID:       request.CompanyID,
		PaymentMethodId: request.PaymentMethodID,
	}

	invoice, err := s.InvoiceRepository.Create(invoiceRequestDTO)
	if err != nil {
		return err
	}

	invoiceID := invoice.ID
	if invoice.ID == 0 {
		return fmt.Errorf("failed to create invoice: ID not set")
	}

	for _, purchase := range request.Purchaseds {
		if purchase.IsSaleableProduct == true {
			invoiceSaleableRequestDTO := &Dto.InvoiceSaleableRequestDTO{
				InvoiceID:         invoiceID,
				SaleableProductID: purchase.ID,
				QuantitySold:      purchase.QuantitySold,
				CompanyID:         request.CompanyID,
			}

			if err := s.InvoiceSaleableRepository.Create(invoiceSaleableRequestDTO); err != nil {
				return err
			}
		}

		if purchase.IsSaleableProduct == false {
			invoiceMaterialRequestDTO := &Dto.InvoiceMaterialRequestDTO{
				InvoiceID:         invoiceID,
				MaterialProductID: purchase.ID,
				QuantitySold:      purchase.QuantitySold,
				CompanyID:         request.CompanyID,
			}

			if err := s.InvoiceMaterialRepository.Create(invoiceMaterialRequestDTO); err != nil {
				return err
			}
		}
	}

	return nil
}
