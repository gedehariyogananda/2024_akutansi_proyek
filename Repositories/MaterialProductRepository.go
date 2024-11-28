package Repositories

import (
	"2024_akutansi_project/Models"
	"fmt"

	"gorm.io/gorm"
)

type (
	IMaterialProductRepository interface {
		FindByStatus(companyID string, status bool) ([]*Models.MaterialProduct, error)
		Update(materialProductId string, materialProduct *Models.MaterialProduct) error
		UpdateTrx(trx *gorm.DB, materialProductId string, materialProduct *Models.MaterialProduct) error
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

	if err := r.DB.Where("company_id = ? AND status = ?", companyID, status).Find(&materialProduct).Error; err != nil {
		return nil, fmt.Errorf("material product with status %w not found: %w", status, err)
	}

	return materialProduct, nil
}

func (r *MaterialProductRepository) Update(materialProductId string, materialProduct *Models.MaterialProduct) error {
	if err := r.DB.Model(&Models.MaterialProduct{}).
		Where("id = ?", materialProductId).
		Updates(materialProduct).Error; err != nil {
		return fmt.Errorf("material product %w not updated: %w", materialProductId, err)
	}

	return nil
}

func (r *MaterialProductRepository) UpdateTrx(trx *gorm.DB, materialProductId string, materialProduct *Models.MaterialProduct) error {
	if err := trx.Model(&Models.MaterialProduct{}).
		Where("id = ?", materialProductId).
		Updates(materialProduct).Error; err != nil {
		return fmt.Errorf("material product %w not updated: %w", materialProductId, err)
	}

	return nil
}
