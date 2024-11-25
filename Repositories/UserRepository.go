package Repositories

import (
	"fmt"

	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IUserRepository interface {
		Create(userClient *Models.User) (*Models.User, error)
		FindByID(userID string) (user *Models.User, err error)
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
		return nil, fmt.Errorf("account tidak ditemukan!")
	}

	return user, nil
}

func (h *UserRepository) FindByID(userID string) (user *Models.User, err error) {
	user = &Models.User{}

	if err := h.DB.
		Where("id = ?", userID).
		First(user).Error; err != nil {
		return nil, fmt.Errorf("account tidak ditemukan!")
	}

	return user, nil
}
