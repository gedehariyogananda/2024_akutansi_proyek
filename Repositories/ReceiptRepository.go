package Repositories

import (
	"2024_akutansi_project/Models"
	"fmt"

	"gorm.io/gorm"
)

type (
	IReceiptRepository interface {
		FindAll(sellableProductID string) (receiptProducts []*Models.Receipt, err error)
		Create(receipt *Models.Receipt) (err error)
		DeleteByProductId(sellableProductID string) (err error)
		DeleteByMaterialId(materialProductID string) (err error)
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

func (r *ReceiptRepository) Create(receipt *Models.Receipt) (err error) {
	err = r.DB.Create(receipt).Error

	if err != nil {

	}
	return nil
}

func (r *ReceiptRepository) DeleteByProductId(sellableProductID string) (err error) {
	if err := r.DB.Where("sellable_product_id = ?", sellableProductID).Delete(&Models.Receipt{}).Error; err != nil {
		return fmt.Errorf("error when deleting receipt product: %w", err)
	}

	return nil
}

func (r *ReceiptRepository) DeleteByMaterialId(materialProductID string) (err error) {
	if err := r.DB.Where("material_product_id = ?", materialProductID).Delete(&Models.Receipt{}).Error; err != nil {
		return fmt.Errorf("error when deleting receipt product: %w", err)
	}

	return nil
}
