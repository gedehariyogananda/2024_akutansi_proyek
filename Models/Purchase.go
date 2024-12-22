package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	ID                  string    `json:"id"`
	TotalPurchaseAmount int       `json:"total_purchase_amount"`
	CompanyID           string    `json:"company_id"`
	Tax                 float32   `json:"tax"`
	Discount            float32   `json:"discount"`
	PaymentType         string    `json:"payment_type"`
	DueDate             time.Time `json:"due_date"`
	CreatedAt           time.Time `json:"created_at"`
}

func (p *Purchase) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}

	return
}
