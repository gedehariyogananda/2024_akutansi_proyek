package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IPurchaseMaterialProductRepository interface {
		Create(purchaseMaterialProduct *Models.PurchaseMaterialProduct) (*Models.PurchaseMaterialProduct, error)
		DeleteByPurchaseID(purchaseID string) error
	}

	PurchaseMaterialProductRepository struct {
		DB *gorm.DB
	}
)

func PurchaseMaterialProductRepositoryProvider(db *gorm.DB) *PurchaseMaterialProductRepository {
	return &PurchaseMaterialProductRepository{DB: db}
}

func (r *PurchaseMaterialProductRepository) Create(purchaseMaterialProduct *Models.PurchaseMaterialProduct) (*Models.PurchaseMaterialProduct, error) {
	if err := r.DB.Create(purchaseMaterialProduct).Error; err != nil {
		return nil, err
	}

	return purchaseMaterialProduct, nil
}

func (r *PurchaseMaterialProductRepository) DeleteByPurchaseID(purchaseID string) error {
	if err := r.DB.Where("purchase_id = ?", purchaseID).Delete(&Models.PurchaseMaterialProduct{}).Error; err != nil {
		return err
	}

	return nil
}
