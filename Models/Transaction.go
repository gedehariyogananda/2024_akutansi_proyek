package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID                    string                  `json:"id"`
	Title                 string                  `json:"title"`
	Name                  string                  `json:"name"`
	AdditionalData        *map[string]interface{} `json:"additional_data,omitempty" gorm:"type:jsonb"`
	Date                  time.Time               `json:"date"`
	DueDate               time.Time               `json:"due_date,omitempty"`
	Note                  *string                 `json:"note,omitempty"`
	TransactionRecordCode string                  `json:"transaction_record_code"`
	TransactionID         *string                 `json:"transaction_id,omitempty"`
	PaymentMethod         string                  `json:"payment_method"`
	PaymentType           string                  `json:"payment_type"`
	Amount                float64                 `json:"amount"`
	CompanyID             string                  `json:"company_id"`
	CreatedAt             *time.Time              `json:"created_at"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}

		t.ID = uuid.String()
	}

	if t.CreatedAt == nil {
		now := time.Now()
		t.CreatedAt = &now
	}

	return

}
