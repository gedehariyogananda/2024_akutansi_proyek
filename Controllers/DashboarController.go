package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	IDashboardController interface {
		GetSalesResume(ctx *gin.Context)
		GetBestSalesProduct(ctx *gin.Context)
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

	fmt.Printf(startDate)
	fmt.Println(endDate)

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
