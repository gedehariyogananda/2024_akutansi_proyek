package Services

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"fmt"
	"time"
)

type (
	IDashboardService interface {
		GetSalesResume(companyId string) (res Response.SalesResumeResponse, err error)
		GetBestSellingProducts(companyId, startDate, endDate string, limit int) (res []Response.BestSellingResponse, err error)
		GetRevenue(companyId string, year int) (res []Response.MonthlyRevenueResponse, err error)
		GetExpense(companyId string, year int) (res []Response.MonthlyExpenseResponse, err error)
	}

	DashboardService struct {
		invoiceItemRepository Repositories.IInvoiceItemRepository
		invoiceRepository     Repositories.IInvoiceRepository
		purchaseRepository    Repositories.IPurchaseRepository
	}
)

func DashboardServiceProvider(invoiceItemRepository Repositories.IInvoiceItemRepository, invoiceRepository Repositories.IInvoiceRepository, purchaseRepository Repositories.IPurchaseRepository) *DashboardService {
	return &DashboardService{
		invoiceItemRepository: invoiceItemRepository,
		invoiceRepository:     invoiceRepository,
		purchaseRepository:    purchaseRepository,
	}
}

func (s *DashboardService) GetSalesResume(companyId string) (res Response.SalesResumeResponse, err error) {
	now := time.Now()
	formattedDateNow := now.Format("2006-01-02")

	yesterday := now.AddDate(0, 0, -1)
	formattedDateYesterday := yesterday.Format("2006-01-02")

	salesByDate, err := s.invoiceRepository.SumSalesByDate(companyId, formattedDateNow)

	if err != nil {
		return res, err
	}

	salesYesterday, err := s.invoiceRepository.SumSalesByDate(companyId, formattedDateYesterday)

	if err != nil {
		return res, err
	}

	comparisonTodayWithyesterday := Helper.CalculateComparisonYesterday(salesByDate, salesYesterday)

	month := now.Month()
	year := now.Year()

	salesByMonth, err := s.invoiceRepository.SumSalesByYearMonth(companyId, year, int(month))

	if err != nil {
		return res, err
	}

	salesLastMonth, err := s.invoiceRepository.SumSalesByYearMonth(companyId, year, int(month)-1)

	if err != nil {
		return res, err
	}

	comparisonThisMonthWithLastMonth := Helper.CalculateComparisonYesterday(salesByMonth, salesLastMonth)

	invoiceItem, err := s.invoiceItemRepository.GetMostProductSold(companyId, formattedDateNow)

	if err != nil {
		return res, err
	}

	res = Response.SalesResumeResponse{
		SalesThisMonth:                salesByMonth,
		SalesThisDay:                  salesByDate,
		BestSellingProduct:            invoiceItem.ProductName,
		LastDayComparison:             comparisonTodayWithyesterday,
		LastMonthComparison:           comparisonThisMonthWithLastMonth,
		NumberOfBestSellingPoductSold: invoiceItem.CountSale,
	}

	return res, nil
}

func (s *DashboardService) GetBestSellingProducts(companyId, startDate, endDate string, limit int) (res []Response.BestSellingResponse, err error) {
	if limit == 0 {
		limit = 5
	}

	fmt.Println("start date service", startDate)

	invoiceItems, err := s.invoiceItemRepository.GetBestSellingProducts(companyId, startDate, endDate, limit)

	if err != nil {
		return res, err
	}

	for _, invoiceItem := range invoiceItems {
		res = append(res, Response.BestSellingResponse{
			Name:          invoiceItem.ProductName,
			TotalQuantity: int(invoiceItem.CountSale),
			TotalRevenue:  *invoiceItem.TotalRevenue,
		})
	}

	return res, nil
}

func (s *DashboardService) GetRevenue(companyId string, year int) (res []Response.MonthlyRevenueResponse, err error) {
	res, err = s.invoiceRepository.GetMonthlyRevenue(companyId, year)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *DashboardService) GetExpense(companyId string, year int) (res []Response.MonthlyExpenseResponse, err error) {
	res, err = s.purchaseRepository.GetMonthlyExpense(companyId, year)

	if err != nil {
		return res, err
	}

	return res, nil
}
