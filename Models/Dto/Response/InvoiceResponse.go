package Response

import "2024_akutansi_project/Models"

type InvoiceResponse struct {
	ID            int                  `json:"id"`
	InvoiceNumber string               `json:"invoice_number"`
	StatusInvoice Models.StatusInvoice `json:"status_invoice"`
	MoneyReceived float64              `json:"money_received"`
	MoneyBack     float64              `json:"money_back"`
	TotalAmount   float64              `json:"total_amount"`
}
