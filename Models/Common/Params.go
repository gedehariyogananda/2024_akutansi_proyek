package Common

import "time"

type JournalEntryParams struct {
	AccountID       string
	SubTotal        float64
	Tax             float64
	Date            *time.Time
	CompanyID       string
	Note            string
	CreditAt        *float64
	DebitAt         *float64
	TransactionCode string
	AdditionalData  *map[string]interface{}
}
