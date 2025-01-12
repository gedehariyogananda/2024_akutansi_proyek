package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Utils"
	"fmt"

	"gorm.io/gorm"
)

type (
	ISellableProductRepository interface {
		GetAll(companyID string, onlyActive bool, query *Common.Query) (sellableProducts []*Models.SellableProduct, totalData int64, err error)
		Update(id string, sellableProduct *Models.SellableProduct) error
		Find(id string) (sellableProduct *Models.SellableProduct, err error)
		UpdateCurrent(trx *gorm.DB, sellableProductID string, QtyClient int) error
		Create(sellableProduct *Models.SellableProduct) (*Models.SellableProduct, error)
		Delete(id string) error
		FindByID(id string) (*Models.SellableProduct, error)
		FindAllFilterReceipt(companyId string, HasReceipt bool) (sellableProducts []*Models.SellableProduct, err error)
	}

	SellableProductRepository struct {
		DB *gorm.DB
	}
)

func SellableProductRepositoryProvider(db *gorm.DB) *SellableProductRepository {
	return &SellableProductRepository{DB: db}
}

func (sellableProductRepository *SellableProductRepository) GetAll(companyID string, onlyActive bool, query *Common.Query) (sellableProducts []*Models.SellableProduct, totalData int64, err error) {
	totalCountInit := sellableProductRepository.DB.Model(&Models.SellableProduct{}).
		Where("company_id = ?", companyID)

	if onlyActive {
		totalCountInit = totalCountInit.Where("status = ?", true)
	}

	if err := totalCountInit.Scopes(
		Helper.FilterSearchProduct(query.Search),
		Helper.FilterCategoryID(query.CategoryID)).
		Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	db := sellableProductRepository.DB.Model(&Models.SellableProduct{}).
		Where("company_id = ?", companyID)

	if onlyActive {
		db = db.Where("status = ?", true).
			Select("id", "name", "image", "description", "current_quantity", "price", "status")
	} else {
		db = db.Preload("PromoItems", func(promoItemPayload *gorm.DB) *gorm.DB {
			return promoItemPayload.Preload("Promo", func(promoItem *gorm.DB) *gorm.DB {
				return promoItem.Select("id, name, start_date, end_date")
			})
		}).
			Preload("Category", func(db *gorm.DB) *gorm.DB {
				return db.Select("id, name")
			})
	}

	if err := db.Scopes(
		Utils.Paginate(query.Page, query.Limit),
		Helper.FilterSearchProduct(query.Search),
		Helper.FilterCategoryID(query.CategoryID),
	).Find(&sellableProducts).Error; err != nil {
		return nil, 0, err
	}

	return sellableProducts, totalData, nil
}

func (sellableProductRepository *SellableProductRepository) Update(id string, sellableProduct *Models.SellableProduct) error {
	if err := sellableProductRepository.DB.Model(&Models.SellableProduct{}).
		Where("id = ?", id).
		Updates(sellableProduct).Error; err != nil {
		return fmt.Errorf("error saat update sellable products: %w", err)
	}

	return nil
}

func (sellableProductRepository *SellableProductRepository) Find(id string) (sellableProduct *Models.SellableProduct, err error) {
	sellableProduct = &Models.SellableProduct{}

	if err = sellableProductRepository.DB.Where("id = ?", id).Preload("Unit").First(sellableProduct).Error; err != nil {
		return nil, fmt.Errorf("sellable product not found: %w", err)
	}

	return sellableProduct, nil
}

func (sellableProductRepository *SellableProductRepository) UpdateCurrent(trx *gorm.DB, sellableProductID string, qtyClient int) error {

	db := trx
	if db == nil {
		db = sellableProductRepository.DB
	}

	if err := db.Model(&Models.SellableProduct{}).
		Where("id = ?", sellableProductID).
		Update("current_quantity", gorm.Expr("current_quantity - ?", qtyClient)).Error; err != nil {
		return fmt.Errorf("error when updating stock: %w", err)
	}

	return nil
}

func (sellableProductRepository *SellableProductRepository) Create(sellableProduct *Models.SellableProduct) (*Models.SellableProduct, error) {
	if err := sellableProductRepository.DB.Create(sellableProduct).Error; err != nil {
		return nil, fmt.Errorf("error saat membuat sellable product: %w", err)
	}

	return sellableProduct, nil
}

func (sellableProductRepository *SellableProductRepository) Delete(id string) error {
	if err := sellableProductRepository.DB.Where("id = ?", id).Delete(&Models.SellableProduct{}).Error; err != nil {
		return fmt.Errorf("error saat menghapus sellable product: %w", err)
	}

	return nil
}

func (sellableProductRepository *SellableProductRepository) FindByID(id string) (*Models.SellableProduct, error) {
	var sellableProduct Models.SellableProduct

	if err := sellableProductRepository.DB.Where("id = ?", id).Preload("Unit").Preload("Category").Preload("Receipts.MaterialProduct.Unit").First(&sellableProduct).Error; err != nil {
		return nil, fmt.Errorf("sellable product not found: %w", err)
	}

	return &sellableProduct, nil
}
func (sellableProductRepository SellableProductRepository) FindAllFilterReceipt(companyId string, HasReceipt bool) (sellableProducts []*Models.SellableProduct, err error) {
	if err := sellableProductRepository.DB.Model(&Models.SellableProduct{}).
		Where("has_receipt = ?", HasReceipt).
		Find(&sellableProducts).Error; err != nil {
		return nil, err
	}

	return sellableProducts, nil
}
