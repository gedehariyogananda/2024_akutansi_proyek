package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IPurchaseRepository interface {
		Create(purchase *Models.Purchase) (*Models.Purchase, error)
	}

	PurchaseRepository struct {
		DB *gorm.DB
	}
)

func PurchaseRepositoryProvider(db *gorm.DB) *PurchaseRepository {
	return &PurchaseRepository{DB: db}
}

func (r *PurchaseRepository) Create(purchase *Models.Purchase) (*Models.Purchase, error) {
	if err := r.DB.Create(purchase).Error; err != nil {
		return nil, err
	}

	return purchase, nil
}
