package Models

import "time"

type PaymentMethod struct {
	ID         int       `json:"id"`
	MethodName string    `json:"method_name"`
	CompanyID  int       `json:"company_id"`
	CreatedAt  time.Time `json:"-"`
	Company    Company   `gorm:"foreignKey:CompanyID" json:"-"`
}
