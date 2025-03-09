package Response

import (
	"2024_akutansi_project/Models"
)

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
	RefundAt      *string               `json:"refund_at,omitempty"`
	InvoiceItems  *[]Models.InvoiceItem `json:"invoice_items,omitempty"`
}

type InvoiceItemResponse struct {
	SellableProductID string                  `json:"sellable_product_id"`
	CountSale         float64                 `json:"count_sale"`
	ProductName       string                  `json:"product_name"`
	SellableProduct   *Models.SellableProduct `json:"sellable_product,omitempty"`
	TotalRevenue      *float64                `json:"total_revenue,omitempty"`
}

type StatisticResponse struct {
}

type CoreInvoiceRes struct {
	ID            string       `json:"id"`
	CustomerName  string       `json:"customer_name,omitempty"`
	PhoneNumber   *string      `json:"phone_number,omitempty"`
	Note          *string      `json:"note,omitempty"`
	TaxID         *string      `json:"tax_id,omitempty"`
	PaymentMethod *string      `json:"payment_method,omitempty"`
	InvoiceNumber string       `json:"invoice_number"`
	CompanyID     *string      `json:"company_id,omitempty"`
	Status        *string      `json:"status,omitempty"`
	Tax           *float64     `json:"tax,omitempty"`
	SubTotal      float64      `json:"sub_total,omitempty"`
	MoneyReceived *float64     `json:"money_received,omitempty"`
	MoneyBack     *float64     `json:"money_back,omitempty"`
	Total         float64      `json:"total,omitempty"`
	CreatedAt     string       `json:"created_at,omitempty"`
	CountSale     int          `json:"count_sale,omitempty"`
	RefundAt      *string      `json:"refund_at,omitempty"`
	InvoiceItems  []InvItemRes `json:"invoice_items"`
}

type InvItemRes struct {
	SellableProductID string        `json:"sellable_product_id"`
	Quantity          int           `json:"quantity"`
	Name              string        `json:"name"`
	Price             float64       `json:"price"`
	Promo             *Models.Promo `json:"promo,omitempty"`
	ResultTotal       *float64      `json:"result_total,omitempty"`
	PromoAmount       *float64      `json:"promo_amount,omitempty"`
}
