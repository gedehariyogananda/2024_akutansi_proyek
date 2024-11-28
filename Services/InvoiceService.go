package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"errors"
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type (
	IInvoiceService interface {
		CreateInvoicePurchased(requestClient *Dto.InvoiceRequestDTO, companyID string) (invoice *Models.Invoice, statusCode int, err error)
	}

	InvoiceService struct {
		invoiceRepository         Repositories.IInvoiceRepository
		invoiceItemRepository     Repositories.IInvoiceItemRepository
		sellableProductRepository Repositories.ISellableProductRepository
		receiptProductRepository  Repositories.IReceiptRepository
		materialProductRepository Repositories.IMaterialProductRepository
		DB                        *gorm.DB
	}
)

func InvoiceServiceProvider(invoiceRepository Repositories.IInvoiceRepository, invoiceItemRepository Repositories.IInvoiceItemRepository, sellableProductRepository Repositories.ISellableProductRepository, receiptProductRepository Repositories.IReceiptRepository, materialProductRepository Repositories.IMaterialProductRepository, DB *gorm.DB) *InvoiceService {
	return &InvoiceService{
		invoiceRepository:         invoiceRepository,
		invoiceItemRepository:     invoiceItemRepository,
		sellableProductRepository: sellableProductRepository,
		receiptProductRepository:  receiptProductRepository,
		materialProductRepository: materialProductRepository,
		DB:                        DB,
	}
}

func (invoiceService *InvoiceService) CreateInvoicePurchased(requestClient *Dto.InvoiceRequestDTO, companyID string) (invoice *Models.Invoice, statusCode int, err error) {

	// intial db transaction
	trx := invoiceService.DB.Begin()
	if trx.Error != nil {
		return nil, http.StatusInternalServerError, trx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
			err = fmt.Errorf("panic occurred: %v", r)
		} else if err != nil {
			trx.Rollback()
		} else {
			trx.Commit()
		}
	}()

	invoiceDataClient := &Models.Invoice{
		CustomerName:  requestClient.CustomerName,
		PhoneNumber:   requestClient.PhoneNumber,
		Note:          requestClient.Notes,
		TaxID:         requestClient.TaxID,
		PaymentMethod: requestClient.PaymentMethod,
		InvoiceNumber: requestClient.InvoiceNumber,
		CompanyID:     companyID,
		Status:        requestClient.Status,
		Tax:           requestClient.Tax,
		SubTotal:      requestClient.SubTotal,
	}

	invoice, err = invoiceService.invoiceRepository.StoreTrx(trx, invoiceDataClient)
	if err != nil {
		trx.Rollback()
		log.Println("ROLLBACK!")
		return nil, http.StatusInternalServerError, err
	}

	log.Println("LOG: SUCCESS CREATED INVOICE!")

	for _, purchasedItem := range requestClient.Purchaseds {
		sellableProduct, err := invoiceService.sellableProductRepository.Find(purchasedItem.ID)
		if err != nil {
			return nil, http.StatusNotFound, err
		}

		// update curren_quantity sellable product
		if sellableProduct.CurrentQuantity < purchasedItem.Qty {
			return nil, http.StatusBadRequest, errors.New("stock not enough")
		}

		err = invoiceService.sellableProductRepository.UpdateTrx(trx, purchasedItem.ID, &Models.SellableProduct{
			CurrentQuantity: sellableProduct.CurrentQuantity - purchasedItem.Qty,
		})

		if err != nil {
			log.Println("LOG: ROLLBACK!")
			return nil, http.StatusInternalServerError, err
		}

		log.Println("LOG: SUCCESS UPDATED SELLABLE PRODUCT QUANTITY!")

		invoiceItem := &Models.InvoiceItem{
			InvoiceID:         invoice.ID,
			SellableProductID: sellableProduct.ID,
			Quantity:          purchasedItem.Qty,
			CompanyID:         companyID,
			Price:             purchasedItem.PriceAll,
			PromoID:           purchasedItem.PromoID,
			PromoAmount:       purchasedItem.PromoAmount,
		}

		_, err = invoiceService.invoiceItemRepository.StoreTrx(trx, invoiceItem)
		if err != nil {
			log.Println("LOG: ROLLBACK!")
			return nil, http.StatusInternalServerError, err
		}

		log.Println("LOG: SUCCESS CREATED INVOICE ITEM!")

		if sellableProduct.HasReceipt {

			materialProductData, err := invoiceService.materialProductRepository.FindByStatus(companyID, true)
			if err != nil {
				return nil, http.StatusNotFound, err
			}

			receiptAllProduct, err := invoiceService.receiptProductRepository.FindAll(sellableProduct.ID)
			if err != nil {
				return nil, http.StatusNotFound, err
			}

			for _, receiptProduct := range receiptAllProduct {
				for _, materialProduct := range materialProductData {
					if materialProduct.ID == receiptProduct.MaterialProductID {
						if materialProduct.CurrentQuantity < purchasedItem.Qty*receiptProduct.Quantity {
							return nil, http.StatusBadRequest, errors.New("stock not enough")
						}

						err = invoiceService.materialProductRepository.UpdateTrx(trx, materialProduct.ID,
							&Models.MaterialProduct{
								CurrentQuantity: materialProduct.CurrentQuantity - (purchasedItem.Qty * receiptProduct.Quantity),
							})

						log.Println("LOG: SUCCESS UPDATED MATERIAL QUANTITY!")

						if err != nil {
							log.Println("ROLLBACK!")
							return nil, http.StatusInternalServerError, err
						}
					}
				}
			}
		}
	}

	responPrefix := &Models.Invoice{
		InvoiceNumber: invoice.InvoiceNumber,
		CustomerName:  invoice.CustomerName,
		SubTotal:      invoice.SubTotal,
	}

	return responPrefix, http.StatusOK, nil

}
