package Mapper

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto/Response"
)

func ToCompanyResponseDTO(userCompany Models.UserCompany) Response.CompanyResponseDTO {
	return Response.CompanyResponseDTO{
		ID:           userCompany.Company.ID,
		Name:         userCompany.Company.Name,
		Address:      userCompany.Company.Address,
		CodeCompany:  userCompany.Company.CodeCompany,
		ImageCompany: userCompany.Company.ImageCompany,
		CreatedAt:    userCompany.Company.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    userCompany.Company.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
