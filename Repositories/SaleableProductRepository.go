package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	ISaleableProductRepository interface {
		FindAll(company_id int) (saleableProduct *[]Models.SaleableProduct, err error)
		FindByCategory(company_id int, category_id int) (saleableProduct *[]Models.SaleableProduct, err error)
	}

	SaleableProductRepository struct {
		DB *gorm.DB
	}
)

func SaleableProductRepositoryProvider(db *gorm.DB) *SaleableProductRepository {
	return &SaleableProductRepository{DB: db}
}

func (r *SaleableProductRepository) FindAll(company_id int) (saleableProduct *[]Models.SaleableProduct, err error) {

	saleableProduct = &[]Models.SaleableProduct{}

	if err := r.DB.Where("company_id = ?", company_id).
		Preload("Category").
		Find(&saleableProduct).Error; err != nil {
		return nil, err
	}

	return saleableProduct, nil
}

func (r *SaleableProductRepository) FindByCategory(company_id int, category_id int) (saleableProduct *[]Models.SaleableProduct, err error) {

	saleableProduct = &[]Models.SaleableProduct{}

	if err := r.DB.Where("company_id = ? AND category_id = ?", company_id, category_id).
		Preload("Category").
		Find(&saleableProduct).Error; err != nil {
		return nil, err
	}

	return saleableProduct, nil
}
