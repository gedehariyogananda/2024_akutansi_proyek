package Controllers

import (
	"2024_akutansi_project/Consts"
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	IJournalEntriesController interface {
		FindAll(ctx *gin.Context)
		TrialBalanceReport(ctx *gin.Context)
		FinancialBalanceReport(ctx *gin.Context)
		ProfitOrLossReport(ctx *gin.Context)
	}

	JournalEntriesController struct {
		JournalEntriesService Services.IJournalEntriesService
	}
)

func JournalEntriesProvider(JournalEntriesService Services.IJournalEntriesService) *JournalEntriesController {
	return &JournalEntriesController{JournalEntriesService: JournalEntriesService}
}

func (controller *JournalEntriesController) FindAll(ctx *gin.Context) {
	var request Dto.GetJournalRequest
	query := Utils.InsertParams(ctx)

	request.AccountID = ctx.Query("account_id")
	request.StartDate = ctx.Query("start_date")
	request.EndDate = ctx.Query("end_date")
	request.Query = query

	key := Consts.JOURNAL_ENTRY
	data, meta, err := controller.JournalEntriesService.FindAll(request, &key)

	if err != nil {
		statusCode := Utils.HandleStatusCode(err)
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetPaginationResponse(ctx,
		"Berhasil mendapatkan data jurnal",
		data,
		meta,
		http.StatusOK,
	)
}

func (controller *JournalEntriesController) TrialBalanceReport(ctx *gin.Context) {
	var request Dto.GetJournalRequest
	query := Utils.InsertParams(ctx)

	yesterdayDate := Utils.ShiftDate(time.Now(), -1)

	request.PeriodDate = ctx.DefaultQuery("period_date", yesterdayDate.Format("2006-01-02"))
	request.Query = query

	key := Consts.TRIAL_BALANCE_REPORT
	data, meta, err := controller.JournalEntriesService.FindAll(request, &key)

	if err != nil {
		statusCode := Utils.HandleStatusCode(err)
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetPaginationResponse(ctx,
		"Berhasil mendapatkan laporan neraca saldo",
		data,
		meta,
		http.StatusOK,
	)
}

func (controller *JournalEntriesController) FinancialBalanceReport(ctx *gin.Context) {
	var request Dto.GetJournalRequest
	query := Utils.InsertParams(ctx)

	yesterdayDate := Utils.ShiftDate(time.Now(), -1)

	request.PeriodDate = ctx.DefaultQuery("period_date", yesterdayDate.Format("2006-01-02"))
	request.Query = query

	key := Consts.FINANCIAL_BALANCE_REPORT
	data, meta, err := controller.JournalEntriesService.FindAll(request, &key)

	if err != nil {
		statusCode := Utils.HandleStatusCode(err)
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetPaginationResponse(ctx,
		"Berhasil mendapatkan laporan neraca keuangan",
		data,
		meta,
		http.StatusOK,
	)
}

func (controller *JournalEntriesController) ProfitOrLossReport(ctx *gin.Context) {
	yearString := ctx.Query("year")
	year, err := strconv.Atoi(yearString)

	if err != nil {
		Helper.SetErrorResponse(ctx, "Tahun tidak valid", http.StatusBadRequest)
		return
	}

	res, err := controller.JournalEntriesService.ProfitLossReport(year)

	if err != nil {
		statusCode := Utils.HandleStatusCode(err)
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil mendapatkan laporan laba rugi", res, http.StatusOK)
}
