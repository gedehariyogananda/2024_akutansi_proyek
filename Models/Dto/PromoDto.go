package Dto

type CreatePromoDto struct {
	Name               string   `json:"name" validate:"required"`
	Type               string   `json:"type" validate:"required"`
	StartDate          string   `json:"start_date" validate:"required,date"`
	EndDate            string   `json:"end_date" validate:"required,date"`
	IsAll              bool     `json:"is_all"`
	Amount             float64  `json:"amount" validate:"required"`
	SellableProductIDS []string `json:"sellable_product_ids"`
	CompanyID          string   `json:"-"`
}

type AsignPromoDto struct {
	SellableProductIDS []string `json:"sellable_product_ids" validate:"required"`
	PromoID            string   `json:"promo_id" validate:"required"`
}

type CreatePromoOnly struct {
	Name      string  `json:"name" validate:"required"`
	Type      string  `json:"type" validate:"required"`
	StartDate string  `json:"start_date" validate:"required,date"`
	EndDate   string  `json:"end_date" validate:"required,date"`
	Amount    float64 `json:"amount" validate:"required"`
	CompanyID string  `json:"-"`
}

type UpdatePromoDto struct {
	Name               string   `json:"name"`
	Type               string   `json:"type"`
	StartDate          string   `json:"start_date"`
	EndDate            string   `json:"end_date"`
	IsAll              bool     `json:"is_all"`
	Amount             float64  `json:"amount"`
	SellableProductIDS []string `json:"sellable_product_ids"`
	CompanyID          string   `json:"-"`
}
