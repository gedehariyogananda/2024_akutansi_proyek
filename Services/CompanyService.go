package Services

import (
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
)

type (
	ICompanyService interface {
		GetDetailCompany(companyID string) (res *Response.CompanyResponse, err error)
	}

	CompanyService struct {
		companyRepository Repositories.ICompanyRepository
	}
)

func CompanyServiceProvider(companyRepository Repositories.ICompanyRepository) *CompanyService {
	return &CompanyService{companyRepository: companyRepository}
}

func (c *CompanyService) GetDetailCompany(companyID string) (res *Response.CompanyResponse, err error) {
	company, err := c.companyRepository.FindByID(companyID)
	if err != nil {
		return nil, err
	}

	res = Response.ToCompanyResponse(company)

	return res, nil
}
