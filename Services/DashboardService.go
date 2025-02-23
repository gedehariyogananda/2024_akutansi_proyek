package Services

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"time"
)

type (
	IDashboardService interface {
		GetSalesResume(companyId string) (res Response.SalesResumeResponse, err error)
	}

	DashboardService struct {
		invoiceItemRepository Repositories.IInvoiceItemRepository
		invoiceRepository     Repositories.IInvoiceRepository
	}
)

func DashboardServiceProvider(invoiceItemRepository Repositories.IInvoiceItemRepository, invoiceRepository Repositories.IInvoiceRepository) *DashboardService {
	return &DashboardService{
		invoiceItemRepository: invoiceItemRepository,
		invoiceRepository:     invoiceRepository,
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
