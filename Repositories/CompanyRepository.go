package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"

	"gorm.io/gorm"
)

type (
	ICompanyRepository interface {
		InsertCompany(request *Dto.MakeCompanyRequest) (company *Models.Company, err error)
		Update(request *Dto.EditCompanyRequest, company_id string) (company *Models.Company, err error)
		Delete(company_id string) (err error)
		GetCompany(company_id string) (company *Models.Company, err error)
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

func (h *CompanyRepository) Update(request *Dto.EditCompanyRequest, company_id string) (company *Models.Company, err error) {
	company = &Models.Company{
		Name:         request.Name,
		Address:      request.Address,
		ImageCompany: request.ImageCompany,
	}

	if err := h.DB.Model(company).Where("id = ?", company_id).Updates(company).Error; err != nil {
		return nil, err
	}

	return company, nil
}

func (h *CompanyRepository) Delete(company_id string) (err error) {
	userCompanyModel := &Models.UserCompany{}
	companyModel := &Models.Company{}

	if err := h.DB.Where("id = ?", company_id).Delete(companyModel).Error; err != nil {
		return err
	}

	// cascade delete safety case
	if err := h.DB.Where("company_id = ?", company_id).Delete(userCompanyModel).Error; err != nil {
		return err
	}

	return nil
}

func (h *CompanyRepository) GetCompany(company_id string) (company *Models.Company, err error) {
	company = &Models.Company{}

	if err := h.DB.Where("id = ?", company_id).First(company).Error; err != nil {
		return nil, err
	}

	return company, nil
}
