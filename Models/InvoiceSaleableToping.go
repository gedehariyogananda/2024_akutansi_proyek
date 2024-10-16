package Models

type InvoiceSaleableToping struct {
	InvoiceSaleableProductID string                 `json:"invoice_saleable_product_id"`
	TopingID                 string                 `json:"toping_id"`
	CompanyID                string                 `json:"company_id"`
	InvoiceSaleableProduct   InvoiceSaleableProduct `gorm:"foreignKey:InvoiceSaleableProductID" json:"-"`
	Toping                   Toping                 `gorm:"foreignKey:TopingID" json:"toping"`
	Company                  Company                `gorm:"foreignKey:CompanyID" json:"-"`
}
