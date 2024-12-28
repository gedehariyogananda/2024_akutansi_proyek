package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TypeTransactionClient string

const (
	TIPE_WITHDRAWAL         TypeTransactionClient = "TIPE_WITHDRAWAL"         // prive
	TIPE_CAPITAL_ADDITION   TypeTransactionClient = "TIPE_CAPITAL_ADDITION"   // penambahan modal
	TIPE_DEBT_PAYMENT       TypeTransactionClient = "TIPE_DEBT_PAYMENT"       // pembayaran hutang
	TIPE_RECEIVABLE_PAYMENT TypeTransactionClient = "TIPE_RECEIVABLE_PAYMENT" // pembayaran piutang
	TIPE_SALE               TypeTransactionClient = "TIPE_SALE"               // penjualan
	TIPE_PURCHASE           TypeTransactionClient = "TIPE_PURCHASE"           // pembelian
	TIPE_TAX_PAYMENT        TypeTransactionClient = "TIPE_TAX_PAYMENT"        // pembayaran pajak
	TYPE_EXPENSE            TypeTransactionClient = "TYPE_EXPENSE"            // bayar beban/tanggungan/jasa
	PAYMENT_METHOD_CASH     TypeTransactionClient = "PAYMENT_METHOD_CASH"     // cash
	PAYMENT_METHOD_DEBIT    TypeTransactionClient = "PAYMENT_METHOD_DEBIT"    // debit

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
