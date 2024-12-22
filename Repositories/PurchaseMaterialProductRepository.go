package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IPurchaseMaterialProductRepository interface {
		Create(purchaseMaterialProduct *Models.PurchaseMaterialProduct) (*Models.PurchaseMaterialProduct, error)
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
