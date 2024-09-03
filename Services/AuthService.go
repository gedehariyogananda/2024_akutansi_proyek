package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type (
	IAuthService interface {
		Register(request *Dto.RegisterRequest) (user *Models.User, err error)
		Login(request *Dto.LoginRequest) (user *Models.User, token string, err error)
	}

	AuthService struct {
		repo       Repositories.IAuthRepository
		jwtService IJwtService
	}
)

func AuthServiceProvider(repo Repositories.IAuthRepository, jwtService IJwtService) *AuthService {
	return &AuthService{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (h *AuthService) Register(request *Dto.RegisterRequest) (user *Models.User, err error) {

	if err := h.repo.CheckUniqueField(request); err != nil {
		return nil, errors.New("email already exist")
	}

	user, err = h.repo.InsertForRegister(request)
	if err != nil {
		return nil, errors.New("error insert user")
	}

	user, err = h.repo.GetUser(user.ID)

	if err != nil {
		return nil, errors.New("error get user")
	}

	return user, nil

}

func (h *AuthService) Login(request *Dto.LoginRequest) (user *Models.User, token string, err error) {
	userInit, err := h.repo.FindEmail(request.Email)

	if err != nil {
		return nil, "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInit.Password), []byte(request.Password)); err != nil {
		return nil, "", errors.New("password not match")
	}

	token, err = h.jwtService.GenerateToken(userInit.ID)

	if err != nil {
		return nil, "", errors.New("error generate token")
	}

	if err := h.repo.UpdateToken(token, userInit.ID); err != nil {
		return nil, "", errors.New("error update token")
	}

	user, err = h.repo.GetUser(userInit.ID)

	if err != nil {
		return nil, "", errors.New("error get user")
	}

	return user, token, err

}
