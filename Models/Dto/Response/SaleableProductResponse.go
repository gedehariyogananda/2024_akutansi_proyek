package Response

import (
	"2024_akutansi_project/Models"
	"time"

	"gorm.io/gorm"
)

type SellableResponse struct {
	ID              string                     `json:"id"`
	Name            string                     `json:"name"`
	CompanyID       string                     `json:"company_id"`
	SmallestUnitID  string                     `json:"smallest_unit_id"`
	CategoryID      string                     `json:"category_id"`
	Sku             string                     `json:"sku"`
	Image           string                     `json:"image"`
	Description     string                     `json:"description"`
	Status          bool                       `json:"status"`
	StatusDisplay   string                     `json:"status_display"`
	HasReceipt      bool                       `json:"has_receipt"`
	CurrentQuantity int                        `json:"current_quantity"`
	Price           float64                    `json:"price"`
	CreatedAt       time.Time                  `json:"created_at"`
	UpdatedAt       time.Time                  `json:"updated_at"`
	DeletedAt       gorm.DeletedAt             `json:"deleted_at,omitempty"`
	Unit            *Models.Unit               `json:"unit,omitempty" gorm:"foreignKey:SmallestUnitID"`
	Category        *Models.Category           `json:"category,omitempty"`
	Materials       []*MaterialProduckResponse `json:"materials,omitempty"`
	PromoItems      []*Models.PromoItem        `json:"promo_items,omitempty"`
}

func ToSellableResponse(sellableProduct *Models.SellableProduct) *SellableResponse {

	var materials []*MaterialProduckResponse

	if sellableProduct.Receipts != nil {
		for _, receipt := range sellableProduct.Receipts {
			materials = append(materials, ToMaterialProduckResponse(receipt.MaterialProduct))
		}
	}

	return &SellableResponse{
		ID:              sellableProduct.ID,
		Name:            sellableProduct.Name,
		Sku:             sellableProduct.Sku,
		CompanyID:       sellableProduct.CompanyID,
		SmallestUnitID:  sellableProduct.SmallestUnitID,
		CategoryID:      sellableProduct.CategoryID,
		Image:           sellableProduct.Image,
		Description:     sellableProduct.Description,
		Status:          *sellableProduct.Status,
		HasReceipt:      sellableProduct.HasReceipt,
		CurrentQuantity: sellableProduct.CurrentQuantity,
		Price:           sellableProduct.Price,
		CreatedAt:       *sellableProduct.CreatedAt,
		UpdatedAt:       *sellableProduct.UpdatedAt,
		DeletedAt:       sellableProduct.DeletedAt,
		Unit:            sellableProduct.Unit,
		Category:        sellableProduct.Category,
		PromoItems:      sellableProduct.PromoItems,
		StatusDisplay:   "Active",
		Materials:       materials,
	}

}

func ToSellableResponsSlice(sellableProducts []*Models.SellableProduct) []*SellableResponse {
	var res []*SellableResponse

	for _, sellableProduct := range sellableProducts {
		res = append(res, ToSellableResponse(sellableProduct))
	}

	return res
}
