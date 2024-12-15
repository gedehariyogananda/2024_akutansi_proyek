package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IJournalEntriesController interface {
		FindAll(c *gin.Context)
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

	data, meta, err := controller.JournalEntriesService.FindAll(request)

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
