package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	ISaleableProductRepository interface {
		FindAll(company_id string) (saleableProduct *[]Models.SaleableProduct, err error)
		FindByCategory(company_id string, category_ids []string) (saleableProduct *[]Models.SaleableProduct, err error)
		CheckProductExist(company_id string, productId string) (isExist bool, err error)
		FindByName(company_id string, productName string, categoryId []string) (saleableProduct *[]Models.SaleableProduct, err error)
	}

	SaleableProductRepository struct {
		DB *gorm.DB
	}
)

func SaleableProductRepositoryProvider(db *gorm.DB) *SaleableProductRepository {
	return &SaleableProductRepository{DB: db}
}

func (r *SaleableProductRepository) FindAll(company_id string) (saleableProduct *[]Models.SaleableProduct, err error) {

	saleableProduct = &[]Models.SaleableProduct{}

	if err := r.DB.Where("company_id = ?", company_id).
		Preload("Category").
		Find(&saleableProduct).Error; err != nil {
		return nil, err
	}

	return saleableProduct, nil
}

func (r *SaleableProductRepository) FindByCategory(company_id string, category_ids []string) (saleableProduct *[]Models.SaleableProduct, err error) {
	saleableProduct = &[]Models.SaleableProduct{}

	if err := r.DB.Where("company_id = ? AND category_id IN (?)", company_id, category_ids).
		Preload("Category").
		Find(&saleableProduct).Error; err != nil {
		return nil, err
	}

	return saleableProduct, nil
}

func (r *SaleableProductRepository) CheckProductExist(company_id string, productId string) (isExist bool, err error) {
	saleableProduct := &Models.SaleableProduct{}

	if err := r.DB.Where("company_id = ? AND id = ?", company_id, productId).
		First(&saleableProduct).Error; err != nil {
		return false, nil
	}

	return true, nil
}

func (r *SaleableProductRepository) FindByName(company_id string, productName string, categoryId []string) (saleableProduct *[]Models.SaleableProduct, err error) {
	saleableProduct = &[]Models.SaleableProduct{}

	query := r.DB.Where("company_id = ? AND product_name LIKE ?", company_id, "%"+productName+"%")

	if len(categoryId) > 0 {
		query = query.Where("category_id IN ?", categoryId)
	}

	if err := query.Preload("Category").Find(&saleableProduct).Error; err != nil {
		return nil, err
	}

	return saleableProduct, nil
}
