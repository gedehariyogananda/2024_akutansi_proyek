package Dto

type CreateMaterialConversionDto struct {
	Quantity int    `json:"quantity" binding:"required"`
	UniID    string `json:"unit_id" binding:"required"`
}

type UpdateMaterialConversionDto struct {
	Quantity int    `json:"quantity"`
	UniID    string `json:"unit_id"`
}
