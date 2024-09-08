package Models

import "time"

type Category struct {
	ID           int       `json:"id"`
	CategoryName string    `json:"category_name"`
	CompanyID    int       `json:"-"`
	CreatedAt    time.Time `json:"-"`
	Company      Company   `gorm:"foreignKey:CompanyID" json:"-"`
}
