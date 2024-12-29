package Dto

type CreateMaterialProductDto struct {
	Name           string `json:"name" validate:"required"`
	Sku            string `json:"sku" validate:"required"`
	CategoryID     string `json:"category_id" validate:"required"`
	CompanyID      string `json:"_" `
	SmallestUnitID string `json:"smallest_unit_id" validate:"required"`
	Status         bool   `json:"status"`
	// MaterialConversions []CreateMaterialConversionDto `json:"" validate:"required"`
}

type UpdateMaterialProductDto struct {
	Name           string `json:"name"`
	Sku            string `json:"sku"`
	SmallestUnitID string `json:"smallest_unit_id"`
	CategoryID     string `json:"category_id"`
	Status         bool   `json:"status"`
	CompanyID      string `json:"-"`
	// MaterialConversions []UpdateMaterialConversionDto `json:"material_conversions"`
}
