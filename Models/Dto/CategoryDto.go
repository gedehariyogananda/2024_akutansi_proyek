package Dto

type CreateCategoryRequestDTO struct {
	CategoryName string `json:"category_name" binding:"required"`
}

type UpdateCategoryRequestDTO struct {
	CategoryName string `json:"category_name"`
}
