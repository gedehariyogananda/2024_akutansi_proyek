package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Models/Mapper"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Utils"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type (
	ICompanyService interface {
		AddCompany(request *Dto.MakeCompanyRequest, userID string, fileName string) (company *Models.Company, err error, filePath string, statusCode int)
		GetAllCompanyUser(user_id string) (companyResponse *[]Response.CompanyResponseDTO, err error, statusCode int)
		UpdateCompany(request *Dto.EditCompanyRequest, company_id string, user_id string, fileName string) (company *Models.Company, statusCode int, filePath string, err error)

		DeleteCompany(company_id string, user_id string) (statusCode int, err error)
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

func (s *CompanyService) AddCompany(request *Dto.MakeCompanyRequest, userID string, fileName string) (company *Models.Company, err error, filePath string, statusCode int) {

	fileName = Utils.GenerateUniqueFileName(fileName)
	request.ImageCompany = "/company-file/" + fileName

	filePath = os.Getenv("UPLOAD_DIR") + "/company-file/" + fileName

	var codeCompany string

	if request.CodeCompany != "" {
		suffix := Utils.GenerateUniqueSuffix()
		codeCompany = fmt.Sprintf("%s-%s", request.CodeCompany, suffix)
	} else {
		codeCompany = Utils.GenerateCodeCompany(request.Name)
	}

	company, err = s.companyRepository.InsertCompany(request, codeCompany)
	if err != nil {
		return nil, fmt.Errorf("error insert company: %w", err), "", http.StatusBadRequest
	}

	companyUser := &Dto.MakeUserCompanyRequest{
		UserId:    userID,
		CompanyId: company.ID,
	}

	if err := s.userCompanyRepository.InsertUserCompany(companyUser); err != nil {
		return nil, fmt.Errorf("error insert user company: %w", err), "", http.StatusBadRequest
	}

	if err := s.paymentMethodRepository.CreateDefaultPaymentMethod(company.ID); err != nil {
		return nil, fmt.Errorf("error create default payment method: %w", err), "", http.StatusBadRequest
	}

	return company, nil, filePath, http.StatusOK
}

func (s *CompanyService) GetAllCompanyUser(user_id string) (companyResponse *[]Response.CompanyResponseDTO, err error, statusCode int) {
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

func (s *CompanyService) UpdateCompany(request *Dto.EditCompanyRequest, company_id string, user_id string, fileName string) (company *Models.Company, statusCode int, filePath string, err error) {
	userCompany, _ := s.userCompanyRepository.Bind(company_id)

	if userCompany == nil {
		return nil, http.StatusNotFound, "", fmt.Errorf("company not found")
	}

	if userCompany.UserID != user_id {
		return nil, http.StatusForbidden, "", fmt.Errorf("you are not allowed to update this company")
	}

	if fileName != "" {
		fileName = Utils.GenerateUniqueFileName(fileName)
		request.ImageCompany = "/company-file/" + fileName
		filePath = os.Getenv("UPLOAD_DIR") + "/company-file/" + fileName

		nameOldPath := userCompany.Company.ImageCompany

		if strings.HasPrefix(nameOldPath, "/") {
			nameOldPath = nameOldPath[1:]
		}

		oldImageCompanyPath := os.Getenv("UPLOAD_DIR") + nameOldPath

		if userCompany.Company.ImageCompany != "" {
			if err := os.Remove(oldImageCompanyPath); err != nil {
				return nil, http.StatusBadRequest, "", fmt.Errorf("error remove old image company: %w", err)
			}
		}
	} else {
		// save without image
		request.ImageCompany = userCompany.Company.ImageCompany
	}

	company, err = s.companyRepository.Update(request, company_id)

	if err != nil {
		return nil, http.StatusBadRequest, "", fmt.Errorf("error update company: %w", err)
	}

	if filePath != "" {
		return company, http.StatusOK, filePath, nil
	} else {
		return company, http.StatusOK, "", nil
	}
}

func (s *CompanyService) DeleteCompany(company_id string, user_id string) (statusCode int, err error) {
	userCompany, _ := s.userCompanyRepository.Bind(company_id)

	if userCompany == nil {
		return http.StatusNotFound, fmt.Errorf("company not found")
	}

	if userCompany.UserID != user_id {
		return http.StatusForbidden, fmt.Errorf("you are not allowed to delete this company")
	}

	if userCompany.Company.ImageCompany != "" {
		nameOldPath := userCompany.Company.ImageCompany

		if strings.HasPrefix(nameOldPath, "/") {
			nameOldPath = nameOldPath[1:]
		}

		oldImageCompanyPath := os.Getenv("UPLOAD_DIR") + nameOldPath

		if err := os.Remove(oldImageCompanyPath); err != nil {
			return http.StatusBadRequest, fmt.Errorf("error remove old image company: %w", err)
		}
	}

	if err := s.companyRepository.Delete(company_id); err != nil {
		return http.StatusBadRequest, fmt.Errorf("error delete company: %w", err)
	}

	return http.StatusOK, nil

}
