package Common

type JournalEntryParams struct {
	AccountID       string
	SubTotal        float64
	Tax             float64
	CompanyID       string
	Note            string
	TransactionCode string
	AdditionalData  *map[string]interface{}
}
