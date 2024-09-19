package Response

type SaleableResponseDTO struct {
	ID                string  `json:"id"`
	ProductName       string  `json:"product_name"`
	UnitPrice         float64 `json:"unit_price"`
	CompanyID         string  `json:"company_id"`
	CategoryName      string  `json:"category_name"`
	IsSaleableProduct bool    `json:"is_saleable_product"`
}

type MaterialResponseDTO struct {
	ID                string  `json:"id"`
	ProductName       string  `json:"product_name"`
	UnitPrice         float64 `json:"unit_price"`
	CompanyID         string  `json:"company_id"`
	CategoryName      string  `json:"category_name"`
	IsSaleableProduct bool    `json:"is_saleable_product"`
}

type DetailSaleableResponseDTO struct {
	ID           string  `json:"id"`
	ProductName  string  `json:"product_name"`
	UnitPrice    float64 `json:"unit_price"`
	Qty          int     `json:"qty"`
	CategoryName string  `json:"category_name"`
	TotalPrice   float64 `json:"total_price"`
}
