package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Utils"

	"gorm.io/gorm"
)

type (
	IUnitRepository interface {
		Create(unit *Models.Unit) (*Models.Unit, error)
		Update(unit *Models.Unit, id string) (*Models.Unit, error)
		Delete(id string) error
		FindById(id string) (*Models.Unit, error)
		FindAll(company_id string, meta *Common.Meta, filter *Common.Filter) (*[]Models.Unit, int, error)
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
	unit := &Models.Unit{}

	if err := r.DB.Where("id = ?", id).First(unit).Error; err != nil {
		return nil, err
	}

	return unit, nil
}

func (r *UnitRepository) FindAll(companyID string, meta *Common.Meta, filter *Common.Filter) (*[]Models.Unit, int, error) {
	var units []Models.Unit
	var total int64

	err := r.DB.
		Scopes(Utils.Paginate(meta.Page, meta.PerPage), Helper.FilterCompanyID(companyID), Helper.FilterStatus(filter.Status), Helper.FilterSearch(*filter.Name)).Find(&units).Error

	if err != nil {
		return nil, 0, err
	}

	err = r.DB.Model(&Models.Unit{}).Scopes(Helper.FilterCompanyID(companyID), Helper.FilterStatus(filter.Status), Helper.FilterSearch(*filter.Name)).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	return &units, int(total), nil
}
