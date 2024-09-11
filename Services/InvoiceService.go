package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"fmt"
	"net/http"
	"time"
)

type (
	IInvoiceService interface {
		CreateInvoicePurchased(request *Dto.InvoiceRequestClient, company_id int) (invoice *Models.Invoice, err error, statusCode int)
		UpdateStatusInvoice(request *Dto.InvoiceStatusRequestDTO, invoice_id int, company_id int) (invoice *Models.Invoice, err error, statusCode int)
		UpdateMoneyReveived(request *Dto.InvoiceMoneyReceivedRequestDTO, invoice_id int, company_id int) (invoice *Response.InvoiceResponse, err error, statusCode int)
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

func (s *InvoiceService) CreateInvoicePurchased(request *Dto.InvoiceRequestClient, company_id int) (invoice *Models.Invoice, err error, statusCode int) {
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

	invoice, err = s.InvoiceRepository.Create(invoiceRequestDTO, company_id)
	if err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err), http.StatusInternalServerError
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
				return nil, fmt.Errorf("failed to create saleable product for invoice: %w", err), http.StatusInternalServerError
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
				return nil, fmt.Errorf("failed to create material product for invoice: %w", err), http.StatusInternalServerError
			}
		}
	}

	return invoice, nil, http.StatusOK
}

func (s *InvoiceService) UpdateStatusInvoice(request *Dto.InvoiceStatusRequestDTO, invoice_id int, company_id int) (invoice *Models.Invoice, err error, statusCode int) {
	invoice, err = s.InvoiceRepository.UpdateStatus(request, invoice_id)

	if err != nil {
		return nil, err, http.StatusNotFound
	}

	if invoice.CompanyID != company_id {
		return nil, fmt.Errorf("access forbidden: company_id mismatch"), http.StatusForbidden
	}

	return invoice, nil, http.StatusOK
}

func (s *InvoiceService) UpdateMoneyReveived(request *Dto.InvoiceMoneyReceivedRequestDTO, invoice_id int, company_id int) (invoice *Response.InvoiceResponse, err error, statusCode int) {
	invoice, err = s.InvoiceRepository.UpdateMoneyReceived(request, invoice_id)
	if err != nil {
		return nil, err, http.StatusNotFound
	}

	if invoice.StatusInvoice == "CANCEL" {
		return nil, fmt.Errorf("invoice is canceled"), http.StatusForbidden
	}

	if invoice.CompanyID != company_id {
		return nil, fmt.Errorf("access forbidden: company_id mismatch"), http.StatusForbidden
	}

	_, err = s.InvoiceRepository.UpdateStatus(&Dto.InvoiceStatusRequestDTO{StatusInvoice: "PROCESS"}, invoice_id)
	if err != nil {
		return nil, err, http.StatusNotFound
	}

	return invoice, nil, http.StatusOK
}
