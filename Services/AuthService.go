package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Utils"
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type (
	IAuthService interface {
		Register(request *Dto.RegisterRequest) (user *Models.User, statusCode int, err error)
		LoginOwner(ctx context.Context, request *Dto.LoginOwnerRequest) (token string, statusCode int, err error)
		LoginEmployee(ctx context.Context, request *Dto.LoginEmployeeRequest) (token string, statusCode int, err error)
	}

	AuthService struct {
		userRepository    Repositories.IUserRepository
		subUserRepository Repositories.ISubUserRepository
		companyRepository Repositories.ICompanyRepository
		jwtService        IJwtService
		redisClient       *redis.Client
	}
)

func AuthServiceProvider(userRepository Repositories.IUserRepository, jwtService IJwtService, companyRepository Repositories.ICompanyRepository, subUser Repositories.ISubUserRepository, redisClient *redis.Client) *AuthService {
	return &AuthService{
		userRepository:    userRepository,
		companyRepository: companyRepository,
		jwtService:        jwtService,
		subUserRepository: subUser,
		redisClient:       redisClient,
	}
}

func (service *AuthService) Register(request *Dto.RegisterRequest) (user *Models.User, statusCode int, err error) {

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

func (service *AuthService) LoginOwner(ctx context.Context, request *Dto.LoginOwnerRequest) (token string, statusCode int, err error) {
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

	err = service.redisClient.Set(ctx, ownerData.ID, token, duration).Err()

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("error set redis")
	}

	return token, http.StatusOK, err
}

func (service *AuthService) LoginEmployee(ctx context.Context, request *Dto.LoginEmployeeRequest) (token string, statusCode int, err error) {
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

	err = service.redisClient.Set(ctx, employeeData.ID, token, duration).Err()

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
