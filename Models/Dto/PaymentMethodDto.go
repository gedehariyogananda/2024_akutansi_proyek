package Dto

type CreatePaymentMethodRequestDTO struct {
	MethodName string `json:"method_name" binding:"required"`
}

type UpdatePaymentMethodRequestDTO struct {
	MethodName string `json:"method_name"`
}
