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
		ImageCompany: userCompany.Company.ImageCompany,
	}
}
