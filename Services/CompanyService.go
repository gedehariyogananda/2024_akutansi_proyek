package Services

import (
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
)

type (
	ICompanyService interface {
		GetDetailCompany(companyID string) (res *Response.CompanyResponse, err error)
	}

	CompanyService struct {
		companyRepository Repositories.ICompanyRepository
		storageService    IStorageService
	}
)

func CompanyServiceProvider(companyRepository Repositories.ICompanyRepository, storageService IStorageService) *CompanyService {
	return &CompanyService{companyRepository: companyRepository, storageService: storageService}
}

func (c *CompanyService) GetDetailCompany(companyID string) (res *Response.CompanyResponse, err error) {
	company, err := c.companyRepository.FindByID(companyID)
	if err != nil {
		return nil, err
	}

	url, err := c.storageService.SignedUrl(Dto.StorageRequest{
		ObjectKey: *company.Image,
	})

	if err != nil {
		return nil, err
	}

	company.Image = &url

	res = Response.ToCompanyResponse(company)

	return res, nil
}
