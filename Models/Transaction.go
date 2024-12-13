package Models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID                    string     `json:"id"`
	Title                 string     `json:"title"`
	Name                  string     `json:"name"`
	AdditionalData        *string    `json:"additional_data,omitempty"`
	Date                  time.Time  `json:"date"`
	DueDate               *time.Time `json:"due_date,omitempty"`
	Note                  *string    `json:"note,omitempty"`
	TransactionRecordCode *string    `json:"transaction_record_code"`
	TransactionID         *string    `json:"transaction_id,omitempty"`
	PaymentMethod         string     `json:"payment_method"`
	PaymentType           string     `json:"payment_type"`
	Amount                float64    `json:"amount"`
}

func (t *Transaction) BeforeCreate(ctx *gin.Context, tx *gorm.DB) (err error) {
	if t.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}

		t.ID = uuid.String()
	}

	return

}
