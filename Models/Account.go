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

type Account struct {
	ID          string      `json:"id"`
	NameAccount string      `json:"name_account"`
	CompanyID   string      `json:"-"`
	TypeAccount TypeAccount `json:"type_account"`
	CodeAccount string      `json:"code_account"`
	DeletedAt   time.Time   `json:"deleted_at"`
}

// create uuid
func (account *Account) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if account.ID == "" {
		account.ID = uuid.New().String()
	}

	return
}
