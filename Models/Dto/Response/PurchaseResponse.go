package Response

import (
	"2024_akutansi_project/Models"
	"time"
)

type DropDownPurchase struct {
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

func ToDropDownPurchase(products []Models.SellableProduct, materials []Models.MaterialProduct) DropDownPurchase {
	return DropDownPurchase{
		Products:  toProductResponseSlice(products),
		Materials: toMaterialResponseSlice(materials),
	}
}

type PurchaseResponse struct {
	ID            string  `json:"id"`
	PurchaseCode  string  `json:"purchase_code"`
	Date          string  `json:"date"`
	Status        string  `json:"status"`
	RestOfTheBill float32 `json:"rest_of_the_bill"`
	Total         float32 `json:"total"`
	IsAbleDelete  bool    `json:"is_able_delete"`
}

type PurchasesWithStatisticResponse struct {
	Purchases []PurchaseResponse `json:"purchases"`
	Statistic StatisTicPurchase  `json:"statistic"`
}

type StatisTicPurchase struct {
	PurchaseMonth float32 `json:"purchase_month"`
	DueMonth      float32 `json:"due_month"`
	PaidMonth     float32 `json:"paid_month"`
}

func ToPurchaseResponse(purchase Models.Purchase) PurchaseResponse {
	var status string
	if purchase.Payment == "cash" {
		status = "Lunas"
	} else {
		status = "Belum Lunas"
	}

	var restOfTheBill float32

	if purchase.Payment == "hutang" {
		restOfTheBill = purchase.TotalPurchaseAmount
	} else {
		restOfTheBill = 0
	}

	var isAbleDelete bool

	oneDay := time.Hour * 24

	if time.Now().Sub(purchase.CreatedAt) < oneDay {
		isAbleDelete = true
	} else {
		isAbleDelete = false
	}

	return PurchaseResponse{
		ID:            purchase.ID,
		PurchaseCode:  purchase.PurchaseNumber,
		Date:          purchase.CreatedAt.Format("2 January 2006"),
		Status:        status,
		RestOfTheBill: restOfTheBill,
		Total:         purchase.TotalPurchaseAmount,
		IsAbleDelete:  isAbleDelete,
	}
}

func ToPurchaseResponseSlice(purchases []*Models.Purchase, statistic StatisTicPurchase) PurchasesWithStatisticResponse {
	purchaseResponses := []PurchaseResponse{}
	for _, purchase := range purchases {
		purchaseResponses = append(purchaseResponses, ToPurchaseResponse(*purchase))
	}
	return PurchasesWithStatisticResponse{
		Purchases: purchaseResponses,
		Statistic: statistic,
	}
}
