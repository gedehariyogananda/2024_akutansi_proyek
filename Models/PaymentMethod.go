package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID         string    `json:"id"`
	MethodName string    `json:"method_name"`
	CompanyID  string    `json:"company_id"`
	CreatedAt  time.Time `json:"-"`
	Company    Company   `gorm:"foreignKey:CompanyID" json:"-"`
}

// create uuid setup
func (paymentMethod *PaymentMethod) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if paymentMethod.ID == "" {
		paymentMethod.ID = uuid.New().String()
	}

	return
}
