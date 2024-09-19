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
		CreateInvoicePurchased(request *Dto.InvoiceRequestClient, company_id string) (invoice *Models.Invoice, err error, statusCode int)
		UpdateStatusInvoice(request *Dto.InvoiceUpdateRequestDTO, invoice_id string, company_id string) (invoice *Models.Invoice, err error, statusCode int)
		UpdateMoneyReveived(request *Dto.InvoiceMoneyReceivedRequestDTO, invoice_id string, company_id string) (invoice *Models.Invoice, MoneyBack float64, err error, statusCode int)
		GetAllInvoices(company_id string, filterDate string) (invoices *[]Models.Invoice, err error, statusCode int)
		UpdateInvoiceCustomer(company_id string, invoice_id string, request *Dto.InvoiceUpdateRequestDTO) (invoice *Models.Invoice, err error, statusCode int)
		GetInvoiceDetail(invoice_id string) (invoiceDetail *[]Response.InvoiceDetailResponse, err error, statusCode int)
		GetInvoice(invoice_id string) (invoiceSet *Models.Invoice, invoiceRes *[]Response.DetailSaleableResponseDTO, err error, statusCode int)
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

func (s *InvoiceService) CreateInvoicePurchased(request *Dto.InvoiceRequestClient, company_id string) (invoice *Models.Invoice, err error, statusCode int) {
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
		InvoiceDate:     time.Now().Format("2006-01-02 15:04:05"),
		TotalAmount:     totalAmount,
		CompanyID:       company_id,
		PaymentMethodId: request.PaymentMethodID,
	}

	invoice, err = s.InvoiceRepository.Create(invoiceRequestDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err), http.StatusBadRequest
	}

	// Handle saleable products
	for _, purchase := range request.Purchaseds {
		if purchase.IsSaleableProduct {
			invoiceSaleableRequestDTO := &Dto.InvoiceSaleableRequestDTO{
				InvoiceID:         invoice.ID,
				SaleableProductID: purchase.ID,
				QuantitySold:      purchase.QuantitySold,
				CompanyID:         company_id,
			}

			if err := s.InvoiceSaleableRepository.Create(invoiceSaleableRequestDTO); err != nil {
				return nil, fmt.Errorf("failed to create saleable product for invoice: %w", err), http.StatusBadRequest
			}
		}
	}

	// Handle material products
	for _, purchase := range request.Purchaseds {
		if !purchase.IsSaleableProduct {
			invoiceMaterialRequestDTO := &Dto.InvoiceMaterialRequestDTO{
				InvoiceID:         invoice.ID,
				MaterialProductID: purchase.ID,
				QuantitySold:      purchase.QuantitySold,
				CompanyID:         company_id,
			}

			if err := s.InvoiceMaterialRepository.Create(invoiceMaterialRequestDTO); err != nil {
				return nil, fmt.Errorf("failed to create material product for invoice: %w", err), http.StatusBadRequest
			}
		}
	}

	invoice, err = s.InvoiceRepository.FindSelectRelasi(invoice.ID)

	if err != nil {
		return nil, err, http.StatusNotFound
	}

	return invoice, nil, http.StatusOK
}

func (s *InvoiceService) UpdateStatusInvoice(request *Dto.InvoiceUpdateRequestDTO, invoice_id string, company_id string) (invoice *Models.Invoice, err error, statusCode int) {
	invoice, err = s.InvoiceRepository.FindById(invoice_id)
	if err != nil {
		return nil, fmt.Errorf("invoice Not Found : %w", err), http.StatusNotFound
	}

	if invoice.CompanyID != company_id {
		return nil, fmt.Errorf("access forbidden: company_id mismatch"), http.StatusForbidden
	}

	switch request.StatusInvoice {
	case "DONE":
		invoice.StatusInvoice = Models.DONE
	case "CANCEL":
		invoice.StatusInvoice = Models.CANCEL
	case "PROCESS":
		invoice.StatusInvoice = Models.PROCESS
	default:
		return nil, fmt.Errorf("invalid status invoice"), http.StatusBadRequest
	}

	if err := s.InvoiceRepository.Update(invoice); err != nil {
		return nil, fmt.Errorf("failed to update invoice: %w", err), http.StatusInternalServerError
	}

	invoice, err = s.InvoiceRepository.FindSelectRelasi(invoice_id)

	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return invoice, nil, http.StatusOK
}

