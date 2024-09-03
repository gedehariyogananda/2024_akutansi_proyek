package Models

type PurchaseMaterialProduct struct {
	PurchaseID        int
	MaterialProductID int
	UnitPrice         float64
	QuantityPurchased int
	CompanyID         int
	Purchase          Purchase        `gorm:"foreignKey:PurchaseID"`
	MaterialProduct   MaterialProduct `gorm:"foreignKey:MaterialProductID"`
	Company           Company         `gorm:"foreignKey:CompanyID"`
}
