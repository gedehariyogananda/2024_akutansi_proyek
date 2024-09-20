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
		GetInvoice(invoice_id string) (invoiceSet *Models.Invoice, invoiceRes *[]Response.DetailSaleableResponseDTO, err error, statusCode int)
		DeleteInvoice(invoice_id string, company_id string) (statusCode int, err error)
		UpdateInvoiceDetail(company_id string, invoice_id string, request *Dto.InvoiceRequestClient) (invoice *Models.Invoice, err error, statusCode int)
	}

	InvoiceService struct {
		InvoiceRepository         Repositories.IInvoiceRepository
		InvoiceMaterialRepository Repositories.IInvoiceMaterialRepository
		InvoiceSaleableRepository Repositories.IInvoiceSaleableRepository
		SaleableProductRepository Repositories.ISaleableProductRepository
		PaymentMethodRepository   Repositories.IPaymentMethodRepository
		CompanyRepository         Repositories.ICompanyRepository
	}
)

func InvoiceServiceProvider(invoiceRepository Repositories.IInvoiceRepository, invoiceMaterialRepository Repositories.IInvoiceMaterialRepository, invoiceSaleableRepository Repositories.IInvoiceSaleableRepository, saleableProductRepository Repositories.ISaleableProductRepository, paymentMethodRepository Repositories.IPaymentMethodRepository, companyRepository Repositories.ICompanyRepository) *InvoiceService {
	return &InvoiceService{
		InvoiceRepository:         invoiceRepository,
		InvoiceMaterialRepository: invoiceMaterialRepository,
		InvoiceSaleableRepository: invoiceSaleableRepository,
		SaleableProductRepository: saleableProductRepository,
		PaymentMethodRepository:   paymentMethodRepository,
		CompanyRepository:         companyRepository,
	}
}

