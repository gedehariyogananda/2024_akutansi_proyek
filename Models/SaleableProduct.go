package Models

import "time"

type SaleableProduct struct {
	ID          int       `json:"id"`
	ProductName string    `json:"product_name"`
	UnitPrice   float64   `json:"unit_price"`
	CompanyID   int       `json:"company_id"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Company     Company   `gorm:"foreignKey:CompanyID"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
}
