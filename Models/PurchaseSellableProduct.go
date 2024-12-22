package Models

import "github.com/google/uuid"

type PurchaseSellableProduct struct {
	ID                string  `json:"id"`
	PurchaseID        string  `json:"purchase_id"`
	SellableProductID string  `json:"sellable_product_id"`
	UnitPrice         float64 `json:"unit_price"`
	Quantity          int     `json:"quantity"`
	CompanyID         string  `json:"company_id"`
}

func (p *PurchaseSellableProduct) BeforeCreate() (err error) {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}
