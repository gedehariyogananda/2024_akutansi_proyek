package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"time"

	"gorm.io/gorm"
)

type (
	ICompanyRepository interface {
		InsertCompany(request *Dto.MakeCompanyRequest) (company *Models.Company, err error)
		Update(request *Dto.EditCompanyRequest, company_id int) (company *Models.Company, err error)
		Delete(company_id int) (err error)
		GetCompany(company_id int) (company *Models.Company, err error)
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
		CreatedAt:    time.Now(),
	}

	if err := h.DB.Create(company).Error; err != nil {
		return nil, err
	}

	return company, nil
}

func (h *CompanyRepository) Update(request *Dto.EditCompanyRequest, company_id int) (company *Models.Company, err error) {
	company = &Models.Company{
		Name:         request.Name,
		Address:      request.Address,
		ImageCompany: request.ImageCompany,
		UpdatedAt:    time.Now(),
	}

	if err := h.DB.Model(&Models.Company{}).Where("id = ?", company_id).Updates(company).Error; err != nil {
		return nil, err
	}

	return company, nil
}

func (h *CompanyRepository) Delete(company_id int) (err error) {
	userCompanyModel := &Models.UserCompany{}

	if err := h.DB.Where("id = ?", company_id).Delete(&Models.Company{}).Error; err != nil {
		return err
	}

	if err := h.DB.Where("company_id = ?", company_id).Delete(userCompanyModel).Error; err != nil {
		return err
	}

	return nil
}

func (h *CompanyRepository) GetCompany(company_id int) (company *Models.Company, err error) {
	company = &Models.Company{}

	if err := h.DB.Where("id = ?", company_id).First(company).Error; err != nil {
		return nil, err
	}

	return company, nil
}
