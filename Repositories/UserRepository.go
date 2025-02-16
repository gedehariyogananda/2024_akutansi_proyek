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
		UpdateAvatar(userID string, fileName string) (err error)
		UpdatePassword(userID string, newPassword string) (err error)
		UpdateStatus(userID string, status bool) (err error)
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

func (h *UserRepository) UpdateAvatar(userID string, fileName string) (err error) {
	if err := h.DB.Model(&Models.User{}).
		Where("id = ?", userID).
		Update("avatar", fileName).Error; err != nil {
		return fmt.Errorf("error saat update avatar: %w", err)
	}

	return nil
}

func (h *UserRepository) UpdatePassword(userID string, newPassword string) (err error) {
	if err := h.DB.Model(&Models.User{}).
		Where("id = ?", userID).
		Update("password", newPassword).Error; err != nil {
		return fmt.Errorf("error saat update password: %w", err)
	}

	return nil
}

func (h *UserRepository) UpdateStatus(userID string, status bool) (err error) {
	if err := h.DB.Model(&Models.User{}).
		Where("id = ?", userID).
		Update("is_active", status).Error; err != nil {
		return fmt.Errorf("error saat update status: %w", err)
	}

	return nil
}
