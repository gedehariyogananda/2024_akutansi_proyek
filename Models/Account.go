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
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Type      TypeAccount `json:"type"`
	CompanyID string      `json:"company_id"`
	Code      string      `json:"code"`
	IsLocked  bool        `json:"is_locked"`
	Status    bool        `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// create uuid
func (account *Account) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if account.ID == "" {
		account.ID = uuid.New().String()
	}

	return
}
