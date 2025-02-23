package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	IDashboardController interface {
		GetSalesResume(ctx *gin.Context)
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
