package Models

import "time"

type StatusInvoice string

const (
	WAITING StatusInvoice = "WAITING"
	PROCESS StatusInvoice = "PROCESS"
	DONE    StatusInvoice = "DONE"
	CANCEL  StatusInvoice = "CANCEL"
)

type Invoice struct {
	ID              int           `gorm:"primaryKey;autoIncrement" json:"id"`
	InvoiceCustomer string        `json:"invoice_customer"`
	InvoiceNumber   string        `json:"invoice_number"`
	InvoiceDate     string        `json:"invoice_date"`
	TotalAmount     float64       `json:"total_amount"`
	MoneyReceived   float64       `json:"money_received"`
	StatusInvoice   StatusInvoice `json:"status_invoice"`
	CompanyID       int           `json:"-"`
	PaymentMethodID int           `json:"-"`
	CreatedAt       time.Time     `json:"-"`
	UpdatedAt       time.Time     `json:"-"`
	Company         Company       `gorm:"foreignKey:CompanyID" json:"-"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"payment_method"`
}
