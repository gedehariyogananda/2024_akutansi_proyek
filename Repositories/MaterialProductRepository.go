package Repositories

import (
	"2024_akutansi_project/Models"
	"fmt"

	"gorm.io/gorm"
)

type (
	IMaterialProductRepository interface {
		FindByStatus(companyID string, status bool) ([]*Models.MaterialProduct, error)
		UpdateCurrentQty(trx *gorm.DB, materialProductId string, qtyClient int) error
	}

	MaterialProductRepository struct {
		DB *gorm.DB
	}
)

func MaterialProductRepositoryProvider(db *gorm.DB) *MaterialProductRepository {
	return &MaterialProductRepository{DB: db}
}

func (r *MaterialProductRepository) FindByStatus(companyID string, status bool) ([]*Models.MaterialProduct, error) {
	var materialProduct []*Models.MaterialProduct

	if err := r.DB.Where("company_id = ? AND status = ?", companyID, status).Preload("Unit").Find(&materialProduct).Error; err != nil {
		return nil, fmt.Errorf("material product not found: %w", err)
	}

	return materialProduct, nil
}

func (r *MaterialProductRepository) UpdateCurrentQty(trx *gorm.DB, materialProductId string, qtyClient int) error {
	if err := trx.Model(&Models.MaterialProduct{}).
		Where("id = ?", materialProductId).
		Update("current_quantity", gorm.Expr("current_quantity - ?", qtyClient)).Error; err != nil {
		return fmt.Errorf("error when updating stock: %w", err)
	}

	return nil
}
