package Repositories

import (
	"errors"

	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"gorm.io/gorm"
)

type (
	IAuthRepository interface {
		InsertForRegister(request *Dto.RegisterRequest) (err error)
		CheckUniqueField(request *Dto.RegisterRequest) (err error)
		FindEmail(email string) (user *Models.User, err error)
	}

	AuthRepository struct {
		DB *gorm.DB
	}
)

func AuthRepositoryProvider(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (h *AuthRepository) InsertForRegister(request *Dto.RegisterRequest) (err error) {
	user := &Models.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := h.DB.Create(&user); err != nil {
		return err.Error
	}

	return nil
}

func (h *AuthRepository) CheckUniqueField(request *Dto.RegisterRequest) (err error) {
	user := &Models.User{}

	if err = h.DB.Where("username = ? OR email = ? OR phone = ?", request.Username, request.Email, request.Phone).First(&user).Error; err == nil {
		return errors.New("account already exist")
	}

	return nil
}

func (h *AuthRepository) FindEmail(email string) (user *Models.User, err error) {
	userInit := &Models.User{}

	if err := h.DB.Where("email = ?", email).First(&userInit).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return userInit, nil
}
