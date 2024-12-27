package Services

import (
	"2024_akutansi_project/Consts"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
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
		LoginMobile(ctx context.Context, request *Dto.LoginMobileRequest) (token string, typeUser string, statusCode int, err error)
		GetProfile(id string) (profile *Response.Profile, statusCode int, err error)
	}

	AuthService struct {
		userRepository    Repositories.IUserRepository
		subUserRepository Repositories.ISubUserRepository
		companyRepository Repositories.ICompanyRepository
		jwtService        IJwtService
		redisClient       *redis.Client
		accountRepository Repositories.IAccountRepository
	}
)

func AuthServiceProvider(userRepository Repositories.IUserRepository, jwtService IJwtService, companyRepository Repositories.ICompanyRepository, subUser Repositories.ISubUserRepository, redisClient *redis.Client, accountRepository Repositories.IAccountRepository) *AuthService {
	return &AuthService{
		userRepository:    userRepository,
		companyRepository: companyRepository,
		jwtService:        jwtService,
		subUserRepository: subUser,
		redisClient:       redisClient,
		accountRepository: accountRepository,
	}
}

func (service *AuthService) Register(request *Dto.RegisterRequest) (user *Models.User, statusCode int, err error) {

	checkEmail, _ := service.userRepository.FindEmail(request.Email)
	if checkEmail != nil {
		return nil, http.StatusBadRequest, errors.New("email sudah terdaftar di sistem kami!")
	}

	company, err := service.companyRepository.Create(&Models.Company{
		Code: Utils.GenerateCodeCompany(request.CompanyName),
		Name: request.CompanyName,
	})

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("kesalahan saat membuat company")
	}

	user, err = service.userRepository.Create(&Models.User{
		Username:  Utils.FormatUsernameClient(request.Name),
		Email:     request.Email,
		Phone:     request.Phone,
		Password:  request.Password,
		Name:      request.Name,
		CompanyID: company.ID,
	})

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("kesalahan saat membuat account")
	}

	// create account for company
	if err := service.accountRepository.InsertDefaultAccounts(company.ID); err != nil {
		return nil, http.StatusInternalServerError, errors.New("kesalahan saat membuat account")
	}

	return user, http.StatusCreated, nil
}

func (service *AuthService) LoginOwner(ctx context.Context, request *Dto.LoginOwnerRequest) (token string, statusCode int, err error) {
	ownerData, err := service.userRepository.FindEmail(request.Email)

	if err != nil {
		return "", http.StatusNotFound, errors.New("email tidak ditemukan")
	}

	if err := Utils.ComparePassword(ownerData.Password, request.Password); err != nil {
		return "", http.StatusUnauthorized, errors.New("password salah!")
	}

	company, err := service.companyRepository.FindByID(ownerData.CompanyID)

	if err != nil {
		return "", http.StatusNotFound, errors.New("company tidak ditemukan")
	}

	codeCompany := Utils.SuffixDigitsToUpper(company.Name, Consts.DigitCompanyCode)

	token, duration, err := service.jwtService.GenerateToken(ownerData.ID, ownerData.CompanyID, ownerData.Name, codeCompany, false, true)

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("error generate token")
	}

	// delete token redis existing before
	existingToken, _ := service.redisClient.Get(ctx, ownerData.ID).Result()
	if existingToken != "" {
		service.redisClient.Del(ctx, ownerData.ID)
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
		return "", http.StatusNotFound, errors.New("karyawan tidak ditemukan!")
	}

	if err := Utils.ComparePassword(employeeData.Password, request.Password); err != nil {
		return "", http.StatusUnauthorized, errors.New("password salah!")
	}

	company, err := service.companyRepository.FindByID(employeeData.CompanyID)

	if err != nil {
		return "", http.StatusNotFound, errors.New("company tidak ditemukan")
	}

	codeCompany := Utils.SuffixDigitsToUpper(company.Name, Consts.DigitCompanyCode)

	token, duration, err := service.jwtService.GenerateToken(employeeData.ID, employeeData.CompanyID, employeeData.Name, codeCompany, true, false)

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("error generate token")
	}

	// delete token redis existing before
	existingToken, _ := service.redisClient.Get(ctx, employeeData.ID).Result()
	if existingToken != "" {
		service.redisClient.Del(ctx, employeeData.ID)
	}

	err = service.redisClient.Set(ctx, employeeData.ID, token, duration).Err()

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("error set redis")
	}

	return token, http.StatusOK, err
}

