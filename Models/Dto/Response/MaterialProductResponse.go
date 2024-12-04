package Response

import "2024_akutansi_project/Models"

type MaterialProductResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	SmallestUnitID  string `json:"smallest_unit_id"`
	Sku             string `json:"sku"`
	CategoryID      string `json:"category_id"`
	CompanyID       string `json:"company_id"`
	CurrentQuantity int    `json:"current_quantity"`
}

func ToMaterialResponse(materialProduct *Models.MaterialProduct) *MaterialProductResponse {
	return &MaterialProductResponse{
		ID:              materialProduct.ID,
		Name:            materialProduct.Name,
		SmallestUnitID:  materialProduct.SmallestUnitID,
		CategoryID:      materialProduct.CategoryID,
		CompanyID:       materialProduct.CompanyID,
		CurrentQuantity: materialProduct.CurrentQuantity,
	}
}

func ToMaterialResponseSlice(materialProducts []*Models.MaterialProduct) []*MaterialProductResponse {
	materialProductResponses := []*MaterialProductResponse{}
	for _, materialProduct := range materialProducts {
		materialProductResponses = append(materialProductResponses, ToMaterialResponse(materialProduct))
	}
	return materialProductResponses
}
