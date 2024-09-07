package Services

import (
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"fmt"
	"net/http"
	"time"
)

type (
	IInvoiceService interface {
		CreateInvoicePurchased(request *Dto.InvoiceRequestClient, company_id int) (err error, statusCode int)
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

func (s *InvoiceService) CreateInvoicePurchased(request *Dto.InvoiceRequestClient, company_id int) (err error, statusCode int) {
	dateInvoice := time.Now().Format("2006/01/02")
	invoiceNumber := fmt.Sprintf("%s-%s", dateInvoice, request.InvoiceCustomer)

	// Calculate total amount
	totalAmount := 0
	for _, purchase := range request.Purchaseds {
		totalAmount += purchase.TotalPrice
	}

	invoiceRequestDTO := &Dto.InvoiceRequestDTO{
		InvoiceNumber:   invoiceNumber,
		InvoiceCustomer: request.InvoiceCustomer,
		InvoiceDate:     time.Now().Format("2006-01-02"),
		TotalAmount:     totalAmount,
		CompanyID:       company_id,
		PaymentMethodId: request.PaymentMethodID,
	}

	invoice, err := s.InvoiceRepository.Create(invoiceRequestDTO, company_id)
	if err != nil {
		return fmt.Errorf("failed to create invoice: %w", err), http.StatusInternalServerError
	}

	if invoice == nil || invoice.ID == 0 {
		return fmt.Errorf("failed to create invoice: invoice ID is 0"), http.StatusInternalServerError
	}

	invoiceID := invoice.ID

	// Handle saleable products
	for _, purchase := range request.Purchaseds {
		if purchase.IsSaleableProduct {
			invoiceSaleableRequestDTO := &Dto.InvoiceSaleableRequestDTO{
				InvoiceID:         invoiceID,
				SaleableProductID: purchase.ID,
				QuantitySold:      purchase.QuantitySold,
				CompanyID:         company_id,
			}

			if err := s.InvoiceSaleableRepository.Create(invoiceSaleableRequestDTO); err != nil {
				return fmt.Errorf("failed to create invoice saleable product: %w", err), http.StatusInternalServerError
			}
		}
	}

	// Handle material products
	for _, purchase := range request.Purchaseds {
		if !purchase.IsSaleableProduct {
			invoiceMaterialRequestDTO := &Dto.InvoiceMaterialRequestDTO{
				InvoiceID:         invoiceID,
				MaterialProductID: purchase.ID,
				QuantitySold:      purchase.QuantitySold,
				CompanyID:         company_id,
			}

			if err := s.InvoiceMaterialRepository.Create(invoiceMaterialRequestDTO); err != nil {
				return fmt.Errorf("failed to create invoice material product: %w", err), http.StatusInternalServerError
			}
		}
	}

	return nil, http.StatusCreated
}
