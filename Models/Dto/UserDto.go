package Dto

import "mime/multipart"

type UploadAvatarDto struct {
	Avatar *multipart.FileHeader `form:"avatar" validate:"required"`
}

type ChangePasswordDto struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
