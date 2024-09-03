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
	ID              int
	InvoiceName     string
	InvoiceNumber   string
	InvoiceDate     time.Time
	TotalAmount     float64
	MoneyReceived   float64
	StatusInvoice   StatusInvoice
	CompanyID       int
	PaymentMethodID int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Company         Company       `gorm:"foreignKey:CompanyID"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
}
