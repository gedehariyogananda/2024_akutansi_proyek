package Models

import "time"

type MaterialProduct struct {
	ID                  int
	MaterialProductName string
	StockTypeID         int
	IsAvailableForSale  bool
	UnitPriceForSelling float64
	TotalStock          int
	CompanyID           int
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Company             Company   `gorm:"foreignKey:CompanyID"`
	StockType           StockType `gorm:"foreignKey:StockTypeID"`
}
