package Repositories

import (
	"errors"

	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IUserRepository interface {
		Create(userClient *Models.User) (*Models.User, error)
		GetUser(user_id string) (user *Models.User, err error)
		FindEmail(email string) (user *Models.User, err error)
	}

	UserRepository struct {
		DB *gorm.DB
	}
)

func UserRepositoryProvider(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (h *UserRepository) Create(userClient *Models.User) (*Models.User, error) {
	if err := h.DB.Create(userClient).Error; err != nil {
		return nil, err
	}

	return userClient, nil
}

func (h *UserRepository) FindEmail(email string) (user *Models.User, err error) {
	user = &Models.User{}

	if err := h.DB.
		Where("email = ?", email).
		First(user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (h *UserRepository) GetUser(user_id string) (user *Models.User, err error) {
	user = &Models.User{}

	if err := h.DB.
		Where("id = ?", user_id).
		First(user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
