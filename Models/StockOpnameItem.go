package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockOpnameItem struct {
	ID                 string       `json:"id"`
	StockOpnameID      string       `json:"stock_opname_id,omitempty"`
	Quantity           int          `json:"quantity,omitempty"`
	StockID            string       `json:"stock_id,omitempty"`
	ProductType        string       `json:"product_type,omitempty"`
	DifferenceQuantity int          `json:"difference_quantity,omitempty"`
	Name               string       `json:"name,omitempty"`
	StockOpname        *StockOpname `json:"stock_opname,omitempty"`
}

func (item *StockOpnameItem) BeforeCreate(tx *gorm.DB) (err error) {
	if item.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}

		item.ID = uuid.String()
	}

	return
}
