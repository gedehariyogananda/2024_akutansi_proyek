package Models

import "time"

type Category struct {
	ID           int
	CategoryName string
	CompanyID    int
	CreatedAt    time.Time
	Company      Company `gorm:"foreignKey:CompanyID"`
}
