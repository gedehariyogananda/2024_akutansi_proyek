package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	ID            string         `json:"id"`
	CustomerName  string         `json:"customer_name,omitempty"`
	PhoneNumber   *string        `json:"phone_number,omitempty"`
	Note          string         `json:"note,omitempty"`
	TaxID         string         `json:"tax_id,omitempty"`
	PaymentMethod string         `json:"payment_method,omitempty"`
	InvoiceNumber string         `json:"invoice_number,omitempty"`
	CompanyID     string         `json:"company_id,omitempty"`
	Status        bool           `json:"status,omitempty"`
	Tax           float64        `json:"tax,omitempty"`
	SubTotal      float64        `json:"sub_total,omitempty"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	UpdatedAt     time.Time      `json:"updated_at,omitempty"`
	RefundAt      *time.Time     `json:"refund_at,omitempty"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty"`
	InvoiceItems  []InvoiceItem  `json:"invoice_items,omitempty"`
}

// create uuid setup
func (invoice *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if invoice.ID == "" {
		invoice.ID = uuid.New().String()
	}

	return
}
