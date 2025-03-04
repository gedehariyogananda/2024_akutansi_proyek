package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	IDashboardController interface {
		GetSalesResume(ctx *gin.Context)
		GetBestSalesProduct(ctx *gin.Context)
		GetRevenue(ctx *gin.Context)
		GetExpense(ctx *gin.Context)
	}

	DashboardController struct {
		DashboardService Services.IDashboardService
	}
)

func DashboardControllerProvider(dashboardService Services.IDashboardService) *DashboardController {
	return &DashboardController{DashboardService: dashboardService}
}

func (controller *DashboardController) GetSalesResume(ctx *gin.Context) {
	res, err := controller.DashboardService.GetSalesResume(ctx.GetString("company_id"))

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	Helper.SetSuccessResponse(ctx, "Get Sales Resume Success!", res, http.StatusOK)
}

func (controller *DashboardController) GetBestSalesProduct(ctx *gin.Context) {
	limit := ctx.Query("limit")
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	limitInt := 0
	var err error

	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
	}
	if err != nil {
		Helper.SetErrorResponse(ctx, "Limit must be a number", http.StatusBadRequest)
		return
	}
	res, err := controller.DashboardService.GetBestSellingProducts(ctx.GetString("company_id"), startDate, endDate, limitInt)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	Helper.SetSuccessResponse(ctx, "Get Best Selling Product Success!", res, http.StatusOK)
}

func (controller *DashboardController) GetRevenue(ctx *gin.Context) {
	yearStr := ctx.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
	yearInt, err := strconv.Atoi(yearStr)
	if err != nil {
		Helper.SetErrorResponse(ctx, "Invalid year format", http.StatusBadRequest)
		return
	}

	companyID := ctx.GetString("company_id")
	if companyID == "" {
		Helper.SetErrorResponse(ctx, "Company ID is required", http.StatusBadRequest)
		return
	}
	res, err := controller.DashboardService.GetRevenue(companyID, yearInt)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}
	Helper.SetSuccessResponse(ctx, "Get Revenue Success!", res, http.StatusOK)
}

func (controller *DashboardController) GetExpense(ctx *gin.Context) {
	yearStr := ctx.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
	yearInt, err := strconv.Atoi(yearStr)
	if err != nil {
		Helper.SetErrorResponse(ctx, "Invalid year format", http.StatusBadRequest)
		return
	}

	companyID := ctx.GetString("company_id")
	if companyID == "" {
		Helper.SetErrorResponse(ctx, "Company ID is required", http.StatusBadRequest)
		return
	}
	res, err := controller.DashboardService.GetExpense(companyID, yearInt)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}
	Helper.SetSuccessResponse(ctx, "Get Expense Success!", res, http.StatusOK)
}
