package Services

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Utils"
	"errors"
	"strconv"
)

type (
	IUserService interface {
		UploadAvatar(userID string, avatar string) (err error)
		GetCurrentUser(userID string) (res *Response.UserResponse, err error)
		ChangePassword(dto Dto.ChangePasswordDto) (statusCode int, err error)
		SendOtp(userID string) (res *Response.SendOtpResponse, err error)
	}

	UserService struct {
		userRepository Repositories.IUserRepository
		mailService    IEmailService
		jwtService     IJwtService
	}
)

func UserServiceProvider(userRepository Repositories.IUserRepository, mailService IEmailService, jwtService IJwtService) *UserService {
	return &UserService{
		userRepository: userRepository,
		mailService:    mailService,
		jwtService:     jwtService,
	}
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

func (u *UserService) ChangePassword(dto Dto.ChangePasswordDto) (statusCode int, err error) {
	claims, err := u.jwtService.ParseTokenOtp(dto.Token)

	if err != nil {
		return 400, errors.New("token tidak valid")
	}

	userId := claims["id"].(string)
	otp := claims["otp"].(string)

	if dto.Otp != otp {
		return 400, errors.New("otp tidak sesuai")
	}

	user, err := u.userRepository.FindByID(userId)

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

	err = u.userRepository.UpdatePassword(userId, hashedPassword)

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (u *UserService) SendOtp(userID string) (res *Response.SendOtpResponse, err error) {
	user, err := u.userRepository.FindByID(userID)

	if err != nil {
		return
	}

	otp := strconv.Itoa(Helper.GenerateRandomNumber(6))

	err = u.mailService.Send(*Models.ToSendOTPMessage(user.Email, user.Name, otp))

	if err != nil {
		return
	}

	token, err := u.jwtService.GenerateTokenForOtp(userID, otp)

	if err != nil {
		return
	}

	res = Response.ToSendOtpResponse(token, "OTP berhasil dikirim, Periksa Email Anda")

	return res, nil
}
