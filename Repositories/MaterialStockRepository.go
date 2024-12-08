package Repositories

import (
	"2024_akutansi_project/Models"
	"fmt"

	"gorm.io/gorm"
)

type (
	IMaterialStockRepository interface {
		FindByMaterialNotExp(materialStockID string) ([]*Models.MaterialStock, error)
		UpdateCurrent(trx *gorm.DB, materialStockID string, qtyClient int) error
		GetAvailableStock(companyId string) (materialStock []*Models.MaterialStock, err error)
	}

	MaterialStockRepository struct {
		DB *gorm.DB
	}
)

func MaterialStockRepositoryProvider(db *gorm.DB) *MaterialStockRepository {
	return &MaterialStockRepository{DB: db}
}

func (r *MaterialStockRepository) FindByMaterialNotExp(materialStockID string) ([]*Models.MaterialStock, error) {
	var materialStock []*Models.MaterialStock

	if err := r.DB.Where("material_product_id = ?", materialStockID).
		Where("expired_date > now()").
		Order("created_at asc").
		Find(&materialStock).Error; err != nil {
		return nil, fmt.Errorf("material stock not found: %w", err)
	}

	return materialStock, nil
}

func (r *MaterialStockRepository) UpdateCurrent(trx *gorm.DB, materialStockID string, qtyClient int) error {

	db := trx
	if db == nil {
		db = r.DB
	}

	if err := db.Model(&Models.MaterialStock{}).
		Where("id = ?", materialStockID).
		Update("current_quantity", gorm.Expr("current_quantity - ?", qtyClient)).Error; err != nil {
		return fmt.Errorf("error when updating stock: %w", err)
	}

	return nil
}

func (r *MaterialStockRepository) GetAvailableStock(companyId string) (materialStock []*Models.MaterialStock, err error) {
	if err := r.DB.
		Preload("MaterialProduct").
		Where("company_id = ?", companyId).
		Find(&materialStock).Error; err != nil {
		return nil, fmt.Errorf("error when getting available stock: %w", err)
	}

	return materialStock, nil
}
