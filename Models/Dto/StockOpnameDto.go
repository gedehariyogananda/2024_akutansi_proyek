package Dto

type CreateStockOpnameDto struct {
	Title     string `json:"judul" binding:"required"`
	Amount    int    `json:"jumlah" binding:"required"`
	CompanyID string `json:"-"`
}

type UpdateStockOpnameDto struct {
	ID   string `json:"id" binding:"required"`
	User string `json:"editor" binding:"required"`
}
