package Models

type InvoiceMaterialProduct struct {
	InvoiceID         string          `json:"invoice_id"`
	MaterialProductID string          `json:"material_product_id"`
	QuantitySold      int             `json:"quantity_sold"`
	CompanyID         string          `json:"company_id"`
	Invoice           Invoice         `gorm:"foreignKey:InvoiceID" json:"invoice"`
	MaterialProduct   MaterialProduct `gorm:"foreignKey:MaterialProductID" json:"material_product"`
	Company           Company         `gorm:"foreignKey:CompanyID" json:"company"`
}
