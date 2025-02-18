package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"
	"encoding/json"

	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ISellableProductController interface {
		GetAllSellableProduct(ctx *gin.Context)
		UpdateSellableProduct(ctx *gin.Context)
		Create(ctx *gin.Context)
		AssignMaterial(ctx *gin.Context)
		UnAssignMAterial(ctx *gin.Context)
		FindByIdSetMaterial(ctx *gin.Context)
		FindByIdSetStock(ctx *gin.Context)
		Delete(ctx *gin.Context)
		Update(ctx *gin.Context)
	}

	SellableProductController struct {
		SellableProductService Services.ISellableProductService
		StorageService         Services.IStorageService
	}
)

func SellableProductControllerProvider(SellableProductService Services.ISellableProductService, StorageService Services.IStorageService) *SellableProductController {
	return &SellableProductController{
		SellableProductService: SellableProductService,
		StorageService:         StorageService,
	}
}

func (controller *SellableProductController) GetAllSellableProduct(ctx *gin.Context) {
	var request Dto.GetSellableProduct
	query := Utils.InsertParams(ctx)
	inStatus := ctx.Query("in_status")
	request.InStatus = &inStatus

	request.Query = query

	sellableProducts, meta, statusCode, err := controller.SellableProductService.GetAll(ctx.GetString("company_id"), &request)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetPaginationResponse(ctx,
		"Berhasil mendapatkan data sellable product",
		sellableProducts,
		meta,
		statusCode)
}

func (controller *SellableProductController) UpdateSellableProduct(ctx *gin.Context) {
	var updateSellableDTO Dto.SellableProductDTO

	if err := ctx.ShouldBindJSON(&updateSellableDTO); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	if validationErrors := Utils.ValidateRequest(ctx, &updateSellableDTO); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	statusCode, addMessage, err := controller.SellableProductService.UpdateStock(ctx.Param("id"), &updateSellableDTO, ctx.GetString("company_id"))
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	if addMessage != nil {
		Helper.SetSuccessResponse(ctx, *addMessage, gin.H{
			"SET_MESSAGE_PROMO": true,
		}, statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil mengupdate stock sellable product", nil, statusCode)

}

func (controller *SellableProductController) Create(ctx *gin.Context) {
	var createSellableProduct Dto.CreateSellableProductDTO

	// BINDING SECTION START
	if err := ctx.ShouldBind(&createSellableProduct); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	var materials *[]Dto.ReceiptMaterialDto
	if createSellableProduct.Materials != "" {
		if err := json.Unmarshal([]byte(createSellableProduct.Materials), &materials); err != nil {
			Helper.SetErrorResponse(ctx, "Format Input Data Resep Salah", 400)
			return
		}

		createSellableProduct.MaterialsObj = materials
	}
	// BINDING SECTION END

	// VALIDATION SECTION START

	if validationErrors := Utils.ValidateRequest(ctx, &createSellableProduct); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	if materials != nil {
		for _, material := range *materials {
			if material.MaterialID == "" || material.Quantity <= 0 {
				Helper.SetErrorResponse(ctx, "Kesalahan Input Data Resep", 400)
				return
			}
		}
	}
	// VALIDATION SECTION END

	image, err := ctx.FormFile("image")
	if err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	filePath, err := controller.StorageService.UploadFile(Dto.StorageRequest{
		File:      image,
		ObjectKey: "products",
	}, true)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), 400)
		return
	}

	createSellableProduct.Image = filePath

	createSellableProduct.CompanyID = ctx.GetString("company_id")

	res, err := controller.SellableProductService.Create(&createSellableProduct)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil membuat sellable product", res, http.StatusCreated)
}

func (controller *SellableProductController) AssignMaterial(ctx *gin.Context) {
	var assignMaterial Dto.AssignMaterialDtos

	if err := ctx.Bind(&assignMaterial); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	if validationErrors := Utils.ValidateRequest(ctx, &assignMaterial); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	statusCode, err := controller.SellableProductService.AssignMaterial(&assignMaterial)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil menambahkan material", nil, statusCode)
}

func (controller *SellableProductController) UnAssignMAterial(ctx *gin.Context) {
	var unAssignMaterial Dto.UnAssignMaterialDto

	if err := ctx.Bind(&unAssignMaterial); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	if validationErrors := Utils.ValidateRequest(ctx, &unAssignMaterial); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	statusCode, err := controller.SellableProductService.UnAssignMaterial(&unAssignMaterial)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil menghapus material", nil, statusCode)
}

func (controller *SellableProductController) FindByIdSetMaterial(ctx *gin.Context) {
	id := ctx.Param("id")

	res, statusCode, err := controller.SellableProductService.FindById(id, true)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil mendapatkan data sellable product", res, statusCode)
}

func (controller *SellableProductController) FindByIdSetStock(ctx *gin.Context) {
	id := ctx.Param("id")

	res, statusCode, err := controller.SellableProductService.FindById(id, false)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil mendapatkan data sellable product", res, statusCode)
}

func (controller *SellableProductController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	statusCode, objectKey, err := controller.SellableProductService.Delete(id)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	// delete image from minio
	if err = controller.StorageService.DeleteFile(Dto.StorageRequest{
		ObjectKey: objectKey,
	}); err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil menghapus sellable product", nil, statusCode)
}

func (controller *SellableProductController) Update(ctx *gin.Context) {
	var updateSellableProduct Dto.UpdateSellableProductDTO

	if err := ctx.ShouldBind(&updateSellableProduct); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	if validationErrors := Utils.ValidateRequest(ctx, &updateSellableProduct); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	id := ctx.Param("id")

	image, err := ctx.FormFile("image")
	if err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	filePath, err := controller.StorageService.UploadFile(Dto.StorageRequest{
		File:      image,
		ObjectKey: "products",
	}, true)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	updateSellableProduct.Image = filePath

	statusCode, oldImage, err := controller.SellableProductService.Update(&updateSellableProduct, id)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	if oldImage != "" {
		if err = controller.StorageService.DeleteFile(Dto.StorageRequest{
			ObjectKey: oldImage,
		}); err != nil {
			Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	Helper.SetSuccessResponse(ctx, "Berhasil mengupdate sellable product", nil, statusCode)
}
