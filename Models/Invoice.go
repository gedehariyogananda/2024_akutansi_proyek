package Models

import (
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
	ID            string  `json:"id"`
	CustomerName  string  `json:"customer_name"`
	PhoneNumber   *string `json:"phone_number"`
	Note          string  `json:"note"`
	TaxID         string  `json:"tax_id"`
	PaymentMethod string  `json:"payment_method"`
	InvoiceNumber string  `json:"invoice_number"`
	CompanyID     string  `json:"company_id"`
	Status        bool    `json:"status"`
	Tax           float64 `json:"tax"`
	SubTotal      float64 `json:"sub_total"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	DeletedAt     gorm.DeletedAt
}

// create uuid setup
func (invoice *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if invoice.ID == "" {
		invoice.ID = uuid.New().String()
	}

	return
}
