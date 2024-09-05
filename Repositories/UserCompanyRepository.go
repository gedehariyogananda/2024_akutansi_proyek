package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"

	"gorm.io/gorm"
)

type (
	IUserCompanyRepository interface {
		InsertUserCompany(request *Dto.MakeUserCompanyRequest) (err error)
		FindAll(user_id int) (userCompany *[]Models.UserCompany, err error)
		Bind(company_id int) (userCompany *Models.UserCompany, err error)
	}

	UserCompanyRepository struct {
		DB *gorm.DB
	}
)

func UserCompanyRepositoryProvider(db *gorm.DB) *UserCompanyRepository {
	return &UserCompanyRepository{DB: db}
}

func (r *UserCompanyRepository) InsertUserCompany(request *Dto.MakeUserCompanyRequest) (err error) {
	userCompany := Models.UserCompany{
		UserID:    request.UserId,
		CompanyID: request.CompanyId,
	}

	if err := r.DB.Create(&userCompany).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserCompanyRepository) FindAll(user_id int) (userCompany *[]Models.UserCompany, err error) {
	userCompany = &[]Models.UserCompany{}

	if err := r.DB.Where("user_id = ?", user_id).
		Preload("Company").
		Find(&userCompany).Error; err != nil {
		return nil, err
	}

	return userCompany, nil
}

func (r *UserCompanyRepository) Bind(company_id int) (userCompany *Models.UserCompany, err error) {
	userCompany = &Models.UserCompany{}

	if err := r.DB.Where("company_id = ?", company_id).
		First(&userCompany).Error; err != nil {
		return nil, err
	}

	return userCompany, nil
}
