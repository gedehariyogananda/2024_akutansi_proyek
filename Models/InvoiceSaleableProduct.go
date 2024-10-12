package Models

type InvoiceSaleableProduct struct {
	InvoiceID               string                `json:"invoice_id"`
	SaleableProductID       string                `json:"saleable_product_id"`
	QuantitySold            int                   `json:"quantity_sold"`
	CompanyID               string                `json:"-"`
	SaleableProductTopingID string                `json:"saleable_product_toping_id"`
	Notes                   string                `json:"notes"`
	Invoice                 Invoice               `gorm:"foreignKey:InvoiceID" json:"-"`
	SaleableProductToping   SaleableProductToping `gorm:"foreignKey:SaleableProductTopingID" json:"saleable_product_toping"`
	SaleableProduct         SaleableProduct       `gorm:"foreignKey:SaleableProductID" json:"saleable_product"`
	Company                 Company               `gorm:"foreignKey:CompanyID" json:"-"`
}