func (s *InvoiceService) CreateInvoicePurchased(request *Dto.InvoiceRequestClient, company_id string) (invoice *Models.Invoice, err error, statusCode int) {
	// Calculate total amount
	totalAmount := 0
	for _, purchase := range request.Purchaseds {
		totalAmount += purchase.TotalPrice
	}

	payment, err := s.PaymentMethodRepository.FindById(request.PaymentMethodID)

	if err != nil {
		return nil, fmt.Errorf("payment method not found"), http.StatusNotFound
	}

	company, err := s.CompanyRepository.FindCompany(company_id)

	if err != nil {
		return nil, fmt.Errorf("company not found"), http.StatusNotFound
	}

	statusInv := Models.PROCESS
	moneyReceive := 0

	if payment.MethodName != "Cash" {
		statusInv = Models.PROCESS
		moneyReceive = totalAmount
	}

	invoiceRequestDTO := &Dto.InvoiceRequestDTO{
		InvoiceCustomer: request.InvoiceCustomer,
		InvoiceDate:     time.Now().Format("2006-01-02 15:04:05"),
		TotalAmount:     totalAmount,
		StatusInvoice:   string(statusInv),
		CompanyID:       company_id,
		PaymentMethodId: request.PaymentMethodID,
		MoneyReceived:   moneyReceive,
	}

	invoice, err = s.InvoiceRepository.Create(invoiceRequestDTO, company.CodeCompany, company_id)
	if err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err), http.StatusBadRequest
	}

	// checked exist if saleable product
	for _, purchase := range request.Purchaseds {
		isExist, err := s.SaleableProductRepository.CheckProductExist(company_id, purchase.ID)

		if err != nil {
			return nil, fmt.Errorf("failed to check product exist: %w", err), http.StatusBadRequest
		}

		if !isExist {
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

		if isExist {
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

	// not permision to update status, hanya bisa button aja

	invoice.InvoiceCustomer = request.InvoiceCustomer
	invoice.MoneyReceived = float64(request.MoneyReceived)
	invoice.PaymentMethodID = request.PaymentMethodId

	paymentMethod, err := s.PaymentMethodRepository.FindById(request.PaymentMethodId)

	if err != nil {
		return nil, fmt.Errorf("payment method not found"), http.StatusNotFound
	}

	if paymentMethod.MethodName != "Cash" {
		invoice.MoneyReceived = invoice.TotalAmount
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
			QuantitySold: item.QuantitySold,
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
			QuantitySold: item.QuantitySold,
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

func (s *InvoiceService) DeleteInvoice(invoice_id string, company_id string) (statusCode int, err error) {
	invoice, err := s.InvoiceRepository.FindById(invoice_id)

	if err != nil {
		return http.StatusNotFound, fmt.Errorf("invoice not found")
	}

	if invoice.CompanyID != company_id {
		return http.StatusForbidden, fmt.Errorf("access forbidden: company_id mismatch")
	}

	if err := s.InvoiceRepository.Delete(invoice_id); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to delete invoice: %w", err)
	}

	return http.StatusOK, nil
}

func (s *InvoiceService) UpdateInvoiceDetail(company_id string, invoice_id string, request *Dto.InvoiceRequestClient) (invoice *Models.Invoice, err error, statusCode int) {
	invoice, err = s.InvoiceRepository.FindById(invoice_id)

	if err != nil {
		return nil, err, http.StatusNotFound
	}

	if invoice.CompanyID != company_id {
		return nil, fmt.Errorf("access forbidden: company_id mismatch"), http.StatusForbidden
	}

	totalAmount := 0
	for _, purchase := range request.Purchaseds {
		totalAmount += purchase.TotalPrice
	}

	payment, err := s.PaymentMethodRepository.FindById(request.PaymentMethodID)

	if err != nil {
		return nil, fmt.Errorf("payment method not found"), http.StatusNotFound
	}

	moneyInit := invoice.MoneyReceived
	statusInv := invoice.StatusInvoice

	if payment.MethodName != "Cash" {
		moneyInit = float64(totalAmount)
		statusInv = Models.PROCESS

	}

	invoiceRequestDTO := &Dto.InvoiceRequestDTO{
		InvoiceCustomer: request.InvoiceCustomer,
		TotalAmount:     totalAmount,
		CompanyID:       company_id,
		MoneyReceived:   int(moneyInit),
		PaymentMethodId: request.PaymentMethodID,
		StatusInvoice:   string(statusInv),
	}

	if err := s.InvoiceRepository.UpdateByInvoiceId(invoice_id, company_id, invoiceRequestDTO); err != nil {
		return nil, fmt.Errorf("failed to update invoice: %w", err), http.StatusBadRequest
	}

	for _, purchase := range request.Purchaseds {
		isExist, err := s.SaleableProductRepository.CheckProductExist(company_id, purchase.ID)

		if err != nil {
			return nil, fmt.Errorf("failed to check product exist: %w", err), http.StatusBadRequest
		}

		if !isExist {
			invoiceMaterialRequestDTO := &Dto.InvoiceMaterialRequestDTO{
				InvoiceID:         invoice_id,
				MaterialProductID: purchase.ID,
				QuantitySold:      purchase.QuantitySold,
				CompanyID:         company_id,
			}

			if err := s.InvoiceMaterialRepository.Update(invoiceMaterialRequestDTO, invoice_id); err != nil {
				return nil, fmt.Errorf("failed to create material product for invoice: %w", err), http.StatusBadRequest
			}
		}

		if isExist {
			invoiceSaleableRequestDTO := &Dto.InvoiceSaleableRequestDTO{
				InvoiceID:         invoice_id,
				SaleableProductID: purchase.ID,
				QuantitySold:      purchase.QuantitySold,
				CompanyID:         company_id,
			}

			if err := s.InvoiceSaleableRepository.Update(invoiceSaleableRequestDTO, invoice_id); err != nil {
				return nil, fmt.Errorf("failed to create saleable product for invoice: %w", err), http.StatusBadRequest
			}
		}
	}

	invoice, err = s.InvoiceRepository.FindSelectRelasi(invoice.ID)

	if err != nil {
		return nil, err, http.StatusNotFound
	}

	return invoice, nil, http.StatusOK
}
