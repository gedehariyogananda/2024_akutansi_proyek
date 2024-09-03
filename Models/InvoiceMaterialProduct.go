package Models

type InvoiceMaterialProduct struct {
	InvoiceID         int
	MaterialProductID int
	QuantitySold      int
	CompanyID         int
	Invoice           Invoice         `gorm:"foreignKey:InvoiceID"`
	MaterialProduct   MaterialProduct `gorm:"foreignKey:MaterialProductID"`
	Company           Company         `gorm:"foreignKey:CompanyID"`
}
