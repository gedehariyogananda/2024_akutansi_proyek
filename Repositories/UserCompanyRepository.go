package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"

	"gorm.io/gorm"
)

type (
	IUserCompanyRepository interface {
		InsertUserCompany(request *Dto.MakeUserCompanyRequest) (err error)
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
