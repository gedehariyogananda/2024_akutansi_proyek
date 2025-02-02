package Services

import (
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Utils"
	"errors"
)

type (
	IUserService interface {
		UploadAvatar(userID string, avatar string) (err error)
		GetCurrentUser(userID string) (res *Response.UserResponse, err error)
		ChangePassword(userID string, dto Dto.ChangePasswordDto) (statusCode int, err error)
	}

	UserService struct {
		userRepository Repositories.IUserRepository
	}
)

func UserServiceProvider(userRepository Repositories.IUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (u *UserService) UploadAvatar(userID string, avatar string) (err error) {
	err = u.userRepository.UpdateAvatar(userID, avatar)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetCurrentUser(userID string) (res *Response.UserResponse, err error) {
	user, err := u.userRepository.FindByID(userID)

	if err != nil {
		return nil, err
	}

	res = Response.ToUserResponse(user)

	return res, nil
}

func (u *UserService) ChangePassword(userID string, dto Dto.ChangePasswordDto) (statusCode int, err error) {
	user, err := u.userRepository.FindByID(userID)

	if err != nil {
		return 404, err
	}

	if err := Utils.ComparePassword(user.Password, dto.OldPassword); err != nil {
		return 400, errors.New("password lama tidak sesuai")
	}

	hashedPassword, err := Utils.HashPassword(dto.NewPassword)

	if err != nil {
		return 500, err
	}

	err = u.userRepository.UpdatePassword(userID, hashedPassword)

	if err != nil {
		return 500, err
	}

	return 200, nil
}