func (s *InvoiceService) UpdateMoneyReveived(request *Dto.InvoiceMoneyReceivedRequestDTO, invoice_id string, company_id string) (invoice *Models.Invoice, MoneyBack float64, err error, statusCode int) {
	invoice, err = s.InvoiceRepository.FindById(invoice_id)

	if err != nil {
		return nil, 0, fmt.Errorf("invoice not found, %s", err), http.StatusNotFound
	}

	if invoice.StatusInvoice == "DONE" || invoice.StatusInvoice == "CANCEL" {
		return nil, 0, fmt.Errorf("invoice status is already %s", invoice.StatusInvoice), http.StatusBadRequest
	}

	if invoice.CompanyID != company_id {
		return nil, 0, fmt.Errorf("access forbidden: company_id mismatch"), http.StatusForbidden
	}

	invoice.MoneyReceived = request.MoneyReceived
	invoice.StatusInvoice = Models.PROCESS

	if err := s.InvoiceRepository.Update(invoice); err != nil {
		return nil, 0, fmt.Errorf("failed to update invoice money received: %w", err), http.StatusBadRequest
	}

	// Calculate money back
	MoneyBack = request.MoneyReceived - invoice.TotalAmount

	return invoice, MoneyBack, nil, http.StatusOK
}

func (s *InvoiceService) GetAllInvoices(company_id string, filterDate string) (invoices *[]Models.Invoice, err error, statusCode int) {
	date := filterDate

	if filterDate == "" {
		invoices, err = s.InvoiceRepository.GetAll(company_id, date)

		if err != nil {
			return nil, err, http.StatusNotFound
		}

		return invoices, nil, http.StatusOK

	}

	invoices, err = s.InvoiceRepository.GetAll(company_id, date)

	if err != nil {
		return nil, err, http.StatusNotFound
	}

	return invoices, nil, http.StatusOK

}

func (s *InvoiceService) UpdateInvoiceCustomer(company_id string, invoice_id string, request *Dto.InvoiceUpdateRequestDTO) (invoice *Models.Invoice, err error, statusCode int) {
	invoice, err = s.InvoiceRepository.FindById(invoice_id)

	if err != nil {
		return nil, err, http.StatusNotFound
	}

	if invoice.CompanyID != company_id {
		return nil, fmt.Errorf("access forbidden: company_id mismatch"), http.StatusForbidden
	}

	switch request.StatusInvoice {
	case "DONE":
		invoice.StatusInvoice = Models.DONE
	case "CANCEL":
		invoice.StatusInvoice = Models.CANCEL
	case "PROCESS":
		invoice.StatusInvoice = Models.PROCESS
	default:
		invoice.StatusInvoice = Models.WAITING
	}

	if request.InvoiceCustomer != "" {
		dateInvoice := invoice.CreatedAt.Format("2006/01/02")
		invoice.InvoiceCustomer = request.InvoiceCustomer
		invoice.InvoiceNumber = fmt.Sprintf("%s-%s", dateInvoice, request.InvoiceCustomer)
	}

	if err := s.InvoiceRepository.Update(invoice); err != nil {
		return nil, fmt.Errorf("failed to update invoice: %w", err), http.StatusInternalServerError
	}

	invoice, err = s.InvoiceRepository.FindSelectRelasi(invoice_id)

	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return invoice, nil, http.StatusOK
}

func (s *InvoiceService) GetInvoiceDetail(invoice_id string) (invoiceDetail *[]Response.InvoiceDetailResponse, err error, statusCode int) {
	return
}

func (s *InvoiceService) GetInvoice(invoice_id string) (invoiceSet *Models.Invoice, invoiceRes *[]Response.DetailSaleableResponseDTO, err error, statusCode int) {
	invoices, err := s.InvoiceSaleableRepository.FindByInvoiceId(invoice_id)

	if err != nil {
		return nil, nil, err, http.StatusNotFound

	}

	invoiceMaterial, err := s.InvoiceMaterialRepository.FindByInvoiceId(invoice_id)

	if err != nil {
		return nil, nil, err, http.StatusNotFound
	}

	invoiceRes = &[]Response.DetailSaleableResponseDTO{}

	for _, item := range *invoices {
		invoiceDetail := Response.DetailSaleableResponseDTO{
			ID:           item.SaleableProduct.ID,
			ProductName:  item.SaleableProduct.ProductName,
			Qty:          item.QuantitySold,
			UnitPrice:    item.SaleableProduct.UnitPrice,
			CategoryName: item.SaleableProduct.Category.CategoryName,
			TotalPrice:   item.SaleableProduct.UnitPrice * float64(item.QuantitySold),
		}

		*invoiceRes = append(*invoiceRes, invoiceDetail)
	}

	for _, item := range *invoiceMaterial {
		invoiceDetail := Response.DetailSaleableResponseDTO{
			ID:           item.MaterialProduct.ID,
			ProductName:  item.MaterialProduct.MaterialProductName,
			Qty:          item.QuantitySold,
			UnitPrice:    item.MaterialProduct.UnitPriceForSelling,
			CategoryName: "",
			TotalPrice:   item.MaterialProduct.UnitPriceForSelling * float64(item.QuantitySold),
		}

		*invoiceRes = append(*invoiceRes, invoiceDetail)
	}

	invoiceSet, err = s.InvoiceRepository.FindSelectRelasi(invoice_id)

	if err != nil {
		return nil, nil, err, http.StatusNotFound
	}

	return invoiceSet, invoiceRes, nil, http.StatusOK
}
