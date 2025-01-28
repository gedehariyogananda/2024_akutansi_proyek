package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	IStorageController interface {
		GetSignedUrl(ctx *gin.Context)
	}

	StorageController struct {
		StorageService Services.IStorageService
	}
)

func StorageControllerProvider(StorageService Services.IStorageService) *StorageController {
	return &StorageController{
		StorageService: StorageService,
	}
}

func (c *StorageController) GetSignedUrl(ctx *gin.Context) {
	var req Dto.StorageRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		Helper.SetValidationErrorResponse(ctx, err.Error())
		return
	}

	url, err := c.StorageService.SignedUrl(req)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	Helper.SetSuccessResponse(ctx, "berhasil mendapatkan url signed", url, http.StatusOK)

}
