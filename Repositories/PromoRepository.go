package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IPromoRepository interface {
		Create(promo *Models.Promo) (*Models.Promo, error)
		FindById(id string) (*Models.Promo, error)
	}

	PromoRepository struct {
		DB *gorm.DB
	}
)

func PromoRepositoryProvider(db *gorm.DB) *PromoRepository {
	return &PromoRepository{DB: db}
}

func (r *PromoRepository) Create(promo *Models.Promo) (*Models.Promo, error) {
	if err := r.DB.Create(promo).Error; err != nil {
		return nil, err
	}

	return promo, nil
}

func (r *PromoRepository) FindById(id string) (*Models.Promo, error) {
	var promo Models.Promo

	if err := r.DB.Where("id = ?", id).Preload("PromoItems.SellableProduct.Category").First(&promo).Error; err != nil {
		return nil, err
	}

	return &promo, nil
}
