package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialProduct struct {
	ID                  string
	MaterialProductName string
	StockTypeID         string
	IsAvailableForSale  bool
	UnitPriceForSelling float64
	TotalStock          int
	CompanyID           string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Company             Company   `gorm:"foreignKey:CompanyID"`
	StockType           StockType `gorm:"foreignKey:StockTypeID"`
}

// create uuid setup
func (materialProduct *MaterialProduct) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if materialProduct.ID == "" {
		materialProduct.ID = uuid.New().String()
	}

	return
}
