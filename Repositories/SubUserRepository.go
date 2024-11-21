package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	ISubUserRepository interface {
		FindByEmployeeKey(employeeKey string) (subUser *Models.SubUser, err error)
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
