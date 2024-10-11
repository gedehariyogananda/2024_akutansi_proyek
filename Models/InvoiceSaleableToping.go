package Models

type InvoiceSaleableToping struct {
	InvoiceSaleableProductID string                 `json:"invoice_saleable_product_id"`
	SaleableProductTopingID  string                 `json:"saleable_product_toping_id"`
	InvoiceSaleableProduct   InvoiceSaleableProduct `gorm:"foreignKey:InvoiceSaleableProductID" json:"invoice_saleable_product"`
	SaleableProductToping    SaleableProductToping  `gorm:"foreignKey:SaleableProductTopingID" json:"saleable_product_toping"`
}
