package Repositories

import (
	"2024_akutansi_project/Models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type (
	ISellableProductRepository interface {
		Find(id string) (sellableProduct *Models.SellableProduct, err error)
		UpdateCurrentQty(trx *gorm.DB, sellableProductID string, QtyClient int) error
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
	log.Printf("Finding sellable product with ID: %s", id)

	if err = sellableProductRepository.DB.Where("id = ?", id).First(sellableProduct).Error; err != nil {
		return nil, fmt.Errorf("sellable product not found: %w", err)
	}

	log.Printf("Found sellable product: %+v", sellableProduct)

	return sellableProduct, nil
}

func (sellableProductRepository *SellableProductRepository) UpdateCurrentQty(trx *gorm.DB, sellableProductID string, qtyClient int) error {
	if err := trx.Model(&Models.SellableProduct{}).
		Where("id = ?", sellableProductID).
		Update("current_quantity", gorm.Expr("current_quantity - ?", qtyClient)).Error; err != nil {
		return fmt.Errorf("error when updating stock: %w", err)
	}

	return nil
}
