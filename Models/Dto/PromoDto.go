package Dto

type CreatePromoDto struct {
	Name               string   `json:"name" binding:"required"`
	Type               string   `json:"type" binding:"required"`
	StartDate          string   `json:"start_date" binding:"required,date"`
	EndDate            string   `json:"end_date" binding:"required,date"`
	IsAll              bool     `json:"is_all"`
	Amount             float64  `json:"amount" binding:"required"`
	SellableProductIDS []string `json:"sellable_product_ids"`
	CompanyID          string   `json:"-"`
}

type AsignPromoDto struct {
	SellableProductIDS []string `json:"sellable_product_ids" binding:"required"`
	PromoID            string   `json:"promo_id" binding:"required"`
}

type CreatePromoOnly struct {
	Name      string  `json:"name" binding:"required"`
	Type      string  `json:"type" binding:"required"`
	StartDate string  `json:"start_date" binding:"required,date"`
	EndDate   string  `json:"end_date" binding:"required,date"`
	Amount    float64 `json:"amount" binding:"required"`
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
