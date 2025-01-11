package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Utils"
	"fmt"

	"gorm.io/gorm"
)

type (
	IMaterialProductRepository interface {
		Create(materialProduct *Models.MaterialProduct) (*Models.MaterialProduct, error)
		FindByCompany(companyID string) ([]*Models.MaterialProduct, error)
		UpdateCurrent(trx *gorm.DB, materialProductId string, qtyClient int) error
		FindByID(id string) (*Models.MaterialProduct, error)
		Delete(id string) error
		Update(materialProduct *Models.MaterialProduct, id string) (*Models.MaterialProduct, error)
		FindAll(companyID string, query *Common.Query) ([]*Models.MaterialProduct, int64, error)
	}

	MaterialProductRepository struct {
		DB *gorm.DB
	}
)

func MaterialProductRepositoryProvider(db *gorm.DB) *MaterialProductRepository {
	return &MaterialProductRepository{DB: db}
}

func (r *MaterialProductRepository) Create(materialProduct *Models.MaterialProduct) (*Models.MaterialProduct, error) {
	if err := r.DB.Create(materialProduct).Error; err != nil {
		return nil, fmt.Errorf("error when creating material product: %w", err)
	}

	return materialProduct, nil
}

func (r *MaterialProductRepository) FindByCompany(companyID string) ([]*Models.MaterialProduct, error) {
	var materialProduct []*Models.MaterialProduct

	if err := r.DB.Where("company_id = ?", companyID).Preload("Unit").Find(&materialProduct).Error; err != nil {
		return nil, fmt.Errorf("material product not found: %w", err)
	}

	return materialProduct, nil
}

func (r *MaterialProductRepository) UpdateCurrent(trx *gorm.DB, materialProductId string, qtyClient int) error {

	db := trx
	if db == nil {
		db = r.DB
	}

	if err := db.Model(&Models.MaterialProduct{}).
		Where("id = ?", materialProductId).
		Update("current_quantity", gorm.Expr("current_quantity - ?", qtyClient)).Error; err != nil {
		return fmt.Errorf("error when updating stock: %w", err)
	}

	return nil
}

func (r *MaterialProductRepository) FindByID(id string) (*Models.MaterialProduct, error) {
	var materialProduct Models.MaterialProduct

	if err := r.DB.Where("id = ?", id).Preload("Category").Preload("Unit").Preload("MaterialConversions.Unit").First(&materialProduct).Error; err != nil {
		return nil, fmt.Errorf("material product not found: %w", err)
	}

	fmt.Println(materialProduct.MaterialConversions)
	return &materialProduct, nil
}

func (r *MaterialProductRepository) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&Models.MaterialProduct{}).Error; err != nil {
		return fmt.Errorf("error when deleting material product: %w", err)
	}

	return nil
}

func (r *MaterialProductRepository) Update(materialProduct *Models.MaterialProduct, id string) (*Models.MaterialProduct, error) {
	if err := r.DB.Where("id = ?", id).Updates(materialProduct).Error; err != nil {
		return nil, fmt.Errorf("error when updating material product: %w", err)
	}

	return materialProduct, nil
}

func (r *MaterialProductRepository) FindAll(companyID string, query *Common.Query) ([]*Models.MaterialProduct, int64, error) {
	var materialProducts []*Models.MaterialProduct
	var total int64

	fmt.Println("company_id", companyID)
	fmt.Println("smallest unit id", query.SmallestUnitID)

	err := r.DB.Preload("Unit").Preload("MaterialConversions.Unit").Scopes(Utils.Paginate(query.Page, query.Limit),
		Helper.FilterCompanyID(companyID),
		Helper.FilterStatus(query.Status),
		Helper.FilterSearch(*query.Search),
		Helper.FilterUnitID(query.SmallestUnitID),
	).Find(&materialProducts).Error

	if err != nil {
		return nil, 0, fmt.Errorf("error when finding all material product: %w", err)
	}

	err = r.DB.Model(&Models.MaterialProduct{}).Scopes(
		Helper.FilterCompanyID(companyID),
		Helper.FilterStatus(query.Status),
		Helper.FilterSearch(*query.Search),
		Helper.FilterUnitID(query.SmallestUnitID),
	).Count(&total).Error

	if err != nil {
		return nil, 0, fmt.Errorf("error when finding all material product: %w", err)
	}

	return materialProducts, total, nil
}
