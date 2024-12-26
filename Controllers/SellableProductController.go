package Controllers

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Services"
	"2024_akutansi_project/Utils"
	"fmt"
	"io/ioutil"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type (
	ISellableProductController interface {
		GetAllSellableProduct(ctx *gin.Context)
		GetActiveSellableProduct(ctx *gin.Context)
		UpdateSellableProduct(ctx *gin.Context)
		Create(ctx *gin.Context)
		CreateWithAssignMaterial(ctx *gin.Context)
		AssignMaterial(ctx *gin.Context)
		UnAssignMAterial(ctx *gin.Context)
		FindById(ctx *gin.Context)
		Delete(ctx *gin.Context)
		Update(ctx *gin.Context)
	}

	SellableProductController struct {
		SellableProductService Services.ISellableProductService
	}
)

func SellableProductControllerProvider(SellableProductService Services.ISellableProductService) *SellableProductController {
	return &SellableProductController{
		SellableProductService: SellableProductService,
	}
}

func (controller *SellableProductController) GetAllSellableProduct(ctx *gin.Context) {
	query := Utils.InsertParams(ctx)

	sellableProducts, meta, statusCode, err := controller.SellableProductService.GetAll(ctx.GetString("company_id"), &query, false)
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

func (controller *SellableProductController) GetActiveSellableProduct(ctx *gin.Context) {
	query := Utils.InsertParams(ctx)

	sellableProducts, meta, statusCode, err := controller.SellableProductService.GetAll(ctx.GetString("company_id"), &query, true)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetPaginationResponse(ctx,
		"Berhasil mendapatkan data sellable product active",
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

	statusCode, err := controller.SellableProductService.UpdateStock(ctx.Param("id"), &updateSellableDTO)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil mengupdate stock sellable product", nil, statusCode)

}

func (controller *SellableProductController) Create(ctx *gin.Context) {
	var createSellableProduct Dto.CreateSellableProductDTO

	if err := ctx.ShouldBind(&createSellableProduct); err != nil {
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	if validationErrors := Utils.ValidateRequest(ctx, &createSellableProduct); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	image, err := Utils.UploadFile(ctx, "image", fmt.Sprintf("%s/%s", os.Getenv("UPLOAD_DIR"), "products"))
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	createSellableProduct.Image = image

	createSellableProduct.CompanyID = ctx.GetString("company_id")

	res, err := controller.SellableProductService.Create(&createSellableProduct)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil membuat sellable product", res, http.StatusCreated)
}

func (controller *SellableProductController) CreateWithAssignMaterial(ctx *gin.Context) {

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(string(body)) // Proses data sesuai kebutuhan

	ctx.JSON(http.StatusOK, gin.H{"message": string(body)})

	return

	var createSellableProduct Dto.CreateSellableProductWithAssignMaterialDTO

	// Bind data JSON ke struct
	if err := ctx.ShouldBindWith(&createSellableProduct, binding.Form); err != nil {
		fmt.Println("Binding error:", err)
		Helper.SetErrorResponse(ctx, "Kesalahan Input Data", 400)
		return
	}

	ctx.JSON(http.StatusOK, createSellableProduct)
	return

	// Validasi data
	if validationErrors := Utils.ValidateRequest(ctx, &createSellableProduct); validationErrors != nil {
		Helper.SetValidationErrorResponse(ctx, validationErrors)
		return
	}

	// Proses file gambar
	image, err := Utils.UploadFile(ctx, "image", fmt.Sprintf("%s/%s", os.Getenv("UPLOAD_DIR"), "products"))
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	createSellableProduct.Image = image
	createSellableProduct.CompanyID = ctx.GetString("company_id")

	// Panggil service untuk membuat produk
	res, err := controller.SellableProductService.CreateWithAsignMaterial(&createSellableProduct)
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

func (controller *SellableProductController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	res, statusCode, err := controller.SellableProductService.FindById(id)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil mendapatkan data sellable product", res, statusCode)

}

func (controller *SellableProductController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	statusCode, err := controller.SellableProductService.Delete(id)

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
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

	image, err := Utils.UploadFile(ctx, "image", fmt.Sprintf("%s/%s", os.Getenv("UPLOAD_DIR"), "products"))

	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	updateSellableProduct.Image = image

	statusCode, err := controller.SellableProductService.Update(&updateSellableProduct, id)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	Helper.SetSuccessResponse(ctx, "Berhasil mengupdate sellable product", nil, statusCode)
}
