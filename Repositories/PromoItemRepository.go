package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IPromoItemRepository interface {
		FindBySellableID(sellableProductID string) (*Models.PromoItem, bool, error)
		UpdateOrCreate(promoItem *Models.PromoItem) (*Models.PromoItem, *gorm.DB, error)
		DeleteBySellableProductID(sellableProductID string) (*gorm.DB, error)
		Create(promoItem *Models.PromoItem) (*Models.PromoItem, error)
		DeleteByPromoID(promoID string) error
		FindByPromoID(promoID string) ([]Models.PromoItem, error)
	}

	PromoItemRepository struct {
		DB *gorm.DB
	}
)

func PromoItemRepositoryProvider(db *gorm.DB) *PromoItemRepository {
	return &PromoItemRepository{DB: db}
}

func (promoItemRepository *PromoItemRepository) FindBySellableID(sellableProductID string) (*Models.PromoItem, bool, error) {
	promoItem := &Models.PromoItem{}
	if err := promoItemRepository.DB.
		Where("sellable_product_id = ?", sellableProductID).
		Preload("Promo").
		First(&promoItem).Error; err != nil {
		return nil, false, err
	}

	return promoItem, true, nil
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

func (promoItemRepository *PromoItemRepository) DeleteBySellableProductID(sellableProductID string) error {
	if err := promoItemRepository.DB.
		Where("sellable_product_id = ?", sellableProductID).
		Delete(&Models.PromoItem{}); err != nil {
		return err.Error
	}

	return nil
}

func (promoItemRepository *PromoItemRepository) Create(promoItem *Models.PromoItem) (*Models.PromoItem, error) {
	if err := promoItemRepository.DB.Create(promoItem).Error; err != nil {
		return nil, err
	}

	return promoItem, nil
}

func (promoItemRepository *PromoItemRepository) DeleteByPromoID(promoID string) error {
	if err := promoItemRepository.DB.Where("promo_id = ?", promoID).Delete(&Models.PromoItem{}).Error; err != nil {
		return err
	}

	return nil
}

func (promoItemRepository *PromoItemRepository) FindByPromoID(promoID string) ([]Models.PromoItem, error) {
	var promoItems []Models.PromoItem

	if err := promoItemRepository.DB.Where("promo_id = ?", promoID).Preload("Promo").Find(&promoItems).Error; err != nil {
		return nil, err
	}

	return promoItems, nil
}
