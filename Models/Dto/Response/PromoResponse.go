package Response

import (
	"2024_akutansi_project/Models"
	"time"
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
	Ammount   float64                `json:"ammount"`
	StartDate time.Time              `json:"start_date"`
	EndDate   time.Time              `json:"end_date"`
	IsAll     bool                   `json:"is_all"`
	Type      string                 `json:"type"`
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
			Ammount:   productPromo.Amount,
			StartDate: productPromo.StartDate,
			EndDate:   productPromo.EndDate,
			IsAll:     productPromo.IsAll,
			CompanyID: productPromo.CompanyID,
			Type:      productPromo.Type,
			Products:  toProductPromoResponseSlice(selabelProducts),
		}
	}

	return PromoResponse{
		ID:        productPromo.ID,
		Name:      productPromo.Name,
		Ammount:   productPromo.Amount,
		StartDate: productPromo.StartDate,
		EndDate:   productPromo.EndDate,
		IsAll:     productPromo.IsAll,
		Type:      productPromo.Type,
		CompanyID: productPromo.CompanyID,
	}

}

func ToPromoResponseSlice(promos []Models.Promo) []PromoResponse {
	promoResponses := []PromoResponse{}
	for _, promo := range promos {
		promoResponses = append(promoResponses, ToPromoResponse(promo))
	}
	return promoResponses
}
