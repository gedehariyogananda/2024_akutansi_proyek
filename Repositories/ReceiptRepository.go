package Repositories

import (
	"2024_akutansi_project/Models"
	"fmt"

	"gorm.io/gorm"
)

type (
	IReceiptRepository interface {
		FindAll(sellableProductID string) (receiptProducts []*Models.Receipt, err error)
	}

	ReceiptRepository struct {
		DB *gorm.DB
	}
)

func ReceiptRepositoryProvider(db *gorm.DB) *ReceiptRepository {
	return &ReceiptRepository{DB: db}
}

func (r *ReceiptRepository) FindAll(sellableProductID string) (receiptProducts []*Models.Receipt, err error) {
	if err := r.DB.Where("sellable_product_id = ?", sellableProductID).Find(&receiptProducts).Error; err != nil {
		return nil, fmt.Errorf("receipt products not found: %w", err)
	}

	return receiptProducts, nil
}
