package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Utils"
	"time"

	"gorm.io/gorm"
)

type (
	IPurchaseRepository interface {
		Create(purchase *Models.Purchase) (*Models.Purchase, error)
		FindAll(companyID string, query *Common.Query) ([]*Models.Purchase, int, error)
		GetTotalPurchaseMonth(companyID string) (float32, error)
	}

	PurchaseRepository struct {
		DB *gorm.DB
	}
)

func PurchaseRepositoryProvider(db *gorm.DB) *PurchaseRepository {
	return &PurchaseRepository{DB: db}
}

func (r *PurchaseRepository) Create(purchase *Models.Purchase) (*Models.Purchase, error) {
	if err := r.DB.Create(purchase).Error; err != nil {
		return nil, err
	}

	return purchase, nil
}

func (r *PurchaseRepository) FindAll(companyID string, query *Common.Query) ([]*Models.Purchase, int, error) {
	var purchases []*Models.Purchase
	var total int64

	if err := r.DB.Where("company_id = ?", companyID).Scopes(Utils.Paginate(query.Page, query.Limit)).Find(&purchases).Error; err != nil {
		return nil, 0, err
	}
	r.DB.Model(&Models.Purchase{}).Where("company_id = ?", companyID).Count(&total)

	return purchases, int(total), nil
}

func (r *PurchaseRepository) GetTotalPurchaseMonth(companyID string) (float32, error) {
	now := time.Now()

	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	var total float32

	if err := r.DB.Model(&Models.Purchase{}).Where("company_id = ?", companyID).Where("created_at BETWEEN ? AND ?", startOfMonth, endOfMonth).
		Select("SUM(total_purchase_amount)").Scan(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}
