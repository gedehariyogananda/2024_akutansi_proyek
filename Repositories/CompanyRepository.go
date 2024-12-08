package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	ICompanyRepository interface {
		Create(companyClient *Models.Company) (*Models.Company, error)
	}

	CompanyRepository struct {
		DB *gorm.DB
	}
)

func CompanyRepositoryProvider(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{DB: db}
}

func (h *CompanyRepository) Create(companyClient *Models.Company) (*Models.Company, error) {
	if err := h.DB.Create(companyClient).Error; err != nil {
		return nil, err
	}

	return companyClient, nil
}
