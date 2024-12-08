package Response

import "2024_akutansi_project/Models"

type MaterialConversionResponse struct {
	ID                string `json:"id"`
	MaterialProductID string `json:"material_product_id"`
	Quantity          int    `json:"quantity"`
	UnitID            string `json:"unit_id"`
	Unit              Unit   `json:"unit"`
}

func ToMaterialConversionResponse(materialConversion Models.MaterialConversion) MaterialConversionResponse {
	return MaterialConversionResponse{
		ID:                materialConversion.ID,
		MaterialProductID: materialConversion.MaterialProductID,
		Quantity:          materialConversion.Quantity,
		UnitID:            materialConversion.UnitID,
		Unit:              *ToUnit(&materialConversion.Unit),
	}
}

func ToMaterialConversionResponseSlice(materialConversions []Models.MaterialConversion) []MaterialConversionResponse {
	materialConversionResponses := []MaterialConversionResponse{}
	for _, materialConversion := range materialConversions {
		materialConversionResponses = append(materialConversionResponses, ToMaterialConversionResponse(materialConversion))
	}
	return materialConversionResponses
}
