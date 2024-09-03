package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"

	"gorm.io/gorm"
)

type (
	ICompanyRepository interface {
		InsertCompany(request *Dto.MakeCompanyRequest) (company *Models.Company, err error)
	}

	CompanyRepository struct {
		DB *gorm.DB
	}
)

func CompanyRepositoryProvider(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{DB: db}
}

func (h *CompanyRepository) InsertCompany(request *Dto.MakeCompanyRequest) (company *Models.Company, err error) {
	company = &Models.Company{
		Name:         request.Name,
		Address:      request.Address,
		ImageCompany: request.ImageCompany,
	}

	if err := h.DB.Create(company).Error; err != nil {
		return nil, err
	}

	return company, nil
}
