package Dto

type CreateCategory struct {
	Name      string `json:"name" binding:"required"`
	Code      string `json:"code" binding:"required"`
	Type      string `json:"type" binding:"required,oneof=Produk Bahan"`
	Status    bool   `json:"status"`
	CompanyID string `json:"-"`
}

type UpdateCategory struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	Type      string `json:"type"`
	Status    bool   `json:"status"`
	CompanyID string `json:"-"`
}
