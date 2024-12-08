package Response

import "2024_akutansi_project/Models"

type MaterialProduckResponse struct {
	ID                  string                       `json:"id"`
	Name                string                       `json:"name"`
	Sku                 string                       `json:"sku"`
	CategoryID          string                       `json:"category_id"`
	CompanyID           string                       `json:"company_id"`
	Status              bool                         `json:"status"`
	CurrentQuantity     int                          `json:"current_quantity"`
	SmallestUnitID      string                       `json:"smallest_unit_id"`
	Unit                Unit                         `json:"unit"`
	MaterialConversions []MaterialConversionResponse `json:"material_conversion"`
}

func ToMaterialProduckResponse(materialProduck Models.MaterialProduct) MaterialProduckResponse {
	return MaterialProduckResponse{
		ID:                  materialProduck.ID,
		Name:                materialProduck.Name,
		Sku:                 materialProduck.Sku,
		CategoryID:          materialProduck.CategoryID,
		CompanyID:           materialProduck.CompanyID,
		Status:              materialProduck.Status,
		CurrentQuantity:     materialProduck.CurrentQuantity,
		SmallestUnitID:      materialProduck.SmallestUnitID,
		Unit:                *ToUnit(&materialProduck.Unit),
		MaterialConversions: ToMaterialConversionResponseSlice(materialProduck.MaterialConversions),
	}
}

func ToMaterialProduckResponseSlice(materialProducks []Models.MaterialProduct) []MaterialProduckResponse {
	materialProduckResponses := []MaterialProduckResponse{}
	for _, materialProduck := range materialProducks {
		materialProduckResponses = append(materialProduckResponses, ToMaterialProduckResponse(materialProduck))
	}
	return materialProduckResponses
}
