package Response

type SalesResumeResponse struct {
	SalesThisMonth                float64 `json:"sales_this_month"`
	SalesThisDay                  float64 `json:"sales_this_day"`
	LastMonthComparison           float64 `json:"last_month_comparison"`
	LastDayComparison             float64 `json:"last_day_comparison"`
	BestSellingProduct            string  `json:"best_selling_product"`
	NumberOfBestSellingPoductSold float64 `json:"number_of_best_selling_product_sold"`
}
