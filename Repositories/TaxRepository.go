package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	ITaxRepository interface {
		GetAll() (taxs []*Models.Tax, err error)
	}

	TaxRepository struct {
		DB *gorm.DB
	}
)

func TaxRepositoryProvider(db *gorm.DB) *TaxRepository {
	return &TaxRepository{DB: db}
}

func (r *TaxRepository) GetAll() (taxs []*Models.Tax, err error) {
	if err := r.DB.Find(&taxs).Error; err != nil {
		return nil, err
	}

	return taxs, nil
}
