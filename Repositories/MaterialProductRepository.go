package Repositories

import (
	"2024_akutansi_project/Models"
	"fmt"

	"gorm.io/gorm"
)

type (
	IMaterialProductRepository interface {
		Create(materialProduct *Models.MaterialProduct) (*Models.MaterialProduct, error)
		FindByCompany(companyID string) ([]*Models.MaterialProduct, error)
		UpdateCurrent(trx *gorm.DB, materialProductId string, qtyClient int) error
		FindByID(id string) (*Models.MaterialProduct, error)
	}

	MaterialProductRepository struct {
		DB *gorm.DB
	}
)

func MaterialProductRepositoryProvider(db *gorm.DB) *MaterialProductRepository {
	return &MaterialProductRepository{DB: db}
}

func (r *MaterialProductRepository) Create(materialProduct *Models.MaterialProduct) (*Models.MaterialProduct, error) {
	if err := r.DB.Create(materialProduct).Error; err != nil {
		return nil, fmt.Errorf("error when creating material product: %w", err)
	}

	return materialProduct, nil
}

func (r *MaterialProductRepository) FindByCompany(companyID string) ([]*Models.MaterialProduct, error) {
	var materialProduct []*Models.MaterialProduct

	if err := r.DB.Where("company_id = ?", companyID).Preload("Unit").Find(&materialProduct).Error; err != nil {
		return nil, fmt.Errorf("material product not found: %w", err)
	}

	return materialProduct, nil
}

func (r *MaterialProductRepository) UpdateCurrent(trx *gorm.DB, materialProductId string, qtyClient int) error {

	db := trx
	if db == nil {
		db = r.DB
	}

	if err := db.Model(&Models.MaterialProduct{}).
		Where("id = ?", materialProductId).
		Update("current_quantity", gorm.Expr("current_quantity - ?", qtyClient)).Error; err != nil {
		return fmt.Errorf("error when updating stock: %w", err)
	}

	return nil
}

func (r *MaterialProductRepository) FindByID(id string) (*Models.MaterialProduct, error) {
	var materialProduct Models.MaterialProduct

	if err := r.DB.Where("id = ?", id).Preload("Unit").Preload("MaterialConversions.Unit").First(&materialProduct).Error; err != nil {
		return nil, fmt.Errorf("material product not found: %w", err)
	}

	fmt.Println(materialProduct.MaterialConversions)
	return &materialProduct, nil
}
