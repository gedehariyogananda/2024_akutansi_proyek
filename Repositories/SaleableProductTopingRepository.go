package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"

	"gorm.io/gorm"
)

type (
	ISaleableProductTopingRepository interface {
		Create(request *Dto.TopingsItem, companyID string, saleableProductID string) (err error)
	}

	SaleableProductTopingRepository struct {
		DB *gorm.DB
	}
)

func SaleableProductTopingRepositoryProvider(db *gorm.DB) *SaleableProductTopingRepository {
	return &SaleableProductTopingRepository{DB: db}
}

func (repository *SaleableProductTopingRepository) Create(request *Dto.TopingsItem, companyID string, saleableProductID string) (err error) {
	productToping := &Models.SaleableProductToping{
		TopingID:          request.TopingID,
		SaleableProductID: saleableProductID,
		CompanyID:         companyID,
	}

	if err := repository.DB.Create(productToping).Error; err != nil {
		return err
	}

	return nil
}
