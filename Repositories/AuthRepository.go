package Repositories

import (
	"errors"

	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"

	"gorm.io/gorm"
)

type (
	IAuthRepository interface {
		InsertForRegister(request *Dto.RegisterRequest) (user *Models.User, err error)
		GetUser(user_id int) (user *Models.User, err error)
		UpdateToken(token string, user_id int) (err error)
		CheckUniqueField(request *Dto.RegisterRequest) (err error)
		FindEmail(email string) (user *Models.User, err error)
		CheckToken(token string, user_id int) (err error)
	}

	AuthRepository struct {
		DB *gorm.DB
	}
)

func AuthRepositoryProvider(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (h *AuthRepository) InsertForRegister(request *Dto.RegisterRequest) (user *Models.User, err error) {
	user = &Models.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := h.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (h *AuthRepository) UpdateToken(token string, user_id int) (err error) {
	user := &Models.User{
		Token: token,
	}

	if err := h.DB.Model(&user).Where("id = ?", user_id).Update("token", token).Error; err != nil {
		return err
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
	user = &Models.User{}

	if err := h.DB.
		Where("email = ?", email).
		First(user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (h *AuthRepository) GetUser(user_id int) (user *Models.User, err error) {
	user = &Models.User{}

	if err := h.DB.
		Where("id = ?", user_id).
		First(user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (h *AuthRepository) CheckToken(token string, user_id int) (err error) {
	user := &Models.User{}

	if err := h.DB.Where("id = ?", user_id).Find(&user).Error; err != nil {
		return errors.New("user not found")
	}

	if user.Token != token {
		return errors.New("token not match")
	}

	return nil
}
