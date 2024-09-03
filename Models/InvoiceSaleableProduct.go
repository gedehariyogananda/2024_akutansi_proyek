package Models

type InvoiceSaleableProduct struct {
	InvoiceID         int
	SaleableProductID int
	QuantitySold      int
	CompanyID         int
	Invoice           Invoice         `gorm:"foreignKey:InvoiceID"`
	SaleableProduct   SaleableProduct `gorm:"foreignKey:SaleableProductID"`
	Company           Company         `gorm:"foreignKey:CompanyID"`
}
