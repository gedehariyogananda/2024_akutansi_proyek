package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Utils"

	"gorm.io/gorm"
)

type (
	ISubUserRepository interface {
		FindByEmployeeKey(employeeKey string) (subUser *Models.SubUser, err error)
		Create(subUser *Models.SubUser) (*Models.SubUser, error)
		Update(subUser *Models.SubUser, id string) (*Models.SubUser, error)
		Delete(id string) error
		FindByID(id string) (subUser *Models.SubUser, err error)
		FindAll(companyId string, dto *Common.Query) (subUsers []*Models.SubUser, totalData int64, err error)
	}

	SubUserRepository struct {
		DB *gorm.DB
	}
)

func SubUserRepositoryProvider(db *gorm.DB) *SubUserRepository {
	return &SubUserRepository{DB: db}
}

func (r *SubUserRepository) FindByEmployeeKey(employeeKey string) (subUser *Models.SubUser, err error) {
	subUser = &Models.SubUser{}

	if err := r.DB.Where("employee_key = ?", employeeKey).First(subUser).Error; err != nil {
		return nil, err
	}

	return subUser, nil
}

func (r *SubUserRepository) Create(subUser *Models.SubUser) (*Models.SubUser, error) {
	if err := r.DB.Create(subUser).Error; err != nil {
		return nil, err
	}

	return subUser, nil
}

func (r *SubUserRepository) Update(subUser *Models.SubUser, id string) (*Models.SubUser, error) {
	if err := r.DB.Model(subUser).Where("id = ?", id).Updates(subUser).Update("status", subUser.Status).Error; err != nil {
		return nil, err
	}

	return subUser, nil
}

func (r *SubUserRepository) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&Models.SubUser{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *SubUserRepository) FindByID(id string) (subUser *Models.SubUser, err error) {
	subUser = &Models.SubUser{}

	if err := r.DB.Where("id = ?", id).First(subUser).Error; err != nil {
		return nil, err
	}

	return subUser, nil
}

func (r *SubUserRepository) FindAll(companyId string, dto *Common.Query) (subUsers []*Models.SubUser, totalData int64, err error) {
	err = r.DB.Scopes(Utils.Paginate(dto.Page, dto.Limit), Helper.FilterCompanyID(companyId), Helper.FilterStatus(dto.Status), Helper.FilterSearch(*dto.Search)).Find(&subUsers).Error

	if err != nil {
		return nil, 0, err
	}

	err = r.DB.Model(&Models.SubUser{}).Scopes(Helper.FilterStatus(dto.Status), Helper.FilterSearch(*dto.Search), Helper.FilterCompanyID(companyId)).Count(&totalData).Error

	if err != nil {
		return nil, 0, err
	}

	return subUsers, totalData, nil
}
