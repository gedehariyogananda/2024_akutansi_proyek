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
	IUnitRepository interface {
		Create(unit *Models.Unit) (*Models.Unit, error)
		Update(unit *Models.Unit, id string) (*Models.Unit, error)
		Delete(id string) error
		FindById(id string) (*Models.Unit, error)
		FindAll(company_id string, query *Common.Query) ([]*Models.Unit, int, error)
	}

	UnitRepository struct {
		DB *gorm.DB
	}
)

func UnitProvider(db *gorm.DB) *UnitRepository {
	return &UnitRepository{DB: db}
}

func (r *UnitRepository) Create(unit *Models.Unit) (*Models.Unit, error) {
	if err := r.DB.Create(unit).Error; err != nil {
		return nil, err
	}

	return unit, nil
}

func (r *UnitRepository) Update(unit *Models.Unit, id string) (*Models.Unit, error) {
	if err := r.DB.Where("id = ?", id).Updates(unit).Error; err != nil {
		return nil, err
	}

	return unit, nil
}

func (r *UnitRepository) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&Models.Unit{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UnitRepository) FindById(id string) (*Models.Unit, error) {
	var unit Models.Unit

	if err := r.DB.Where("id = ?", id).First(&unit).Error; err != nil {
		return nil, err
	}

	return &unit, nil
}

func (r *UnitRepository) FindAll(companyID string, query *Common.Query) ([]*Models.Unit, int, error) {
	var units []*Models.Unit
	var total int64

	fmt.Println("query", query.Status)

	err := r.DB.
		Scopes(Utils.Paginate(query.Page, query.Limit), Helper.FilterCompanyID(companyID), Helper.FilterStatus(query.Status), Helper.FilterSearch(*query.Search)).Find(&units).Error

	if err != nil {
		return nil, 0, err
	}

	err = r.DB.Model(&Models.Unit{}).Scopes(Helper.FilterCompanyID(companyID), Helper.FilterStatus(query.Status), Helper.FilterSearch(*query.Search)).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	return units, int(total), nil
}
