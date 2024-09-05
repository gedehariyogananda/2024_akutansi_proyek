package Services

import (
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Models/Mapper"
	"2024_akutansi_project/Repositories"
	"fmt"
)

type (
	ICompanyService interface {
		AddCompany(request *Dto.MakeCompanyRequest, userID int) (err error)
		GetAllCompanyUser(user_id int) (companyResponse *[]Response.CompanyResponseDTO, err error)
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

func (s *CompanyService) AddCompany(request *Dto.MakeCompanyRequest, userID int) (err error) {
	company, err := s.companyRepository.InsertCompany(request)
	if err != nil {
		return fmt.Errorf("error insert company: %w", err)
	}

	if err := s.userCompanyRepository.InsertUserCompany(&Dto.MakeUserCompanyRequest{
		UserId:    userID,
		CompanyId: company.ID,
	}); err != nil {
		return fmt.Errorf("error insert user company: %w", err)
	}

	return nil
}

func (s *CompanyService) GetAllCompanyUser(user_id int) (companyResponse *[]Response.CompanyResponseDTO, err error) {
	userCompany, err := s.userCompanyRepository.FindAll(user_id)
	if err != nil {
		return nil, err
	}

	// for mapping data
	companyResponses := []Response.CompanyResponseDTO{}
	for _, uc := range *userCompany {
		companyResponse := Mapper.ToCompanyResponseDTO(uc)
		companyResponses = append(companyResponses, companyResponse)
	}

	return &companyResponses, nil
}
