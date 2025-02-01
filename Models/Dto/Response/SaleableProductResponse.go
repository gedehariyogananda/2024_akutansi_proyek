package Response

import (
	"2024_akutansi_project/Models"
	"time"

	"gorm.io/gorm"
)

type SellableResponse struct {
	ID              string                     `json:"id"`
	Name            *string                    `json:"name,omitempty"`
	CompanyID       *string                    `json:"company_id,omitempty"`
	SmallestUnitID  *string                    `json:"smallest_unit_id,omitempty"`
	CategoryID      *string                    `json:"category_id,omitempty"`
	Sku             *string                    `json:"sku,omitempty"`
	Image           *string                    `json:"image,omitempty"`
	Description     *string                    `json:"description,omitempty"`
	Status          *bool                      `json:"status,omitempty"`
	StatusDisplay   *string                    `json:"status_display,omitempty"`
	HasReceipt      *bool                      `json:"has_receipt,omitempty"`
	CurrentQuantity *int                       `json:"current_quantity,omitempty"`
	Price           *float64                   `json:"price,omitempty"`
	CreatedAt       *time.Time                 `json:"created_at,omitempty"`
	UpdatedAt       *time.Time                 `json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt             `json:"deleted_at,omitempty"`
	PromoID         *string                    `json:"promo_id,omitempty"`
	Unit            *Models.Unit               `json:"unit,omitempty" gorm:"foreignKey:SmallestUnitID"`
	Category        *Models.Category           `json:"category,omitempty"`
	PromoItems      []*Models.PromoItem        `json:"promo_items,omitempty"`
	Materials       []*MaterialProduckResponse `json:"materials,omitempty"`
	Promo           *Models.Promo              `json:"promo,omitempty"`
}

func ToSellableResponse(sellableProduct *Models.SellableProduct) *SellableResponse {

	var materials []*MaterialProduckResponse

	if sellableProduct.Receipts != nil {
		for _, receipt := range sellableProduct.Receipts {
			materials = append(materials, ToMaterialProduckResponse(receipt.MaterialProduct))
		}
	}

	status := ""

	if *sellableProduct.Status {
		status = "Active"
	} else if !*sellableProduct.Status && sellableProduct.CurrentQuantity <= 0 {
		status = "Habis"
	} else {
		status = "Non-Aktif"
	}

	return &SellableResponse{
		ID:              sellableProduct.ID,
		Name:            &sellableProduct.Name,
		Sku:             &sellableProduct.Sku,
		CompanyID:       &sellableProduct.CompanyID,
		SmallestUnitID:  &sellableProduct.SmallestUnitID,
		CategoryID:      &sellableProduct.CategoryID,
		Image:           &sellableProduct.Image,
		Description:     &sellableProduct.Description,
		Status:          sellableProduct.Status,
		HasReceipt:      &sellableProduct.HasReceipt,
		CurrentQuantity: &sellableProduct.CurrentQuantity,
		Price:           &sellableProduct.Price,
		CreatedAt:       sellableProduct.CreatedAt,
		UpdatedAt:       sellableProduct.UpdatedAt,
		DeletedAt:       sellableProduct.DeletedAt,
		Unit:            sellableProduct.Unit,
		Category:        sellableProduct.Category,
		PromoItems:      sellableProduct.PromoItems,
		StatusDisplay:   &status,
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
