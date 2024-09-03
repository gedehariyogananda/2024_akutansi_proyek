package Services

import (
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"fmt"
)

type (
	ICompanyService interface {
		AddCompany(request *Dto.MakeCompanyRequest, userID int) (err error)
	}

	CompanyService struct {
		companyRepository     Repositories.ICompanyRepository
		userCompanyRepository Repositories.IUserCompanyRepository
	}
)

func CompanyServiceProvider(companyRepository Repositories.ICompanyRepository, userCompanyRepository Repositories.IUserCompanyRepository) *CompanyService {
	return &CompanyService{
		companyRepository:     companyRepository,
		userCompanyRepository: userCompanyRepository,
	}
}

func (h *CompanyService) AddCompany(request *Dto.MakeCompanyRequest, userID int) (err error) {
	company, err := h.companyRepository.InsertCompany(request)
	if err != nil {
		return fmt.Errorf("error insert company: %w", err)
	}

	if err := h.userCompanyRepository.InsertUserCompany(&Dto.MakeUserCompanyRequest{
		UserId:    userID,
		CompanyId: company.ID,
	}); err != nil {
		return fmt.Errorf("error insert user company: %w", err)
	}

	return nil
}
