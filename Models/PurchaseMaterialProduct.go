package Models

type PurchaseMaterialProduct struct {
	PurchaseID        string
	MaterialProductID string
	UnitPrice         float64
	QuantityPurchased int
	CompanyID         string
	Purchase          Purchase        `gorm:"foreignKey:PurchaseID"`
	MaterialProduct   MaterialProduct `gorm:"foreignKey:MaterialProductID"`
	Company           Company         `gorm:"foreignKey:CompanyID"`
}
