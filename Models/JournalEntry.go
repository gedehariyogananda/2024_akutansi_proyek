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
	Amount          float64                 `json:"amount"`
	CompanyID       string                  `json:"company_id"`
	AccountID       string                  `json:"account_id"`
	Type            JournalType             `json:"type"`
	Note            string                  `json:"note"`
	Date            time.Time               `json:"date"`
	TransactionCode string                  `json:"transaction_code"`
	Account         *Account                `json:"accounts,omitempty"`
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

	return
}
