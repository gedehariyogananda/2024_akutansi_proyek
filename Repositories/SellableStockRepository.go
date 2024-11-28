package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	ISellableStockRepository interface {
		FindBySellableStockNotExp(sellableStockID string) (sellableStock []*Models.SellableStock, err error)
		UpdateCurrent(trx *gorm.DB, sellableStockID string, qtyClient int) error
		SumCurrentQuantity(sellableStockID string) (total int, err error)
	}

	SellableStockRepository struct {
		DB *gorm.DB
	}
)

func SellableStockRepositoryProvider(db *gorm.DB) *SellableStockRepository {
	return &SellableStockRepository{DB: db}
}

func (r *SellableStockRepository) FindBySellableStockNotExp(sellableStockID string) (sellableStock []*Models.SellableStock, err error) {
	if err := r.DB.Where("sellable_product_id = ?", sellableStockID).
		Where("expired_date > now()").
		Order("created_at asc").
		Find(&sellableStock).Error; err != nil {
		return nil, err
	}

	return sellableStock, nil
}

func (r *SellableStockRepository) UpdateCurrent(trx *gorm.DB, sellableStockID string, qtyClient int) error {

	db := trx
	if db == nil {
		db = r.DB
	}

	if err := db.Model(&Models.SellableStock{}).
		Where("id = ?", sellableStockID).
		Update("current_quantity", gorm.Expr("current_quantity - ?", qtyClient)).Error; err != nil {
		return err
	}

	return nil
}

func (r *SellableStockRepository) SumCurrentQuantity(sellableStockID string) (total int, err error) {
	if err := r.DB.Model(&Models.SellableStock{}).
		Select("sum(current_quantity) as total").
		Where("sellable_product_id = ?", sellableStockID).
		Where("expired_date > now()").
		Group("sellable_product_id").
		Scan(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}
