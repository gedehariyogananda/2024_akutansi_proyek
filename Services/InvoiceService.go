package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Utils"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type (
	IInvoiceService interface {
		CreateInvoicePurchased(requestClient *Dto.InvoiceRequestDTO, companyID string) (invoice *Models.Invoice, statusCode int, err error)
		GetAllByCompany(companyID string, query *Common.Query) (response []Response.InvoiceResponse, meta Common.Meta, statusCode int, err error)
		GetSpesifySalesHistory(companyID string, invoiceID string) (response Response.InvoiceResponse, statusCode int, err error)
		UpdateRefund(companyID string, id string) (statusCode int, err error)
		StatisticSales(companyID string, date string) (data interface{}, statusCode int, err error)
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
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
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
			return nil, http.StatusForbidden, errors.New("FORBIDDEN_ACCESS")
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

func (invoiceService *InvoiceService) GetAllByCompany(companyID string, query *Common.Query) (response []Response.InvoiceResponse, meta Common.Meta, statusCode int, err error) {
	invoices, totalData, err := invoiceService.invoiceRepository.GetAllByCompany(companyID, query)

	if err != nil {
		return nil, Common.Meta{}, http.StatusInternalServerError, err
	}

	var res []Response.InvoiceResponse

	for _, invoice := range invoices {
		total := 0
		for _, item := range invoice.InvoiceItems {
			total += item.Quantity
		}

		status := ""

		if invoice.Status {
			status = "Lunas"
		} else {
			status = "Belum Lunas"
		}

		res = append(res, Response.InvoiceResponse{
			ID:            invoice.ID,
			CustomerName:  invoice.CustomerName,
			InvoiceNumber: invoice.InvoiceNumber,
			SubTotal:      invoice.SubTotal,
			Status:        &status,
			CreatedAt:     invoice.CreatedAt.Format("02/01/2006"),
			CountSale:     &total,
		})

	}

	meta = Common.Meta{
		TotalData: totalData,
		Limit:     query.Limit,
		Page:      query.Page,
	}

	return res, meta, http.StatusOK, nil
}

func (invoiceService *InvoiceService) GetSpesifySalesHistory(companyID string, invoiceID string) (response Response.InvoiceResponse, statusCode int, err error) {
	invoice, _ := invoiceService.invoiceRepository.GetByInvoiceID(companyID, invoiceID)

	status := ""

	if invoice.Status {
		status = "Lunas"
	} else {
		status = "Belum Lunas"
	}

	total := 0

	for _, item := range invoice.InvoiceItems {
		total += item.Quantity
	}

	res := Response.InvoiceResponse{
		ID:           invoice.ID,
		CustomerName: invoice.CustomerName,
		PhoneNumber:  invoice.PhoneNumber,
		CreatedAt:    invoice.CreatedAt.Format("02/01/2006"),
		Status:       &status,
		Note:         &invoice.Note,
		SubTotal:     invoice.SubTotal,
		Tax:          &invoice.Tax,
		CountSale:    &total,
		InvoiceItems: &invoice.InvoiceItems,
	}

	return res, http.StatusOK, nil
}

func (invoiceService *InvoiceService) UpdateRefund(companyID string, id string) (statusCode int, err error) {
	invoice, err := invoiceService.invoiceRepository.FindByID(id, companyID)

	if err != nil {
		return http.StatusNotFound, err
	}
	if invoice.RefundAt != nil {
		return http.StatusBadRequest, errors.New("invoice sudah di refund")
	}

	if err = invoiceService.invoiceRepository.Update(id, &Models.Invoice{
		RefundAt: func() *time.Time {
			now := time.Now()
			return &now
		}(),
	}); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (invoiceService *InvoiceService) StatisticSales(companyID string, date string) (data interface{}, statusCode int, err error) {

	yearInit, _ := strconv.Atoi(strings.Split(date, "-")[0])
	monthInit, _ := strconv.Atoi(strings.Split(date, "-")[1])
	prevMonth := monthInit - 1

	// set safety first and latest month init
	if prevMonth < 1 {
		prevMonth = 12
		yearInit -= 1
	}

	currentDate, _ := time.Parse("2006-01-02", date)
	prevDay := currentDate.AddDate(0, 0, -1).Format("2006-01-02")

	sumSalesNow, _ := invoiceService.invoiceRepository.SumSalesByDate(companyID, date)
	sumSalesPrev, _ := invoiceService.invoiceRepository.SumSalesByDate(companyID, prevDay)

	sumSalesNowByMonth, _ := invoiceService.invoiceRepository.SumSalesByYearMonth(companyID, yearInit, monthInit)
	sumSalesPrevByMonth, _ := invoiceService.invoiceRepository.SumSalesByYearMonth(companyID, yearInit, prevMonth)

	// calculate peresentage kenaikan
	salesNowPercentage := Utils.CalculatePercentageInit(sumSalesPrev, sumSalesNow)
	salesMonthPercentage := Utils.CalculatePercentageInit(sumSalesPrevByMonth, sumSalesNowByMonth)

	mostProductSold, err := invoiceService.invoiceItemRepository.GetMostProductSold(companyID, date)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	data = map[string]interface{}{
		"sum_sales_now": map[string]interface{}{
			"total":      sumSalesNow,
			"percentage": salesNowPercentage,
		},
		"sum_sales_now_by_month": map[string]interface{}{
			"total":      sumSalesNowByMonth,
			"percentage": salesMonthPercentage,
		},
		"most_product_sold": map[string]interface{}{
			"name":       mostProductSold.ProductName,
			"count_sale": mostProductSold.CountSale,
		},
	}

	return data, http.StatusOK, nil
}
