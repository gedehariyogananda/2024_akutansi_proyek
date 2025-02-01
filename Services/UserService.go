package Services

import (
	"2024_akutansi_project/Repositories"
)

type (
	IUserService interface {
		UploadAvatar(userID string, avatar string) (err error)
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
