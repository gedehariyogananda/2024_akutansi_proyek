package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusInvoice string

const (
	WAITING StatusInvoice = "WAITING"
	PROCESS StatusInvoice = "PROCESS"
	DONE    StatusInvoice = "DONE"
	CANCEL  StatusInvoice = "CANCEL"
)

type Invoice struct {
	ID              string        `json:"id"`
	InvoiceCustomer string        `json:"invoice_customer"`
	InvoiceNumber   string        `json:"invoice_number"`
	InvoiceDate     string        `json:"invoice_date"`
	TotalAmount     float64       `json:"total_amount"`
	MoneyReceived   float64       `json:"money_received"`
	StatusInvoice   StatusInvoice `json:"status_invoice"`
	CompanyID       string        `json:"-"`
	PaymentMethodID string        `json:"-"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Company         Company       `gorm:"foreignKey:CompanyID" json:"company"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"payment_method"`
}

// create uuid setup
func (invoice *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if invoice.ID == "" {
		invoice.ID = uuid.New().String()
	}

	return
}
