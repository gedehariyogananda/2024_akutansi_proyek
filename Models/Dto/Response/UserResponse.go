package Response

import "2024_akutansi_project/Models"

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}

type SendOtpResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

func ToSendOtpResponse(token, message string) *SendOtpResponse {
	return &SendOtpResponse{
		Token:   token,
		Message: message,
	}
}

func ToUserResponse(user *Models.User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Name:     user.Name,
		Avatar:   *user.Avatar,
	}
}
