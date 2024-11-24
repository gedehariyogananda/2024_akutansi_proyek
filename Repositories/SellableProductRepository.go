package Repositories

import (
	"2024_akutansi_project/Models"
	"fmt"

	"gorm.io/gorm"
)

type (
	ISellableProductRepository interface {
		Find(id string) (sellableProduct *Models.SellableProduct, err error)
		Update(sellableProductID string, sellableProduct *Models.SellableProduct) error
		UpdateTrx(trx *gorm.DB, sellableProductID string, sellableProduct *Models.SellableProduct) error
	}

	SellableProductRepository struct {
		DB *gorm.DB
	}
)

func SellableProductRepositoryProvider(db *gorm.DB) *SellableProductRepository {
	return &SellableProductRepository{DB: db}
}

func (sellableProductRepository *SellableProductRepository) Find(id string) (sellableProduct *Models.SellableProduct, err error) {
	sellableProduct = &Models.SellableProduct{}
	if err = sellableProductRepository.DB.Where("id = ?", id).First(sellableProduct).Error; err != nil {
		return nil, fmt.Errorf("sellable product not found: %w", err)
	}

	return sellableProduct, nil
}

func (sellableProductRepository *SellableProductRepository) Update(sellableProductID string, sellableProduct *Models.SellableProduct) error {
	if err := sellableProductRepository.DB.Model(&Models.SellableProduct{}).
		Where("id = ?", sellableProductID).
		Updates(sellableProduct).Error; err != nil {
		return fmt.Errorf("sellable product %w not updated: %w", sellableProductID, err)
	}

	return nil
}

func (sellableProductRepository *SellableProductRepository) UpdateTrx(trx *gorm.DB, sellableProductID string, sellableProduct *Models.SellableProduct) error {
	if err := trx.Model(&Models.SellableProduct{}).
		Where("id = ?", sellableProductID).
		Updates(sellableProduct).Error; err != nil {
		return fmt.Errorf("sellable product %w not updated: %w", sellableProductID, err)
	}

	return nil
}
