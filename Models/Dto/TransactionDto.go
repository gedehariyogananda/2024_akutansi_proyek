package Dto

type CreateTransactionDto struct {
	Title          string  `json:"title" validate:"required"`
	Name           string  `json:"name" validate:"required"`
	AdditionalData *string `json:"additional_data,omitempty"`
	Date           string  `json:"date" validate:"required"`
	DueDate        *string `json:"due_date,omitempty"`
	Note           *string `json:"note,omitempty"`
	TransactionID  *string `json:"transaction_id,omitempty"`
	PaymentMethod  string  `json:"payment_method" validate:"required"`
	PaymentType    string  `json:"payment_type" validate:"required"`
	Amount         float64 `json:"amount" validate:"required"`
	CompanyID      string  `json:"company_id" validate:"required"`
}
