package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Models/Mapper"
	"2024_akutansi_project/Repositories"
	"fmt"
	"net/http"
)

type (
	ICompanyService interface {
		AddCompany(request *Dto.MakeCompanyRequest, userID int) (err error, statusCode int)
		GetAllCompanyUser(user_id int) (companyResponse *[]Response.CompanyResponseDTO, err error, statusCode int)
		UpdateCompany(request *Dto.EditCompanyRequest, company_id int, user_id int) (company *Models.Company, statusCode int, err error)

		DeleteCompany(company_id int, user_id int) (statusCode int, err error)
	}

	CompanyService struct {
		companyRepository       Repositories.ICompanyRepository
		userCompanyRepository   Repositories.IUserCompanyRepository
		paymentMethodRepository Repositories.IPaymentMethodRepository
	}
)

func CompanyServiceProvider(companyRepository Repositories.ICompanyRepository, userCompanyRepository Repositories.IUserCompanyRepository, paymentMethodRepository Repositories.IPaymentMethodRepository) *CompanyService {
	return &CompanyService{
		companyRepository:       companyRepository,
		userCompanyRepository:   userCompanyRepository,
		paymentMethodRepository: paymentMethodRepository,
	}
}

func (s *CompanyService) AddCompany(request *Dto.MakeCompanyRequest, userID int) (err error, statusCode int) {
	company, err := s.companyRepository.InsertCompany(request)
	if err != nil {
		return fmt.Errorf("error insert company: %w", err), http.StatusBadRequest
	}

	if err := s.userCompanyRepository.InsertUserCompany(&Dto.MakeUserCompanyRequest{
		UserId:    userID,
		CompanyId: company.ID,
	}); err != nil {
		return fmt.Errorf("error insert user company: %w", err), http.StatusBadRequest
	}

	if err := s.paymentMethodRepository.CreateDefaultPaymentMethod(company.ID); err != nil {
		return fmt.Errorf("error create payment method: %w", err), http.StatusBadRequest
	}

	return nil, http.StatusCreated
}

func (s *CompanyService) GetAllCompanyUser(user_id int) (companyResponse *[]Response.CompanyResponseDTO, err error, statusCode int) {
	userCompany, err := s.userCompanyRepository.FindAll(user_id)
	if err != nil {
		return nil, fmt.Errorf("error get all company user: %w", err), http.StatusBadRequest
	}

	// for mapping data
	companyResponses := []Response.CompanyResponseDTO{}
	for _, uc := range *userCompany {
		companyResponse := Mapper.ToCompanyResponseDTO(uc)
		companyResponses = append(companyResponses, companyResponse)
	}

	return &companyResponses, nil, http.StatusOK
}

func (s *CompanyService) UpdateCompany(request *Dto.EditCompanyRequest, company_id int, user_id int) (company *Models.Company, statusCode int, err error) {
	userCompany, _ := s.userCompanyRepository.Bind(company_id)

	if userCompany == nil {
		return nil, http.StatusNotFound, fmt.Errorf("company not found")
	}

	if userCompany.UserID != user_id {
		return nil, http.StatusForbidden, fmt.Errorf("you are not allowed to update this company")
	}

	company, err = s.companyRepository.Update(request, company_id)

	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("error update company: %w", err)
	}

	return company, http.StatusOK, nil
}

func (s *CompanyService) DeleteCompany(company_id int, user_id int) (statusCode int, err error) {
	userCompany, _ := s.userCompanyRepository.Bind(company_id)

	if userCompany == nil {
		return http.StatusNotFound, fmt.Errorf("company not found")
	}

	if userCompany.UserID != user_id {
		return http.StatusForbidden, fmt.Errorf("you are not allowed to delete this company")
	}

	if err := s.companyRepository.Delete(company_id); err != nil {
		return http.StatusBadRequest, fmt.Errorf("error delete company: %w", err)
	}

	return http.StatusOK, nil

}
