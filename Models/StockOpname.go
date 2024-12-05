package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type StockOpname struct {
	ID          string     `json:"id"`
	Title       string     `json:"title,omitempty"`
	CompanyID   string     `json:"company_id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	ChangerName *string    `json:"changer_name,omitempty"`
}

func (stockOpname *StockOpname) BeforeCreate(tx *gorm.DB) (err error) {
	if stockOpname.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}

		stockOpname.ID = uuid.String()
	}

	if stockOpname.CreatedAt == nil {
		now := time.Now()
		stockOpname.CreatedAt = &now
	}

	return
}
