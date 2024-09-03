package Models

import "time"

type PaymentMethod struct {
	ID         int
	MethodName string
	CompanyID  int
	CreatedAt  time.Time
	Company    Company `gorm:"foreignKey:CompanyID"`
}