func (service *AuthService) LoginMobile(ctx context.Context, request *Dto.LoginMobileRequest) (token string, typeUser string, statusCode int, err error) {

	var ownerData *Models.User
	var employeeData *Models.SubUser

	ownerData, err = service.userRepository.FindEmail(request.Key)

	if ownerData == nil || err != nil {
		employeeData, err = service.subUserRepository.FindByEmployeeKey(request.Key)

		if employeeData == nil || err != nil {
			return "", "", http.StatusNotFound, errors.New("user tidak ditemukan")
		}

		if err := Utils.ComparePassword(employeeData.Password, request.Password); err != nil {
			return "", "", http.StatusUnauthorized, errors.New("password salah!")
		}

		company, err := service.companyRepository.FindByID(employeeData.CompanyID)

		if err != nil {
			return "", "", http.StatusNotFound, errors.New("company tidak ditemukan")
		}

		codeCompany := Utils.SuffixDigitsToUpper(company.Name, Consts.DigitCompanyCode)

		token, duration, err := service.jwtService.GenerateToken(employeeData.ID, employeeData.CompanyID, employeeData.Name, codeCompany, true, true)

		if err != nil {
			return "", "", http.StatusInternalServerError, errors.New("error generate token")
		}

		// delete token redis existing before
		existingToken, _ := service.redisClient.Get(ctx, employeeData.ID).Result()
		if existingToken != "" {
			service.redisClient.Del(ctx, employeeData.ID)
		}

		err = service.redisClient.Set(ctx, employeeData.ID, token, duration).Err()

		if err != nil {
			return "", "", http.StatusInternalServerError, errors.New("error set redis")
		}

		return token, "EMPLOYEE", http.StatusOK, err
	}

	if err := Utils.ComparePassword(ownerData.Password, request.Password); err != nil {
		return "", "", http.StatusUnauthorized, errors.New("password salah!")
	}

	company, err := service.companyRepository.FindByID(ownerData.CompanyID)

	if err != nil {
		return "", "", http.StatusNotFound, errors.New("company tidak ditemukan")
	}

	codeCompany := Utils.SuffixDigitsToUpper(company.Name, Consts.DigitCompanyCode)

	token, duration, err := service.jwtService.GenerateToken(ownerData.ID, ownerData.CompanyID, ownerData.Name, codeCompany, false, true)

	if err != nil {
		return "", "", http.StatusInternalServerError, errors.New("error generate token")
	}

	// delete token redis existing before
	existingToken, _ := service.redisClient.Get(ctx, ownerData.ID).Result()
	if existingToken != "" {
		log.Print("delete token exist", existingToken)
		service.redisClient.Del(ctx, ownerData.ID)
	}

	err = service.redisClient.Set(ctx, ownerData.ID, token, duration).Err()

	if err != nil {
		return "", "", http.StatusInternalServerError, errors.New("error set redis")
	}

	return token, "OWNER", http.StatusOK, err
}

func (service *AuthService) GetProfile(id string) (profile *Response.Profile, statusCode int, err error) {
	var companyID string

	ownerData, err := service.userRepository.FindByID(id)
	if err == nil {
		companyID = ownerData.CompanyID

		profile = &Response.Profile{
			ID:          ownerData.ID,
			Username:    ownerData.Username,
			Name:        ownerData.Name,
			Email:       ownerData.Email,
			Phone:       ownerData.Phone,
			CompanyID:   ownerData.CompanyID,
			CompanyName: "",
			EmployeeKey: nil,
			IsEmployee:  false,
		}
	}

	if ownerData == nil {
		employeeData, err := service.subUserRepository.FindByID(id)
		if err == nil {
			companyID = employeeData.CompanyID

			profile = &Response.Profile{
				ID:          employeeData.ID,
				Name:        employeeData.Name,
				CompanyID:   employeeData.CompanyID,
				CompanyName: "",
				EmployeeKey: &employeeData.EmployeeKey,
				IsEmployee:  true,
			}
		}
	}

	if profile == nil {
		return nil, http.StatusNotFound, errors.New("data tidak ditemukan di owner maupun employee")
	}

	company, err := service.companyRepository.FindByID(companyID)
	if err != nil {
		return nil, http.StatusNotFound, errors.New("company tidak ditemukan")
	}

	profile.CompanyName = company.Name

	return profile, http.StatusOK, nil
}
