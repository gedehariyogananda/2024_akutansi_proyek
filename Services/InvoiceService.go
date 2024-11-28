package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	invoice, err = invoiceService.invoiceRepository.Store(trx, invoiceDataClient)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var lowStockErrors []string
	var expiredItemsErrors []string

	for _, purchasedItem := range requestClient.Purchaseds {
		sellableProduct, err := invoiceService.sellableProductRepository.Find(purchasedItem.ID)
		if err != nil {
			return nil, http.StatusNotFound, fmt.Errorf("product tidak ditemukan : %s", purchasedItem.ID)
		}

		if sellableProduct.CompanyID != companyID {
			return nil, http.StatusForbidden, errors.New("forbidden access")
		}

		// check stock availability
		if sellableProduct.CurrentQuantity < purchasedItem.Qty {
			lowStockErrors = append(lowStockErrors, fmt.Sprintf("stok produk %s tidak mencukupi (tersedia: %d %s, dibutuhkan: %d %s)",
				sellableProduct.Name, sellableProduct.CurrentQuantity, sellableProduct.Unit.Name, purchasedItem.Qty, sellableProduct.Unit.Name))
			continue
		}

		// update stock sellable product
		if err = invoiceService.sellableProductRepository.UpdateCurrent(trx, purchasedItem.ID, purchasedItem.Qty); err != nil {
			return nil, http.StatusInternalServerError, fmt.Errorf("gagal memperbarui stok produk %s: %v", sellableProduct.Name, err)
		}

		// add invoice item
		invoiceItem := &Models.InvoiceItem{
			InvoiceID:         invoice.ID,
			SellableProductID: sellableProduct.ID,
			Quantity:          purchasedItem.Qty,
			CompanyID:         companyID,
			Price:             purchasedItem.PriceAll,
			PromoID:           &purchasedItem.PromoID,
			PromoAmount:       &purchasedItem.PromoAmount,
		}

		if err = invoiceService.invoiceItemRepository.Store(trx, invoiceItem); err != nil {
			return nil, http.StatusInternalServerError, fmt.Errorf("gagal menyimpan item invoice untuk produk %s: %v", sellableProduct.Name, err)
		}

		// set logic hasReceipt true or false handling
		if sellableProduct.HasReceipt {
			if err = invoiceService.handleMaterialProducts(trx, sellableProduct, purchasedItem.Qty, &expiredItemsErrors, &lowStockErrors); err != nil {
				return nil, http.StatusInternalServerError, err
			}
		} else {
			if err = invoiceService.handleSellableStocks(trx, sellableProduct, purchasedItem.Qty, &expiredItemsErrors, &lowStockErrors); err != nil {
				return nil, http.StatusInternalServerError, err
			}
		}
	}

	if len(lowStockErrors) > 0 || len(expiredItemsErrors) > 0 {
		allErrors := append(lowStockErrors, expiredItemsErrors...)
		return nil, http.StatusBadRequest, fmt.Errorf("terdapat beberapa masalah: %v", strings.Join(allErrors, "; "))
	}

	return invoice, http.StatusOK, nil
}

func (invoiceService *InvoiceService) handleMaterialProducts(trx *gorm.DB, sellableProduct *Models.SellableProduct, qty int, expiredItemsErrors, lowStockErrors *[]string) error {
	materialProductData, err := invoiceService.materialProductRepository.FindByCompany(sellableProduct.CompanyID)
	if err != nil || len(materialProductData) == 0 {
		*expiredItemsErrors = append(*expiredItemsErrors, fmt.Sprintf("tidak ada bahan aktif yang ditemukan untuk produk %s", sellableProduct.Name))
		return nil
	}

	receiptProducts, err := invoiceService.receiptProductRepository.FindAll(sellableProduct.ID)
	if err != nil || len(receiptProducts) == 0 {
		*expiredItemsErrors = append(*expiredItemsErrors, fmt.Sprintf("tidak ada data resep yang ditemukan untuk produk %s", sellableProduct.Name))
		return nil
	}

	for _, receipt := range receiptProducts {
		for _, material := range materialProductData {
			if material.ID == receipt.MaterialProductID {
				requiredQty := qty * receipt.Quantity
				if material.CurrentQuantity < requiredQty {
					*lowStockErrors = append(*lowStockErrors, fmt.Sprintf(
						"stok bahan %s tidak mencukupi (dibutuhkan: %d %s, tersedia: %d %s)",
						material.Name, requiredQty, material.Unit.Name,
						material.CurrentQuantity, material.Unit.Name))
				}

				if material.CurrentQuantity >= requiredQty {

					if err = invoiceService.materialProductRepository.UpdateCurrent(trx, material.ID, requiredQty); err != nil {
						return err
					}

					materialStock, err := invoiceService.materialStockRepository.FindByMaterialNotExp(material.ID)
					if err != nil {
						return err
					}

					remainingQty := requiredQty

					for _, materialStockData := range materialStock {
						if materialStockData.CurrentQuantity > 0 {
							if materialStockData.CurrentQuantity < remainingQty {
								remainingQty -= materialStockData.CurrentQuantity
								if err = invoiceService.materialStockRepository.UpdateCurrent(trx, materialStockData.ID, materialStockData.CurrentQuantity); err != nil {
									trx.Rollback()
									return err
								}

							} else {
								if err = invoiceService.materialStockRepository.UpdateCurrent(trx, materialStockData.ID, remainingQty); err != nil {
									trx.Rollback()
									return err
								}

								break
							}
						}
					}
				}
			}
		}
	}

	return nil
}

func (invoiceService *InvoiceService) handleSellableStocks(trx *gorm.DB, sellableProduct *Models.SellableProduct, qty int, expiredItemsErrors, lowStockErrors *[]string) error {
	sellableStocks, err := invoiceService.sellableStockRepository.FindBySellableStockNotExp(sellableProduct.ID)
	if err != nil || len(sellableStocks) == 0 {
		*expiredItemsErrors = append(*expiredItemsErrors, fmt.Sprintf("tidak ada stok yang ditemukan untuk produk %s", sellableProduct.Name))
		return nil
	}

	sumCurrentStock, _ := invoiceService.sellableStockRepository.SumCurrentQuantity(sellableProduct.ID)
	totalAvailableQty := sumCurrentStock

	log.Println("log: totalAvailableQty", totalAvailableQty)

	// if stock != matched
	if totalAvailableQty < qty {
		*lowStockErrors = append(*lowStockErrors, fmt.Sprintf("stok produk %s tidak cukup (dibutuhkan: %d pcs, tersedia: %d pcs)",
			sellableProduct.Name, qty, totalAvailableQty))
		return nil
	}

	var insufficientStockDetails []string
	for _, stock := range sellableStocks {
		if stock.CurrentQuantity > 0 {
			if stock.CurrentQuantity <= qty {
				qty -= stock.CurrentQuantity
				if err := invoiceService.sellableStockRepository.UpdateCurrent(trx, stock.ID, stock.CurrentQuantity); err != nil {
					return err
				}
			} else {
				if err := invoiceService.sellableStockRepository.UpdateCurrent(trx, stock.ID, qty); err != nil {
					return err
				}
				break
			}
		}
	}

	if len(insufficientStockDetails) > 0 {
		*lowStockErrors = append(*lowStockErrors, fmt.Sprintf("terdapat beberapa masalah: %s", strings.Join(insufficientStockDetails, "; ")))
	}

	return nil
}
