package Response

import (
	"2024_akutansi_project/Models"
)

type productPromoResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func toProductPromoResponse(product *Models.SellableProduct) productPromoResponse {
	return productPromoResponse{
		ID:   product.ID,
		Name: product.Name,
		Type: product.Category.Name,
	}
}

func toProductPromoResponseSlice(products []*Models.SellableProduct) []productPromoResponse {
	productPromoResponses := []productPromoResponse{}
	for _, product := range products {
		productPromoResponses = append(productPromoResponses, toProductPromoResponse(product))
	}
	return productPromoResponses
}

type PromoResponse struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Quantity  int                    `json:"quantity"`
	Ammount   float64                `json:"ammount"`
	StartDate string                 `json:"start_date"`
	EndDate   string                 `json:"end_date"`
	IsAll     bool                   `json:"is_all"`
	CompanyID string                 `json:"company_id"`
	Products  []productPromoResponse `json:"products"`
}

func ToPromoResponse(productPromo Models.Promo) PromoResponse {
	var selabelProducts []*Models.SellableProduct

	if productPromo.PromoItems != nil {

		for _, promoItem := range productPromo.PromoItems {
			selabelProducts = append(selabelProducts, promoItem.SellableProduct)
		}
		return PromoResponse{
			ID:        productPromo.ID,
			Name:      productPromo.Name,
			Quantity:  productPromo.Quantity,
			Ammount:   productPromo.Amount,
			StartDate: productPromo.StartDate,
			EndDate:   productPromo.EndDate,
			IsAll:     productPromo.IsAll,
			CompanyID: productPromo.CompanyID,
			Products:  toProductPromoResponseSlice(selabelProducts),
		}
	}

	return PromoResponse{
		ID:        productPromo.ID,
		Name:      productPromo.Name,
		Quantity:  productPromo.Quantity,
		Ammount:   productPromo.Amount,
		StartDate: productPromo.StartDate,
		EndDate:   productPromo.EndDate,
		IsAll:     productPromo.IsAll,
		CompanyID: productPromo.CompanyID,
	}

}
