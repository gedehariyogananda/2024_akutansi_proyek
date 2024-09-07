package Services

import (
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

		Login(request *Dto.LoginRequest) (user *Models.User, token string, err error, statusCode int)
		TokenCompany(request *Dto.TokenCompanyRequest, user_id int) (token string, company *Models.Company, err error, statusCode int)
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

func (h *AuthService) Login(request *Dto.LoginRequest) (user *Models.User, token string, err error, statusCode int) {
	userInit, err := h.authRepository.FindEmail(request.Email)

	if err != nil {
		return nil, "", errors.New("email not found"), http.StatusNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInit.Password), []byte(request.Password)); err != nil {
		return nil, "", errors.New("password not match"), http.StatusUnauthorized
	}

	token, err = h.jwtService.GenerateToken(userInit.ID)

	if err != nil {
		return nil, "", errors.New("error generate token"), http.StatusInternalServerError
	}

	if err := h.authRepository.UpdateToken(token, userInit.ID); err != nil {
		return nil, "", errors.New("error update token"), http.StatusInternalServerError
	}

	user, err = h.authRepository.GetUser(userInit.ID)

	if err != nil {
		return nil, "", errors.New("error get user"), http.StatusInternalServerError
	}

	return user, token, err, http.StatusOK
}

func (h *AuthService) TokenCompany(request *Dto.TokenCompanyRequest, user_id int) (token string, company *Models.Company, err error, statusCode int) {
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
