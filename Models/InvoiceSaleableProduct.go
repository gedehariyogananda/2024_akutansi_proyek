package Models

type InvoiceSaleableProduct struct {
	InvoiceID         string
	SaleableProductID string
	QuantitySold      int
	CompanyID         string
	Invoice           Invoice         `gorm:"foreignKey:InvoiceID"`
	SaleableProduct   SaleableProduct `gorm:"foreignKey:SaleableProductID"`
	Company           Company         `gorm:"foreignKey:CompanyID"`
}
