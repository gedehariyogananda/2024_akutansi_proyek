package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IPromoRepository interface {
		FindByID(id string) (*Models.Promo, error)
	}

	PromoRepository struct {
		DB *gorm.DB
	}
)

func PromoRepositoryProvider(db *gorm.DB) *PromoRepository {
	return &PromoRepository{DB: db}
}

func (promoRepository *PromoRepository) FindByID(id string) (*Models.Promo, error) {
	var promo *Models.Promo
	if err := promoRepository.DB.Where("id = ?", id).First(&promo).Error; err != nil {
		return nil, err
	}

	return promo, nil
}
