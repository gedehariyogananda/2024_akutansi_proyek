package Response

import (
	"2024_akutansi_project/Models"
	"fmt"
)

type AvailableStockResponse struct {
	ID              string `json:"id"`
	CurrentQuantity int    `json:"quantity"`
	ExpiredDate     string `json:"expired_date"`
	Name            string `json:"name"`
	Type            string `json:"type"`
}

func ToSellStock(availableStock *Models.SellableStock) *AvailableStockResponse {
	fmt.Println(availableStock.SellableProduct)
	return &AvailableStockResponse{
		ID:              availableStock.ID,
		CurrentQuantity: availableStock.CurrentQuantity,
		ExpiredDate:     availableStock.ExpiredDate.Format("2006-01-02"),
		Name:            availableStock.SellableProduct.Name,
		Type:            "product",
	}
}

func ToSellStockSlice(availableStocks []*Models.SellableStock) []*AvailableStockResponse {
	var availableStockResponses []*AvailableStockResponse

	for _, availableStock := range availableStocks {
		availableStockResponses = append(availableStockResponses, ToSellStock(availableStock))
	}

	return availableStockResponses
}

func ToMaterialStock(availableStock *Models.MaterialStock) *AvailableStockResponse {
	return &AvailableStockResponse{
		ID:              availableStock.ID,
		CurrentQuantity: availableStock.CurrentQuantity,
		ExpiredDate:     availableStock.ExpiredDate.Format("2006-01-02"),
		Name:            availableStock.MaterialProduct.Name,
		Type:            "material",
	}
}

func ToMaterialStockSlice(availableStocks []*Models.MaterialStock) []*AvailableStockResponse {
	var availableStockResponses []*AvailableStockResponse

	for _, availableStock := range availableStocks {
		availableStockResponses = append(availableStockResponses, ToMaterialStock(availableStock))
	}

	return availableStockResponses
}
