package Models

import "time"

type SaleableProduct struct {
	ID          int
	ProductName string
	UnitPrice   float64
	CompanyID   int
	CategoryID  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Company     Company  `gorm:"foreignKey:CompanyID"`
	Category    Category `gorm:"foreignKey:CategoryID"`
}
