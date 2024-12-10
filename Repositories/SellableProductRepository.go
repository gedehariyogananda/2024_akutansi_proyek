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
	}

	SellableProductRepository struct {
		DB *gorm.DB
	}
)

func SellableProductRepositoryProvider(db *gorm.DB) *SellableProductRepository {
	return &SellableProductRepository{DB: db}
}

func (sellableProductRepository *SellableProductRepository) GetAll(companyID string, onlyActive bool, query *Common.Query) (sellableProducts []*Models.SellableProduct, totalData int64, err error) {
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
		Helper.FilterSearch(*query.Search),
	).Find(&sellableProducts).Error; err != nil {
		return nil, 0, err
	}

	totalData, err = Utils.CountModelRecords(db, &sellableProducts)
	if err != nil {
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
		return nil, fmt.Errorf("sellable product tidak ditemukan! : %w", err)
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
		return fmt.Errorf("ada kesalahan saat update stock! : %w", err)
	}

	return nil
}
