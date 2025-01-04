package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JournalType string

const (
	DEBIT  JournalType = "DEBIT"
	CREDIT JournalType = "CREDIT"
)

type JournalEntry struct {
	ID              string                  `json:"id"`
	Amount          float64                 `json:"amount,omitempty"`
	CompanyID       string                  `json:"company_id,omitempty"`
	AccountID       string                  `json:"account_id,omitempty"`
	Type            JournalType             `json:"type,omitempty"`
	Note            string                  `json:"note,omitempty"`
	Date            time.Time               `json:"date,omitempty"`
	TransactionCode string                  `json:"transaction_code,omitempty"`
	CreditAt        *float64                `json:"credit_at,omitempty"`
	DebitAt         *float64                `json:"debit_at,omitempty"`
	Account         *Account                `json:"accounts,omitempty"`
	CreatedAt       *time.Time              `json:"created_at,omitempty"`
	AdditionalData  *map[string]interface{} `json:"additional_data,omitempty" gorm:"type:jsonb"`
}

func (journalEntry *JournalEntry) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if journalEntry.ID == "" {
		uuidV7, err := uuid.NewV7()
		if err != nil {
			return err
		}

		journalEntry.ID = uuidV7.String()
	}

	if journalEntry.CreatedAt == nil {
		now := time.Now()
		journalEntry.CreatedAt = &now
	}

	return
}
