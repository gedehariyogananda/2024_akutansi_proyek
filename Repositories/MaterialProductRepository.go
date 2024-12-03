package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IMaterialProductRepository interface {
		FindByAvailableForSale(company_id string) (materialProduct *[]Models.MaterialProduct, err error)
		Create(materialProduct *Models.MaterialProduct) (*Models.MaterialProduct, error)
	}

	MaterialProductRepository struct {
		DB *gorm.DB
	}
)

func MaterialProductRepositoryProvider(db *gorm.DB) *MaterialProductRepository {
	return &MaterialProductRepository{DB: db}
}

func (r *MaterialProductRepository) FindByAvailableForSale(company_id string) (materialProduct *[]Models.MaterialProduct, err error) {

	materialProduct = &[]Models.MaterialProduct{}

	if err := r.DB.Where("company_id = ?", company_id).
		Where("is_available_for_sale = ?", true).
		Find(&materialProduct).Error; err != nil {
		return nil, err
	}

	return materialProduct, nil
}

func (r *MaterialProductRepository) Create(materialProduct *Models.MaterialProduct) (*Models.MaterialProduct, error) {
	if err := r.DB.Create(materialProduct).Error; err != nil {
		return nil, err
	}

	return materialProduct, nil
}
