package Dto

import "mime/multipart"

type (
	StorageRequest struct {
		File      *multipart.FileHeader
		ObjectKey string `json:"object_key" validate:"required"`
	}
)
