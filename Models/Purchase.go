package Models

import "time"

type Purchase struct {
	ID                  int
	TotalPurchaseAmount float64
	CompanyID           int
	CreatedAt           time.Time
	Company             Company `gorm:"foreignKey:CompanyID"`
}
