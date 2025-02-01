package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"

	"github.com/gin-gonic/gin"
)

type (
	IUserController interface {
		UploadAvatar(ctx *gin.Context)
		GetCurrentUser(ctx *gin.Context)
	}

	UserController struct {
		userService    Services.IUserService
		storageService Services.IStorageService
	}
)

func UserControllerProvider(userService Services.IUserService, storageService Services.IStorageService) *UserController {
	return &UserController{
		userService:    userService,
		storageService: storageService,
	}
}

func (c *UserController) UploadAvatar(ctx *gin.Context) {
	var uploadAvatarDTO Dto.UploadAvatarDto

	if err := ctx.ShouldBind(&uploadAvatarDTO); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	if validationErrors := Utils.ValidateRequest(ctx, &uploadAvatarDTO); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	storageDto := Dto.StorageRequest{
		File:      uploadAvatarDTO.Avatar,
		ObjectKey: "avatar",
	}

	path, err := c.storageService.UploadFile(storageDto, true)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), 500)
		return
	}

	userID := ctx.GetString("id")
	if err = c.userService.UploadAvatar(userID, path); err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), 500)
		return
	}

	Helper.SetSuccessResponse(ctx, "Upload Avatar Berhasil!", nil, 200)
}

func (c *UserController) GetCurrentUser(ctx *gin.Context) {
	userID := ctx.GetString("id")

	user, err := c.userService.GetCurrentUser(userID)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), 500)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil mendapatkan data user", user, 200)
}
