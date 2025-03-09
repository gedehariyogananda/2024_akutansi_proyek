package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type (
	ISellableProductRepository interface {
		GetAll(companyID string, query *Dto.GetSellableProduct) (sellableProducts []*Models.SellableProduct, totalData int64, err error)
		Update(id string, sellableProduct *Models.SellableProduct) error
		Find(id string) (sellableProduct *Models.SellableProduct, err error)
		UpdateCurrent(trx *gorm.DB, sellableProductID string, QtyClient int) error
		Create(sellableProduct *Models.SellableProduct) (*Models.SellableProduct, error)
		Delete(id string) error
		FindByID(id string, setWithMaterial bool) (*Models.SellableProduct, error)
		FindAllFilterReceipt(companyId string, HasReceipt bool) (sellableProducts []*Models.SellableProduct, err error)
		// FindByID(id string) (*Models.SellableProduct, error)

		// FindAll(companyID string, query *Common.Query) ([]*Models.SellableProduct, int64, error)
		GetByCompanyID(companyID string) (sellableProducts []Models.SellableProduct, err error)
	}

	SellableProductRepository struct {
		DB *gorm.DB
	}
)

func SellableProductRepositoryProvider(db *gorm.DB) *SellableProductRepository {
	return &SellableProductRepository{DB: db}
}

func (sellableProductRepository *SellableProductRepository) GetAll(companyID string, query *Dto.GetSellableProduct) (sellableProducts []*Models.SellableProduct, totalData int64, err error) {
	totalCountInit := sellableProductRepository.DB.Model(&Models.SellableProduct{}).
		Where("company_id = ?", companyID)

	if err := totalCountInit.Scopes(
		Helper.FilterSearchProduct(query.Search),
		Helper.FilterCategoryID(query.CategoryID),
		Helper.FilterManagementStock(*query.InStatus)).
		Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	db := sellableProductRepository.DB.Model(&Models.SellableProduct{}).
		Where("company_id = ?", companyID)

	db = db.Preload("PromoItems", func(promoItemPayload *gorm.DB) *gorm.DB {
		return promoItemPayload.Joins("JOIN promos ON promo_items.promo_id = promos.id").
			Where("DATE(promos.start_date) <= ? AND DATE(promos.end_date) >= ?", time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02")).
			Preload("Promo", func(promoPayload *gorm.DB) *gorm.DB {
				return promoPayload.Select("id, name, start_date, end_date, amount, type")
			}).
			Select("promo_items.promo_id, promo_items.sellable_product_id")
	}).Preload("Category", func(categoryPayload *gorm.DB) *gorm.DB {
		return categoryPayload.Select("id, name")
	})

	if err := db.Scopes(
		Utils.Paginate(query.Page, query.Limit),
		Helper.FilterSearchProduct(query.Search),
		Helper.FilterCategoryID(query.CategoryID),
		Helper.FilterManagementStock(*query.InStatus),
	).Order("created_at desc").Find(&sellableProducts).Error; err != nil {
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

func (sellableProductRepository *SellableProductRepository) FindByID(id string, setWithMaterial bool) (*Models.SellableProduct, error) {
	var sellableProduct Models.SellableProduct

	db := sellableProductRepository.DB.Where("id = ?", id).Preload("Category", func(categoryPayload *gorm.DB) *gorm.DB {
		return categoryPayload.Select("id, name")
	})

	if !setWithMaterial {
		db = db.Preload("PromoItems.Promo")
	} else {
		db = db.Preload("Receipts.MaterialProduct.Unit").Preload("Unit")
	}

	if err := db.First(&sellableProduct).Error; err != nil {
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

func (SellableProductRepository *SellableProductRepository) GetByCompanyID(companyID string) (sellableProducts []Models.SellableProduct, err error) {
	if err := SellableProductRepository.DB.Where("company_id = ?", companyID).Preload("PromoItems.Promo").Find(&sellableProducts).Error; err != nil {
		return nil, err
	}

	return sellableProducts, nil
}
