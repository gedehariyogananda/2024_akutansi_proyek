package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TypeTransaction string

const (
	DEBIT  TypeTransaction = "DEBIT"
	CREDIT TypeTransaction = "CREDIT"
)

type TransactionRecord struct {
	ID        string          `json:"id"`
	Count     int             `json:"count"`
	AccountID string          `json:"account_id"`
	Type      TypeTransaction `json:"type"`
	CreatedAt string          `json:"created_at"`
	DeletedAt string          `json:"deleted_at"`
}

// create uuid
func (transactionRecord *TransactionRecord) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if transactionRecord.ID == "" {
		transactionRecord.ID = uuid.New().String()
	}

	return
}
