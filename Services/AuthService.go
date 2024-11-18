package Services

import (
	"2024_akutansi_project/Config"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type (
	IAuthService interface {
		Register(request *Dto.RegisterRequest) (user *Models.User, err error, statusCode int)

		Login(request *Dto.LoginOwnerRequest) (token string, statusCode int, err error)
		TokenCompany(request *Dto.TokenCompanyRequest, user_id string) (token string, company *Models.Company, err error, statusCode int)
	}

	AuthService struct {
		authRepository    Repositories.IAuthRepository
		companyRepository Repositories.ICompanyRepository
		jwtService        IJwtService
	}
)

func AuthServiceProvider(authRepository Repositories.IAuthRepository, jwtService IJwtService, companyRepository Repositories.ICompanyRepository) *AuthService {
	return &AuthService{
		authRepository:    authRepository,
		companyRepository: companyRepository,
		jwtService:        jwtService,
	}
}

func (h *AuthService) Register(request *Dto.RegisterRequest) (user *Models.User, err error, statusCode int) {

	if err := h.authRepository.CheckUniqueField(request); err != nil {
		return nil, errors.New("email already exist"), http.StatusConflict
	}

	user, err = h.authRepository.InsertForRegister(request)
	if err != nil {
		return nil, errors.New("error insert user"), http.StatusInternalServerError
	}

	user, err = h.authRepository.GetUser(user.ID)

	if err != nil {
		return nil, errors.New("error get user"), http.StatusInternalServerError
	}

	return user, nil, http.StatusCreated
}

func (h *AuthService) Login(request *Dto.LoginOwnerRequest) (token string, statusCode int, err error) {
	userData, err := h.authRepository.FindEmail(request.Email)

	if err != nil {
		return "", http.StatusNotFound, errors.New("email not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(request.Password)); err != nil {
		return "", http.StatusUnauthorized, errors.New("password not match")
	}

	token, duration, err := h.jwtService.GenerateToken(userData.ID, request.Me)

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("error generate token")
	}

	err = Config.SetToRedis(userData.ID, token, duration)

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("error set redis")
	}

	// checkTokenRedis, err := Config.GetFromRedis(userData.ID)
	// if err != nil {
	// 	return "", http.StatusInternalServerError, errors.New("error get token from redis")
	// }

	// log.Println("log: data token in redis: ", checkTokenRedis)

	return token, http.StatusOK, err
}

func (h *AuthService) TokenCompany(request *Dto.TokenCompanyRequest, user_id string) (token string, company *Models.Company, err error, statusCode int) {
	token, err = h.jwtService.GenerateTokenWithCompany(user_id, request.CompanyID)
	if err != nil {
		return "", nil, errors.New("error generate token"), http.StatusInternalServerError
	}

	if err := h.authRepository.UpdateToken(token, user_id); err != nil {
		return "", nil, errors.New("error update token"), http.StatusInternalServerError
	}

	company, err = h.companyRepository.GetCompany(request.CompanyID)
	if err != nil {
		return "", nil, errors.New("error get company"), http.StatusNotFound
	}

	return token, company, nil, http.StatusOK
}
