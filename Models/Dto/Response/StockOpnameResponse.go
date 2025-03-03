package Response

import (
	"2024_akutansi_project/Models"
	"time"
)

type StockOpnameResponse struct {
	ID          string                    `json:"id"`
	Title       string                    `json:"title"`
	ChangerName *string                   `json:"changer_name"`
	Items       []StockOpnameItemResponse `json:"items"`
}

type StockOpnameItemResponse struct {
	Quantity       int       `json:"quantity"`
	SystemQuantity int       `json:"system_quantity"`
	RealQuantity   int       `json:"real_quantity"`
	InputQuantity  int       `json:"input_quantity"`
	OutputQuantity int       `json:"output_quantity"`
	ExpiredDate    time.Time `json:"expired_date"`
	Name           string    `json:"name"`
}

func MapFromStockOpname(stockOpname Models.StockOpname) *StockOpnameResponse {
	var items []StockOpnameItemResponse

	for _, item := range stockOpname.Items {
		items = append(items, StockOpnameItemResponse{
			Quantity:       item.Quantity,
			SystemQuantity: item.Quantity,
			RealQuantity:   item.Quantity + item.DifferenceQuantity,
			InputQuantity:  item.InitialQuantity,
			OutputQuantity: item.InitialQuantity - item.Quantity,
			ExpiredDate:    item.ExpiredDate,
			Name:           item.Name,
		})
	}

	return &StockOpnameResponse{
		ID:          stockOpname.ID,
		Title:       stockOpname.Title,
		ChangerName: stockOpname.ChangerName,
		Items:       items,
	}
}

type StockOpnameAllResponse struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	ChangerName *string    `json:"changer_name"`
	Amount      int        `json:"amount"`
	CreatedAt   *time.Time `json:"created_at"`
}

func MapFromStockOpnameAll(stockOpname []*Models.StockOpname) []StockOpnameAllResponse {
	var stockOpnameAllResponse []StockOpnameAllResponse

	for _, item := range stockOpname {
		stockOpnameAllResponse = append(stockOpnameAllResponse, StockOpnameAllResponse{
			ID:          item.ID,
			Title:       item.Title,
			ChangerName: item.ChangerName,
			Amount:      len(item.Items),
			CreatedAt:   item.CreatedAt,
		})
	}

	return stockOpnameAllResponse
}
