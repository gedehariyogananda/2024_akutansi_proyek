package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Utils"
	"fmt"
	"time"
)

type (
	IPurchaseService interface {
		GetDropdown(companyID string) (res Response.DropDownPurchase, err error)
		Create(dto Dto.CreatePurchasesDto) (err error)
		GetAllPurchaseWithStatistic(companyID string, query *Common.Query) (res Response.PurchasesWithStatisticResponse, meta Common.Meta, err error)
		Delete(id string) (err error)
	}

	PurchaseService struct {
		ProductRepository          Repositories.ISellableProductRepository
		MaterialRepository         Repositories.IMaterialProductRepository
		MaterialStockRepository    Repositories.IMaterialStockRepository
		ProductStockRepository     Repositories.ISellableStockRepository
		purchaseProductRepository  Repositories.IPurchaseSellableProductRepository
		purchaseMaterialRepository Repositories.IPurchaseMaterialProductRepository
		purchaseRepository         Repositories.IPurchaseRepository
	}
)

func PurchaseServiceProvider(
	productRepository Repositories.ISellableProductRepository,
	materialRepository Repositories.IMaterialProductRepository,
	materialStockRepository Repositories.IMaterialStockRepository,
	productStockRepository Repositories.ISellableStockRepository,
	purchaseProductRepository Repositories.IPurchaseSellableProductRepository,
	purchaseMaterialRepository Repositories.IPurchaseMaterialProductRepository,
	purchaseRepository Repositories.IPurchaseRepository,
) *PurchaseService {
	return &PurchaseService{
		ProductRepository:          productRepository,
		MaterialRepository:         materialRepository,
		MaterialStockRepository:    materialStockRepository,
		ProductStockRepository:     productStockRepository,
		purchaseProductRepository:  purchaseProductRepository,
		purchaseMaterialRepository: purchaseMaterialRepository,
		purchaseRepository:         purchaseRepository,
	}
}

func (p *PurchaseService) GetDropdown(companyID string) (res Response.DropDownPurchase, err error) {
	products, err := p.ProductRepository.FindAllFilterReceipt(companyID, false)
	if err != nil {
		return res, err
	}

	materials, err := p.MaterialRepository.FindByCompany(companyID)
	if err != nil {
		return res, err
	}

	productsValues := make([]Models.SellableProduct, len(products))
	for i, product := range products {
		productsValues[i] = *product
	}

	materialsValues := make([]Models.MaterialProduct, len(materials))
	for i, material := range materials {
		materialsValues[i] = *material
	}

	res = Response.ToDropDownPurchase(productsValues, materialsValues)

	return res, nil
}

func (p *PurchaseService) Create(dto Dto.CreatePurchasesDto) (err error) {
	var dueDate *time.Time

	if dto.DueDate != "" {
		fomatedDate := Utils.ParseDateStringToDate(dto.DueDate, nil)
		dueDate = &fomatedDate
	}

	if err != nil {
		return err
	}

	puchase := Models.Purchase{
		TotalPurchaseAmount: dto.TotalPurchaseAmount,
		CompanyID:           dto.CompanyID,
		Tax:                 dto.Tax,
		Discount:            dto.Discount,
		Payment:             dto.Payment,
		DueDate:             dueDate,
		IsDiscountPercent:   dto.IsDiscountPercent,
	}

	purchase, err := p.purchaseRepository.Create(&puchase)

	if err != nil {
		return err
	}

	for _, product := range dto.Purchases {
		if product.Type == "product" {
			formatedDate := Utils.ParseDateStringToDate(product.ExpDate, nil)
			expDate := &formatedDate

			if err != nil {
				return err
			}

			stockProduct := Models.SellableStock{
				SellableProductID: product.ID,
				Quantity:          product.Quantity,
				ExpiredDate:       *expDate,
				CompanyID:         dto.CompanyID,
			}

			_, err = p.ProductStockRepository.Create(&stockProduct)

			if err != nil {
				return err
			}

			purchaseProduct := Models.PurchaseSellableProduct{
				PurchaseID:        purchase.ID,
				SellableProductID: product.ID,
				UnitPrice:         product.UnitPrice,
				Quantity:          product.Quantity,
				CompanyID:         dto.CompanyID,
			}
			_, err = p.purchaseProductRepository.Create(&purchaseProduct)

			if err != nil {
				return err
			}
		} else {
			fomateddate := Utils.ParseDateStringToDate(product.ExpDate, nil)
			expDate := &fomateddate

			if err != nil {
				return err
			}

			stockMaterial := Models.MaterialStock{
				MaterialProductID: product.ID,
				Quantity:          product.Quantity,
				ExpiredDate:       *expDate,
				CompanyID:         dto.CompanyID,
			}

			_, err = p.MaterialStockRepository.Create(&stockMaterial)

			if err != nil {
				return err
			}

			purchaseMaterial := Models.PurchaseMaterialProduct{
				PurchaseID:        purchase.ID,
				MaterialProductID: product.ID,
				UnitPrice:         product.UnitPrice,
				Quantity:          product.Quantity,
				CompanyID:         dto.CompanyID,
			}

			_, err = p.purchaseMaterialRepository.Create(&purchaseMaterial)

			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (p *PurchaseService) GetAllPurchaseWithStatistic(companyID string, query *Common.Query) (res Response.PurchasesWithStatisticResponse, meta Common.Meta, err error) {
	purchases, total, err := p.purchaseRepository.FindAll(companyID, query)

	if err != nil {
		return
	}

	totalPurchaseMonth, err := p.purchaseRepository.GetTotalPurchaseMonth(companyID)

	if err != nil {
		return
	}

	fmt.Println("totalPurchaseMonth", totalPurchaseMonth)

	purchaseStatistic := Response.StatisTicPurchase{
		PurchaseMonth: float32(totalPurchaseMonth),
	}

	meta = Common.Meta{
		Page:      query.Page,
		Limit:     query.Limit,
		TotalData: int64(total),
		TotalPage: int64(total / query.Limit),
	}

	res = Response.ToPurchaseResponseSlice(purchases, purchaseStatistic)

	return res, meta, nil
}

func (p *PurchaseService) Delete(id string) (err error) {
	err = p.purchaseMaterialRepository.DeleteByPurchaseID(id)

	if err != nil {
		return
	}

	err = p.purchaseProductRepository.DeleteByPurchaseID(id)

	if err != nil {
		return
	}

	err = p.purchaseRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
