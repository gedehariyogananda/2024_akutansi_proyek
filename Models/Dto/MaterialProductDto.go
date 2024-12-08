package Dto

type CreateMaterialProductDto struct {
	Name                string                        `json:"name" binding:"required"`
	Sku                 string                        `json:"sku" binding:"required"`
	CategoryID          string                        `json:"category_id" binding:"required"`
	CompanyID           string                        `json:"_" `
	SmallestUnitID      string                        `json:"smallest_unit_id" binding:"required"`
	Status              bool                          `json:"status"`
	MaterialConversions []CreateMaterialConversionDto `json:"" binding:"required"`
}

type UpdateMaterialProductDto struct {
	Name                string                        `json:"name"`
	Sku                 string                        `json:"sku"`
	SmallestUnitID      string                        `json:"smallest_unit_id"`
	CategoryID          string                        `json:"category_id"`
	Status              bool                          `json:"status"`
	CompanyID           string                        `json:"-"`
	MaterialConversions []UpdateMaterialConversionDto `json:"material_conversions"`
}
