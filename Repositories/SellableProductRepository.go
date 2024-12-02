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
		GetAll(companyID string, status *bool, query *Common.Query) (sellableProducts []*Models.SellableProduct, totalData int64, err error)
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

func (sellableProductRepository *SellableProductRepository) GetAll(companyID string, status *bool, query *Common.Query) (sellableProducts []*Models.SellableProduct, totalData int64, err error) {
	if err := sellableProductRepository.DB.Model(&Models.SellableProduct{}).
		Scopes(Helper.FilterSearch(*query.Search)).
		Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	db := sellableProductRepository.DB.Model(&Models.SellableProduct{}).
		Where("company_id = ?", companyID)

	if status != nil {
		db = db.Where("status = ?", status)
	}

	if err := db.
		Preload("PromoItems", func(promoItemPayload *gorm.DB) *gorm.DB {
			return promoItemPayload.Preload("Promo", func(promoItem *gorm.DB) *gorm.DB {
				return promoItem.Select("id, name, start_date, end_date")
			})
		}).
		Preload("Category", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Scopes(
			Utils.Paginate(query.Page, query.Limit),
			Helper.FilterSearch(*query.Search),
		).Find(&sellableProducts).Error; err != nil {
		return nil, 0, err
	}

	return sellableProducts, totalData, nil
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
