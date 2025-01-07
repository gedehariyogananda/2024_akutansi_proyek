package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TypeAccount string

const (
	ASSET     TypeAccount = "ASSET"
	LIABILITY TypeAccount = "LIABILITY"
	EQUITY    TypeAccount = "EQUITY"
	REVENUE   TypeAccount = "REVENUE"
	EXPENSE   TypeAccount = "EXPENSE"
)

const (
	AccountCash                = "CASH" // Kas
	AccountCashCode            = "1001"
	AccountRevenue             = "REVENUE" // Pendapatan
	AccountRevenueCode         = "4001"
	AccountOutputTax           = "OUTPUT_TAX" // Pajak Luaran
	AccountOutputTaxCode       = "3002"
	AccountInputTax            = "INPUT_TAX" // Pajak Masukan
	AccountInputTaxCode        = "1003"
	AccountProductMaterial     = "PRODUCT_MATERIAL" // Bahan Produk
	AccountProductMaterialCode = "1002"
	AccountBusinessDebt        = "BUSINESS_DEBT" // Hutang Usaha
	AccountBusinessDebtCode    = "3001"
	AccountBusinessCapital     = "BUSINESS_CAPITAL" // Modal Usaha
	AccountBusinessCapitalCode = "2001"
	AccountReceivables         = "RECEIVABLES" // Piutang Usaha
	AccountReceivablesCode     = "1004"
	AccountAssets              = "ASSETS" // Aset
	AccountAssetsCode          = "1000"
	AccountCompanyExpense      = "COMPANY_EXPENSE" // Beban Perusahaan
	AccountCompanyExpenseCode  = "5000"
	AccountWithdrawal          = "WITHDRAWAL" // prive
	AccountWithdrawalCode      = "2002"
)

type Account struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Type      TypeAccount `json:"type"`
	CompanyID string      `json:"company_id"`
	Code      string      `json:"code"`
	IsLock    bool        `json:"is_lock"`
	Status    bool        `json:"status"`
	CreatedAt *time.Time  `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// create uuid
func (account *Account) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if account.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}

		account.ID = uuid.String()
	}

	if account.CreatedAt == nil {
		now := time.Now()
		account.CreatedAt = &now
	}

	return
}
