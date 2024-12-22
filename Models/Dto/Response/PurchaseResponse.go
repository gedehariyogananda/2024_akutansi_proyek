package Response

import "2024_akutansi_project/Models"

type DropDwonPurchase struct {
	Products  []productResponse  `json:"products"`
	Materials []materialResponse `json:"materials"`
}

type productResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func toProductResponseSlice(product []Models.SellableProduct) []productResponse {
	productResponses := []productResponse{}
	for _, product := range product {
		productResponses = append(productResponses, productResponse{
			ID:   product.ID,
			Name: product.Name,
		})
	}
	return productResponses
}

type materialResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func toMaterialResponseSlice(material []Models.MaterialProduct) []materialResponse {
	materialResponses := []materialResponse{}
	for _, material := range material {
		materialResponses = append(materialResponses, materialResponse{
			ID:   material.ID,
			Name: material.Name,
		})
	}
	return materialResponses
}

func ToDropDownPurchase(products []Models.SellableProduct, materials []Models.MaterialProduct) DropDwonPurchase {
	return DropDwonPurchase{
		Products:  toProductResponseSlice(products),
		Materials: toMaterialResponseSlice(materials),
	}
}
