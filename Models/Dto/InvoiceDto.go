package Dto

// ----- client request --- //
type PurchasedItem struct {
	ID                int  `json:"id" binding:"required"`
	QuantitySold      int  `json:"quantity_sold" binding:"required"`
	IsSaleableProduct bool `json:"is_saleable_product"`
	TotalPrice        int  `json:"total_price" binding:"required"`
}

type InvoiceRequestClient struct {
	InvoiceCustomer string          `json:"invoice_customer" binding:"required"`
	PaymentMethodID int             `json:"payment_method_id" binding:"required"`
	Purchaseds      []PurchasedItem `json:"purchaseds" binding:"required,dive"`
}

type InvoiceRequestDTO struct {
	ID              int    `json:"id"` // generated by gorm
	InvoiceNumber   string `json:"invoice_number" binding:"required"`
	InvoiceCustomer string `json:"invoice_customer" binding:"required"`
	InvoiceDate     string `json:"invoice_date" binding:"required"`
	TotalAmount     int    `json:"total_amount"`
	MoneyReceived   int    `json:"money_received"`
	StatusInvoice   string `json:"status_invoice"`
	CompanyID       int    `json:"company_id" binding:"required"`
	PaymentMethodId int    `json:"payment_method_id"`
}

type InvoiceMaterialRequestDTO struct {
	InvoiceID         int `json:"invoice_id" binding:"required"`
	MaterialProductID int `json:"material_product_id" binding:"required"`
	QuantitySold      int `json:"quantity_sold" binding:"required"`
	CompanyID         int `json:"company_id" binding:"required"`
}

type InvoiceSaleableRequestDTO struct {
	InvoiceID         int `json:"invoice_id" binding:"required"`
	SaleableProductID int `json:"saleable_product_id" binding:"required"`
	QuantitySold      int `json:"quantity_sold" binding:"required"`
	CompanyID         int `json:"company_id" binding:"required"`
}

type InvoiceStatusRequestDTO struct {
	StatusInvoice string `json:"status_invoice" binding:"required"`
}

type InvoiceMoneyReceivedRequestDTO struct {
	MoneyReceived float64 `json:"money_received" binding:"required"`
}

type InvoiceUpdateRequestDTO struct {
	InvoiceNumber   string `json:"invoice_number"`
	InvoiceCustomer string `json:"invoice_customer"`
	InvoiceDate     string `json:"invoice_date"`
	TotalAmount     int    `json:"total_amount"`
	MoneyReceived   int    `json:"money_received"`
	StatusInvoice   string `json:"status_invoice"`
	PaymentMethodId int    `json:"payment_method_id"`
}
