package Dto

type CreateStockOpnameDto struct {
	Title       string               `json:"title" binding:"required"`
	CompanyID   string               `json:"-"`
	Items       []StockOpnameItemDTO `json:"items" binding:"required"`
	ChangerName string               `json:"changer_name"`
}

type StockOpnameItemDTO struct {
	Quantity       int    `json:"quantity" binding:"required"`
	StockId        string `json:"stock_id" binding:"required"`
	SystemQuantity int    `json:"system_quantity" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Type           string `json:"type" binding:"required"`
}
