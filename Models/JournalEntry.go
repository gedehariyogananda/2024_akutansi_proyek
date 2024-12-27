package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JournalType string
type TypeTransactionClient string

const (
	DEBIT  JournalType = "DEBIT"
	CREDIT JournalType = "CREDIT"
)

const (
	TIPE_WITHDRAWAL         TypeTransactionClient = "TIPE_WITHDRAWAL"         // prive
	TIPE_CAPITAL_ADDITION   TypeTransactionClient = "TIPE_CAPITAL_ADDITION"   // penambahan modal
	TIPE_DEBT_PAYMENT       TypeTransactionClient = "TIPE_DEBT_PAYMENT"       // pembayaran hutang
	TIPE_RECEIVABLE_PAYMENT TypeTransactionClient = "TIPE_RECEIVABLE_PAYMENT" // pembayaran piutang
	TIPE_SALE               TypeTransactionClient = "TIPE_SALE"               // penjualan
	TIPE_PURCHASE           TypeTransactionClient = "TIPE_PURCHASE"           // pembelian
	TIPE_TAX_PAYMENT        TypeTransactionClient = "TIPE_TAX_PAYMENT"        // pembayaran pajak
	TYPE_EXPENSE            TypeTransactionClient = "TYPE_EXPENSE"            // bayar beban/tanggungan/jasa

	// PAYMENT METHOD
	PAYMENT_METHOD_CASH  TypeTransactionClient = "PAYMENT_METHOD_CASH"  // cash
	PAYMENT_METHOD_DEBIT TypeTransactionClient = "PAYMENT_METHOD_DEBIT" // debit

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
	CreatedAt       *time.Time              `json:"created_at"`
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
