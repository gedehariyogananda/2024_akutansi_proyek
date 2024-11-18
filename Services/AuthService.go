package Services

import (
	"2024_akutansi_project/Config"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Utils"
	"errors"
	"log"
	"net/http"
)

type (
	IAuthService interface {
		Register(request *Dto.RegisterRequest) (user *Models.User, statusCode int, err error)
		LoginOwner(request *Dto.LoginOwnerRequest) (token string, statusCode int, err error)
		LoginEmployee(request *Dto.LoginEmployeeRequest) (token string, statusCode int, err error)
	}

	AuthService struct {
		userRepository    Repositories.IUserRepository
		subUserRepository Repositories.ISubUserRepository
		companyRepository Repositories.ICompanyRepository
		jwtService        IJwtService
	}
)

func AuthServiceProvider(userRepository Repositories.IUserRepository, jwtService IJwtService, companyRepository Repositories.ICompanyRepository, subUser Repositories.ISubUserRepository) *AuthService {
	return &AuthService{
		userRepository:    userRepository,
		companyRepository: companyRepository,
		jwtService:        jwtService,
		subUserRepository: subUser,
	}
}

func (service *AuthService) Register(request *Dto.RegisterRequest) (user *Models.User, statusCode int, err error) {

	// if err := service.userRepository.CheckUniqueField(request); err != nil {
	// 	return nil, errors.New("account already exist"), http.StatusConflict
	// }

	company := &Models.Company{
		Code: Utils.GenerateCodeCompany(request.CompanyName),
		Name: request.CompanyName,
	}

	company, err = service.companyRepository.Create(company)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error insert company")
	}

	user = &Models.User{
		Username:  Utils.FormatUsernameClient(request.Name),
		Email:     request.Email,
		Phone:     request.Phone,
		Password:  request.Password,
		Name:      request.Name,
		CompanyID: company.ID,
	}

	user, err = service.userRepository.Create(user)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error insert user")
	}

	return user, http.StatusCreated, nil
}

func (service *AuthService) LoginOwner(request *Dto.LoginOwnerRequest) (token string, statusCode int, err error) {
	ownerData, err := service.userRepository.FindEmail(request.Email)

	if err != nil {
		return "", http.StatusNotFound, errors.New("email not found")
	}

	if err := Utils.ComparePassword(ownerData.Password, request.Password); err != nil {
		return "", http.StatusUnauthorized, errors.New("password not match")
	}

	token, duration, err := service.jwtService.GenerateToken(ownerData.ID, ownerData.CompanyID, false, request.Me)

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("error generate token")
	}

	err = Config.SetToRedis(ownerData.ID, token, duration)

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("error set redis")
	}

	// checkTokenRedis, err := Config.GetFromRedis(ownerData.ID)
	// if err != nil {
	// 	return "", http.StatusInternalServerError, errors.New("error get token from redis")
	// }

	// log.Println("log: data token in redis: ", checkTokenRedis)

	return token, http.StatusOK, err
}

func (service *AuthService) LoginEmployee(request *Dto.LoginEmployeeRequest) (token string, statusCode int, err error) {
	employeeData, err := service.subUserRepository.FindByEmployeeKey(request.EmployeeKey)

	if err != nil {
		return "", http.StatusNotFound, errors.New("employee key not found")
	}

	if err := Utils.ComparePassword(employeeData.Password, request.Password); err != nil {
		return "", http.StatusUnauthorized, errors.New("password not match")
	}

	token, duration, err := service.jwtService.GenerateToken(employeeData.ID, employeeData.CompanyID, true, request.Me)

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("error generate token")
	}

	err = Config.SetToRedis(employeeData.ID, token, duration)

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("error set redis")
	}

	parseToken, err := service.jwtService.ParseToken(token)
	if err != nil {
		return "", http.StatusInternalServerError, errors.New("error parse token")
	}

	log.Println("log: data token claims redis: ", parseToken)

	return token, http.StatusOK, err
}
