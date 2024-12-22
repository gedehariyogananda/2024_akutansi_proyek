package Models

import "github.com/google/uuid"

type PurchaseMaterialProduct struct {
	ID                string  `json:"id"`
	PurchaseID        string  `json:"purchase_id"`
	MaterialProductID string  `json:"material_product_id"`
	UnitPrice         float64 `json:"unit_price"`
	Quantity          int     `json:"quantity"`
	CompanyID         string  `json:"company_id"`
}

func (p *PurchaseMaterialProduct) BeforeCreate() (err error) {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}
