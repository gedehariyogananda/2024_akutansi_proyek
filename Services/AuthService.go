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
		Register(request *Dto.RegisterRequest) (err error)
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

func (h *AuthService) Register(request *Dto.RegisterRequest) (err error) {

	if err := h.repo.CheckUniqueField(request); err != nil {
		return err
	}

	if err := h.repo.InsertForRegister(request); err != nil {
		return err
	}

	return nil

}

func (h *AuthService) Login(request *Dto.LoginRequest) (user *Models.User, token string, err error) {
	user, err = h.repo.FindEmail(request.Email)

	if err != nil {
		return nil, "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, "", errors.New("password not match")
	}

	token, err = h.jwtService.GenerateToken(user.ID)

	if err != nil {
		return nil, "", errors.New("error generate token")
	}

	return user, token, err

}
