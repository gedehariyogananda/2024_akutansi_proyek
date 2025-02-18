package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Utils"

	"gorm.io/gorm"
)

type (
	IPromoRepository interface {
		Create(promo *Models.Promo) (*Models.Promo, error)
		FindById(id string) (*Models.Promo, error)
		Delete(id string) error
		Update(promo *Models.Promo, id string) (*Models.Promo, error)
		FindAll(companyID string, query Common.Query) ([]Models.Promo, int64, error)
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

func (r *PromoRepository) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&Models.Promo{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *PromoRepository) Update(promo *Models.Promo, id string) (*Models.Promo, error) {
	if err := r.DB.Where("id = ?", id).Updates(promo).Error; err != nil {
		return nil, err
	}

	return promo, nil
}

func (r *PromoRepository) FindAll(companyID string, query Common.Query) ([]Models.Promo, int64, error) {
	var promos []Models.Promo
	var total int64

	if err := r.DB.Preload("PromoItems.SellableProduct.Category").
		Scopes(Utils.Paginate(query.Page, query.Limit),
			Helper.FilterCompanyID(companyID),
			Helper.FilterSearch(*query.Search),
			Helper.FilterType(*query.Type),
		).Find(&promos).Error; err != nil {
		return nil, 0, err
	}

	err := r.DB.Model(&Models.Promo{}).
		Scopes(Helper.FilterCompanyID(companyID), Helper.FilterSearch(*query.Search), Helper.FilterType(*query.Type)).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	return promos, total, nil
}
