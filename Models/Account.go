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
	AccountCash            = "Cash"             // kas
	AccountRevenue         = "Revenue"          // pendapatan
	AccountOutputTax       = "Output Tax"       // pajak luaran
	AccountInputTax        = "Input Tax"        // pajak masukan
	AccountProductMaterial = "Product Material" // bahan produk
	AccountBusinessDebt    = "Business Debt"    // Hutang Usaha
	AccountBusinessCapital = "Business Capital" // modal usaha
	AccountReceivables     = "Receivables"      // piutang usaha
	AccountAssets          = "Assets"           // aset
	AccountCompanyExpense  = "Company Expense"  // beban perusahaan
)

type Account struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Type      TypeAccount `json:"type"`
	CompanyID string      `json:"company_id"`
	Code      string      `json:"code"`
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
