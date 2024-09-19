package Response

import "2024_akutansi_project/Models"

type InvoiceResponse struct {
	ID            string               `json:"id"`
	CompanyID     string               `json:"company_id"`
	InvoiceNumber string               `json:"invoice_number"`
	StatusInvoice Models.StatusInvoice `json:"status_invoice"`
	MoneyReceived float64              `json:"money_received"`
	MoneyBack     float64              `json:"money_back"`
	TotalAmount   float64              `json:"total_amount"`
}

type InvoiceDetailResponse struct {
	ID           string  `json:"id"`
	ProductName  string  `json:"product_name"`
	UnitPrice    float64 `json:"unit_price"`
	CompanyID    string  `json:"company_id"`
	CategoryName string  `json:"category_name"`
	Qty          int     `json:"qty"`
}
