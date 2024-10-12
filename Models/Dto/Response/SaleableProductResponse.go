package Response

type SaleableResponseDTO struct {
	ID                    string      `json:"id"`
	ProductName           string      `json:"product_name"`
	UnitPrice             float64     `json:"unit_price"`
	CategoryName          string      `json:"category_name"`
	SaleableProductToping []TopingDTO `json:"saleable_product_toping"`
}

type TopingDTO struct {
	ID          string  `json:"id"`
	TopingName  string  `json:"toping_name"`
	PriceToping float64 `json:"price_toping"`
}

type DetailSaleableResponseDTO struct {
	ID           string  `json:"id"`
	ProductName  string  `json:"product_name"`
	UnitPrice    float64 `json:"unit_price"`
	QuantitySold int     `json:"quantity_sold"`
	CategoryName string  `json:"category_name"`
	TotalPrice   float64 `json:"total_price"`
}
