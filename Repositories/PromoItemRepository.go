package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IPromoItemRepository interface {
		UpdateOrCreate(promoItem *Models.PromoItem) (*Models.PromoItem, *gorm.DB, error)
		DeleteBySellableProductID(sellableProductID string) (*gorm.DB, error)
	}

	PromoItemRepository struct {
		DB *gorm.DB
	}
)

func PromoItemRepositoryProvider(db *gorm.DB) *PromoItemRepository {
	return &PromoItemRepository{DB: db}
}

func (promoItemRepository *PromoItemRepository) UpdateOrCreate(promoItem *Models.PromoItem) (*Models.PromoItem, *gorm.DB, error) {
	result := promoItemRepository.DB.
		Where("sellable_product_id = ?", promoItem.SellableProductID).
		Assign(&Models.PromoItem{
			PromoID:           promoItem.PromoID,
			SellableProductID: promoItem.SellableProductID,
		}).
		FirstOrCreate(&promoItem)

	if result.Error != nil {
		return nil, nil, result.Error
	}

	return promoItem, result, nil
}

func (promoItemRepository *PromoItemRepository) DeleteBySellableProductID(sellableProductID string) (*gorm.DB, error) {
	result := promoItemRepository.DB.
		Where("sellable_product_id = ?", sellableProductID).
		Delete(&Models.PromoItem{})

	return result, result.Error
}
