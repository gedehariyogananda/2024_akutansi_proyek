package Models

type InvoiceMaterialProduct struct {
	InvoiceID         string
	MaterialProductID string
	QuantitySold      int
	CompanyID         string
	Invoice           Invoice         `gorm:"foreignKey:InvoiceID"`
	MaterialProduct   MaterialProduct `gorm:"foreignKey:MaterialProductID"`
	Company           Company         `gorm:"foreignKey:CompanyID"`
}
