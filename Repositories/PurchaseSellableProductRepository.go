package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IPurchaseSellableProductRepository interface {
		Create(purchaseSellableProduct *Models.PurchaseSellableProduct) (*Models.PurchaseSellableProduct, error)
		DeleteByPurchaseID(purchaseID string) error
	}

	PurchaseSellableProductRepository struct {
		DB *gorm.DB
	}
)

func PurchaseSellableProductRepositoryProvider(db *gorm.DB) *PurchaseSellableProductRepository {
	return &PurchaseSellableProductRepository{DB: db}
}

func (r *PurchaseSellableProductRepository) Create(purchaseSellableProduct *Models.PurchaseSellableProduct) (*Models.PurchaseSellableProduct, error) {
	if err := r.DB.Create(purchaseSellableProduct).Error; err != nil {
		return nil, err
	}

	return purchaseSellableProduct, nil
}

func (r *PurchaseSellableProductRepository) DeleteByPurchaseID(purchaseID string) error {
	if err := r.DB.Where("purchase_id = ?", purchaseID).Delete(&Models.PurchaseSellableProduct{}).Error; err != nil {
		return err
	}

	return nil
}
