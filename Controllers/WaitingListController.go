package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IWaitingListController interface {
		Insert(ctx *gin.Context)
	}

	WaitingListController struct {
		service Services.IWaitingListService
	}
)

func WaitingListControllerProvider(service Services.IWaitingListService) *WaitingListController {
	return &WaitingListController{service: service}
}

func (c *WaitingListController) Insert(ctx *gin.Context) {
	var request Dto.MakeWaitingListRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": "Invalid request body",
		}, http.StatusBadRequest)
		return
	}

	if err := c.service.InsertWaitingList(&request); err != nil {
		Helper.SetResponse(ctx, gin.H{
			"success": false,
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	Helper.SetResponse(ctx, gin.H{
		"success": true,
		"message": "Waiting list created",
	}, http.StatusCreated)
}
