package Models

import "time"

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	CompanyID    int    `json:"company_id"`
	CreatedAt    time.Time
	Company      Company `gorm:"foreignKey:CompanyID"`
}
