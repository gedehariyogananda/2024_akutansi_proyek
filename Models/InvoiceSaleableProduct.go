package Models

type InvoiceSaleableProduct struct {
	InvoiceID         string          `json:"invoice_id"`
	SaleableProductID string          `json:"saleable_product_id"`
	QuantitySold      int             `json:"quantity_sold"`
	CompanyID         string          `json:"-"`
	Invoice           Invoice         `gorm:"foreignKey:InvoiceID" json:"-"`
	SaleableProduct   SaleableProduct `gorm:"foreignKey:SaleableProductID" json:"saleable_product"`
	Company           Company         `gorm:"foreignKey:CompanyID" json:"-"`
}
