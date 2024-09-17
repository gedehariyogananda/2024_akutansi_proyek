package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockType struct {
	ID        string
	TypeName  string
	CreatedAt time.Time
}

// create uuid setup
func (stockType *StockType) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if stockType.ID == "" {
		stockType.ID = uuid.New().String()
	}

	return
}
