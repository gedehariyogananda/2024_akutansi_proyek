package Response

import "2024_akutansi_project/Models"

type InvoiceResponse struct {
	ID            string                `json:"id"`
	CustomerName  string                `json:"customer_name,omitempty"`
	PhoneNumber   *string               `json:"phone_number,omitempty"`
	Note          *string               `json:"note,omitempty"`
	TaxID         *string               `json:"tax_id,omitempty"`
	PaymentMethod *string               `json:"payment_method,omitempty"`
	InvoiceNumber string                `json:"invoice_number,omitempty"`
	CompanyID     *string               `json:"company_id,omitempty"`
	Status        *string               `json:"status,omitempty"`
	Tax           *float64              `json:"tax,omitempty"`
	SubTotal      float64               `json:"sub_total,omitempty"`
	CreatedAt     string                `json:"created_at,omitempty"`
	CountSale     *int                  `json:"count_sale,omitempty"`
	InvoiceItems  *[]Models.InvoiceItem `json:"invoice_items,omitempty"`
}

type InvoiceItemResponse struct {
	SellableProductID string                  `json:"sellable_product_id"`
	CountSale         float64                 `json:"count_sale"`
	ProductName       string                  `json:"product_name"`
	SellableProduct   *Models.SellableProduct `json:"sellable_product,omitempty"`
}

type StatisticResponse struct {
}
