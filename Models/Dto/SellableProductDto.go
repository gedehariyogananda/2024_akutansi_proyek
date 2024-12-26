package Dto

type SellableProductDTO struct {
	PromoID           *string `json:"promo_id"`
	Description       *string `json:"description"`
	Status            *bool   `json:"status"`
	CurrentQuantity   *int    `json:"current_quantity"`
	PrefixDeletePromo *bool   `json:"prefix_delete_promo" `
}

type ReceiptMaterialDto struct {
	MaterialID string  `json:"material_id" validate:"required"`
	Quantity   float64 `json:"quantity" validate:"required"`
}

type CreateSellableProductDTO struct {
	Name           string                `form:"name" validate:"required"`
	SmallestUnitID string                `form:"smallest_unit_id" validate:"required"`
	CategoryID     string                `form:"category_id" validate:"required"`
	Sku            string                `form:"sku" validate:"required"`
	Description    string                `form:"description" validate:"required"`
	Status         *bool                 `form:"status"`
	Image          string                `form:"-"`
	Price          float64               `form:"price" validate:"required"`
	CompanyID      string                `form:"-"`
	Materials      string                `form:"materials"`
	MaterialsObj   *[]ReceiptMaterialDto `form:"-"`
}

type AssignMaterialDtoJson struct {
	MaterialID string `json:"material_id" validate:"required"`
	Quantity   int    `json:"quantity" validate:"required"`
}
type AssignMaterialDtos struct {
	SellableProductID string                   `json:"product_id" validate:"required"`
	Materials         []*AssignMaterialDtoJson `json:"materials" validate:"required"`
}

type AssignMaterialDtoForm struct {
	MaterialID string `from:"material_id" validate:"required"`
	Quantity   int    `form:"quantity" validate:"required"`
}
type CreateSellableProductWithAssignMaterialDTO struct {
	Name           string                   `form:"name" validate:"required"`
	SmallestUnitID string                   `form:"smallest_unit_id" validate:"required"`
	CategoryID     string                   `form:"category_id" validate:"required"`
	Sku            string                   `form:"sku" validate:"required"`
	Description    string                   `form:"description" validate:"required"`
	Status         *bool                    `form:"status"`
	Image          string                   `form:"-"`
	Price          float64                  `form:"price" validate:"required"`
	CompanyID      string                   `form:"-"`
	Materials      []*AssignMaterialDtoForm `form:"materials" validate:"required"`
}

type UnAssignMaterialDto struct {
	MaterialIDS []string `json:"material_ids" validate:"required"`
}

type UpdateSellableProductDTO struct {
	Name           string  `form:"name"`
	SmallestUnitID string  `form:"smallest_unit_id"`
	CategoryID     string  `form:"category_id" `
	Sku            string  `form:"sku" `
	Description    string  `form:"description" `
	Status         *bool   `form:"status"`
	Image          string  `form:"-"`
	Price          float64 `form:"price"`
}
