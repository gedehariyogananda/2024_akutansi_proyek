package Dto

import "mime/multipart"

type UploadAvatarDto struct {
	Avatar *multipart.FileHeader `form:"avatar" validate:"required"`
}
