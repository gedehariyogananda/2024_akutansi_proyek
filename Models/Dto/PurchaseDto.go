package Dto

type createPurchaseDto struct {
	ID        string  `json:"id" validate:"required"`
	Type      string  `json:"type" validate:"required,oneof=material product"`
	ExpDate   string  `json:"exp_date" validate:"required,date"`
	Quantity  int     `json:"quantity" validate:"required"`
	UnitPrice float64 `json:"unit_price" validate:"required"`
}

type CreatePurchasesDto struct {
	Purchases           []createPurchaseDto `json:"purchases" validate:"required,dive"`
	Tax                 float32             `json:"tax" validate:"required"`
	Discount            float32             `json:"discount" validate:"required"`
	IsDiscountPercent   bool                `json:"is_discount_percent" validate:"required"`
	PaymentType         string              `json:"payment_type" validate:"required,oneof=cash hutang"`
	DueDate             string              `json:"due_date,omitempty" validate:"omitempty,date"`
	CompanyID           string              `json:"-"`
	TotalPurchaseAmount int                 `json:"total_purchase_amount" validate:"required"`
}
