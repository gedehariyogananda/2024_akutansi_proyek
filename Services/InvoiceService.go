package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

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
		sellableStockRepository   Repositories.ISellableStockRepository
		materialStockRepository   Repositories.IMaterialStockRepository
		DB                        *gorm.DB
	}
)

func InvoiceServiceProvider(invoiceRepository Repositories.IInvoiceRepository, invoiceItemRepository Repositories.IInvoiceItemRepository, sellableProductRepository Repositories.ISellableProductRepository, receiptProductRepository Repositories.IReceiptRepository, materialProductRepository Repositories.IMaterialProductRepository, sellableStockRepository Repositories.ISellableStockRepository, materialStockRepository Repositories.IMaterialStockRepository, DB *gorm.DB) *InvoiceService {
	return &InvoiceService{
		invoiceRepository:         invoiceRepository,
		invoiceItemRepository:     invoiceItemRepository,
		sellableProductRepository: sellableProductRepository,
		receiptProductRepository:  receiptProductRepository,
		materialProductRepository: materialProductRepository,
		sellableStockRepository:   sellableStockRepository,
		materialStockRepository:   materialStockRepository,
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
		PhoneNumber:   &requestClient.PhoneNumber,
		Note:          requestClient.Notes,
		TaxID:         requestClient.TaxID,
		PaymentMethod: requestClient.PaymentMethod,
		InvoiceNumber: requestClient.InvoiceNumber,
		CompanyID:     companyID,
		Status:        requestClient.Status,
		Tax:           requestClient.Tax,
		SubTotal:      requestClient.SubTotal,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
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

		if sellableProduct.CompanyID != companyID {
			return nil, http.StatusForbidden, errors.New("forbidden access!")
		}

		// update current_quantity sellable product
		if sellableProduct.CurrentQuantity < purchasedItem.Qty {
			return nil, http.StatusBadRequest, errors.New("stock not enough!")
		}

		if err = invoiceService.sellableProductRepository.UpdateCurrentQty(trx, purchasedItem.ID, purchasedItem.Qty); err != nil {
			trx.Rollback()
			return nil, http.StatusInternalServerError, err
		}

		if err = invoiceService.invoiceItemRepository.StoreTrx(trx,
			&Models.InvoiceItem{
				InvoiceID:         invoice.ID,
				SellableProductID: sellableProduct.ID,
				Quantity:          purchasedItem.Qty,
				CompanyID:         companyID,
				Price:             purchasedItem.PriceAll,
				PromoID:           &purchasedItem.PromoID,
				PromoAmount:       &purchasedItem.PromoAmount,
			}); err != nil {
			trx.Rollback()
			return nil, http.StatusInternalServerError, err
		}

		if sellableProduct.HasReceipt {
			materialProductData, err := invoiceService.materialProductRepository.FindByStatus(companyID, true)
			if err != nil {
				trx.Rollback()
				return nil, http.StatusNotFound, err
			}

			if len(materialProductData) == 0 {
				trx.Rollback()
				return nil, http.StatusNotFound, errors.New("material product not found")
			}

			receiptAllProduct, err := invoiceService.receiptProductRepository.FindAll(sellableProduct.ID)
			if err != nil {
				trx.Rollback()
				return nil, http.StatusNotFound, err
			}

			if len(receiptAllProduct) == 0 {
				trx.Rollback()
				return nil, http.StatusNotFound, errors.New("receipt product not found")
			}

			for _, receiptProduct := range receiptAllProduct {
				for _, materialProduct := range materialProductData {
					if materialProduct.ID == receiptProduct.MaterialProductID {
						countQuantity := purchasedItem.Qty * receiptProduct.Quantity

						if materialProduct.CurrentQuantity < countQuantity {
							return nil, http.StatusBadRequest, errors.New("stock not enough")
						}

						if err = invoiceService.materialProductRepository.UpdateCurrentQty(trx, materialProduct.ID, countQuantity); err != nil {
							trx.Rollback()
							return nil, http.StatusInternalServerError, err
						}

						materialStock, err := invoiceService.materialStockRepository.FindByMaterialNotExp(materialProduct.ID)
						if err != nil {
							return nil, http.StatusNotFound, err
						}

						requiredQuantity := purchasedItem.Qty * receiptProduct.Quantity
						remainingQuantity := requiredQuantity

						for _, materialStockData := range materialStock {

							if materialStockData.CurrentQuantity > 0 {
								if materialStockData.CurrentQuantity < remainingQuantity {
									// set sisa remainingQuantity
									remainingQuantity = remainingQuantity - materialStockData.CurrentQuantity
									// then update current_quantity to 0 in latest fifo
									if err = invoiceService.materialStockRepository.UpdateCurrentQty(trx, materialStockData.ID, materialStockData.CurrentQuantity); err != nil {
										trx.Rollback()
										return nil, http.StatusInternalServerError, err
									}

								} else {
									if err = invoiceService.materialStockRepository.UpdateCurrentQty(trx, materialStockData.ID, remainingQuantity); err != nil {
										trx.Rollback()
										return nil, http.StatusInternalServerError, err
									}
									break
								}
							}
						}
					}
				}
			}
		} else {
			sellableStock, err := invoiceService.sellableStockRepository.FindBySellableStockNotExp(sellableProduct.ID)

			if err != nil {
				return nil, http.StatusNotFound, err
			}

			quantityClient := purchasedItem.Qty

			for _, sellableStockData := range sellableStock {
				if sellableStockData.CurrentQuantity < quantityClient {
					quantityClient = quantityClient - sellableStockData.CurrentQuantity
					if err = invoiceService.sellableStockRepository.UpdateCurrentQty(trx, sellableStockData.ID, sellableStockData.CurrentQuantity); err != nil {
						trx.Rollback()
						return nil, http.StatusInternalServerError, err
					}
				} else {
					if err = invoiceService.sellableStockRepository.UpdateCurrentQty(trx, sellableStockData.ID, quantityClient); err != nil {
						trx.Rollback()
						return nil, http.StatusInternalServerError, err
					}
					break
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
