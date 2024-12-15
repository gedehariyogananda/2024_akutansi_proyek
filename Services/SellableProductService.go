package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Utils"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type (
	ISellableProductService interface {
		Create(request *Dto.CreateSellableProductDTO) (res *Response.SellableResponse, err error)
		CreateWithAsignMaterial(request *Dto.CreateSellableProductWithAssignMaterialDTO) (res *Response.SellableResponse, err error)
		GetAll(companyID string, query *Common.Query) (response []*Response.SellableResponse, meta Common.Meta, statusCode int, err error)
		UpdateStock(id string, request *Dto.SellableProductDTO) (statusCode int, err error)
		AssignMaterial(request *Dto.AssignMaterialDtos) (statusCode int, err error)
		UnAssignMaterial(request *Dto.UnAssignMaterialDto) (statusCode int, err error)
		FindById(id string) (res *Response.SellableResponse, statusCode int, err error)
		Delete(id string) (statusCode int, err error)
		Update(request *Dto.UpdateSellableProductDTO, id string) (statusCode int, err error)
	}

	SellableProductService struct {
		SellableProductRepository Repositories.ISellableProductRepository
		PromoItemRepository       Repositories.IPromoItemRepository
		ReceiptRepository         Repositories.IReceiptRepository
	}
)

func SellableProductServiceProvider(sellableProductRepository Repositories.ISellableProductRepository, promoItemRepository Repositories.IPromoItemRepository, receiptRepository Repositories.IReceiptRepository) *SellableProductService {
	return &SellableProductService{
		SellableProductRepository: sellableProductRepository,
		PromoItemRepository:       promoItemRepository,
		ReceiptRepository:         receiptRepository,
	}
}

func (service *SellableProductService) GetAll(companyID string, query *Common.Query) (response []*Response.SellableResponse, meta Common.Meta, statusCode int, err error) {
	sellableProducts, totalData, err := service.SellableProductRepository.GetAll(companyID, nil, query)
	if err != nil {
		return nil, Common.Meta{}, http.StatusInternalServerError, err
	}

	meta = Common.Meta{
		TotalData: totalData,
		Limit:     query.Limit,
		Page:      query.Page,
	}

	// return sellableProducts, meta, http.StatusOK, nil
	var res []*Response.SellableResponse

	for _, sellableProduct := range sellableProducts {
		status := ""
		if *sellableProduct.Status {
			status = "Aktif"
		} else if !*sellableProduct.Status && sellableProduct.CurrentQuantity <= 0 {
			status = "Habis"
		} else {
			status = "Non-Aktif"
		}

		res = append(res, &Response.SellableResponse{
			ID:              sellableProduct.ID,
			Name:            sellableProduct.Name,
			CompanyID:       sellableProduct.CompanyID,
			SmallestUnitID:  sellableProduct.SmallestUnitID,
			CategoryID:      sellableProduct.CategoryID,
			Image:           sellableProduct.Image,
			Description:     sellableProduct.Description,
			Status:          *sellableProduct.Status,
			StatusDisplay:   status,
			HasReceipt:      sellableProduct.HasReceipt,
			CurrentQuantity: sellableProduct.CurrentQuantity,
			Price:           sellableProduct.Price,
			CreatedAt:       *sellableProduct.CreatedAt,
			UpdatedAt:       *sellableProduct.UpdatedAt,
			DeletedAt:       sellableProduct.DeletedAt,
			Unit:            sellableProduct.Unit,
			Category:        sellableProduct.Category,
			PromoItems:      sellableProduct.PromoItems,
		})

	}

	return res, meta, http.StatusOK, nil
}

func (service *SellableProductService) UpdateStock(id string, request *Dto.SellableProductDTO) (statusCode int, err error) {

	if request.PrefixDeletePromo != nil && *request.PrefixDeletePromo {
		result, err := service.PromoItemRepository.DeleteBySellableProductID(id)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		if result.RowsAffected == 0 {
			return http.StatusNotFound, nil
		}

		return http.StatusOK, nil
	}

	if request.PromoID != nil && *request.PromoID != "" {
		promoItems := &Models.PromoItem{}
		if request.PromoID != nil {
			promoItems.SellableProductID = id
			promoItems.PromoID = *request.PromoID
		}

		_, result, err := service.PromoItemRepository.UpdateOrCreate(promoItems)

		if err != nil {
			return http.StatusInternalServerError, err
		}

		if result.RowsAffected == 0 {
			return http.StatusNotFound, nil
		}
	}

	sellableProduct := &Models.SellableProduct{}

	if request.CurrentQuantity != nil {
		sellableProduct.CurrentQuantity = *request.CurrentQuantity
	}
	if request.Status != nil {
		sellableProduct.Status = request.Status
	}
	if request.Description != nil {
		sellableProduct.Description = *request.Description
	}

	if err = service.SellableProductRepository.Update(id, sellableProduct); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (service *SellableProductService) Create(request *Dto.CreateSellableProductDTO) (res *Response.SellableResponse, err error) {
	sellableProduct := &Models.SellableProduct{
		Name:           request.Name,
		SmallestUnitID: request.SmallestUnitID,
		CategoryID:     request.CategoryID,
		Description:    request.Description,
		Status:         request.Status,
		HasReceipt:     false,
		Price:          request.Price,
		Sku:            request.Sku,
		Image:          request.Image,
		CompanyID:      request.CompanyID,
	}

	product, err := service.SellableProductRepository.Create(sellableProduct)

	if err != nil {
		return res, err
	}

	res = Response.ToSellableResponse(product)

	return res, nil
}

func (service *SellableProductService) CreateWithAsignMaterial(request *Dto.CreateSellableProductWithAssignMaterialDTO) (res *Response.SellableResponse, err error) {
	sellableProduct := &Models.SellableProduct{
		Name:           request.Name,
		SmallestUnitID: request.SmallestUnitID,
		CategoryID:     request.CategoryID,
		Description:    request.Description,
		Status:         request.Status,
		HasReceipt:     true,
		Price:          request.Price,
		Sku:            request.Sku,
		Image:          request.Image,
		CompanyID:      request.CompanyID,
	}

	product, err := service.SellableProductRepository.Create(sellableProduct)

	if err != nil {
		return res, err
	}

	fmt.Println(request.Materials)
	fmt.Println("test")

	for _, material := range request.Materials {
		receipt := &Models.Receipt{
			SellableProductID: product.ID,
			MaterialProductID: material.MaterialID,
			Quantity:          material.Quantity,
		}

		if err = service.ReceiptRepository.Create(receipt); err != nil {
			return res, err
		}
	}

	res = Response.ToSellableResponse(product)

	return res, nil
}

func (service *SellableProductService) AssignMaterial(request *Dto.AssignMaterialDtos) (statusCode int, err error) {
	for _, material := range request.Materials {
		receipt := &Models.Receipt{
			SellableProductID: request.SellableProductID,
			MaterialProductID: material.MaterialID,
			Quantity:          material.Quantity,
		}

		if err = service.ReceiptRepository.Create(receipt); err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusCreated, nil
}

func (service *SellableProductService) UnAssignMaterial(request *Dto.UnAssignMaterialDto) (statusCode int, err error) {
	for _, materialID := range request.MaterialIDS {
		if err = service.ReceiptRepository.DeleteByMaterialId(materialID); err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}

func (s *SellableProductService) FindById(id string) (res *Response.SellableResponse, statusCode int, err error) {
	sellableProduct, err := s.SellableProductRepository.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	fmt.Println(sellableProduct.Category.Name)

	response := Response.ToSellableResponse(sellableProduct)

	return response, http.StatusOK, nil
}

func (s *SellableProductService) Delete(id string) (statusCode int, err error) {
	_, err = s.SellableProductRepository.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, err
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = s.SellableProductRepository.Delete(id)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = s.ReceiptRepository.DeleteByProductId(id)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *SellableProductService) Update(dto *Dto.UpdateSellableProductDTO, id string) (statusCode int, err error) {
	product, err := s.SellableProductRepository.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, err
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	var image string

	fmt.Println("test")
	fmt.Println(dto.Image)

	if dto.Image == "" {
		image = product.Image
	} else if dto.Image != product.Image {
		image = dto.Image

		err = Utils.DeleteFile(product.Image)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	sellableProduct := &Models.SellableProduct{
		Name:           dto.Name,
		SmallestUnitID: dto.SmallestUnitID,
		CategoryID:     dto.CategoryID,
		Description:    dto.Description,
		Status:         dto.Status,
		Price:          dto.Price,
		Sku:            dto.Sku,
		Image:          image,
	}

	if err = s.SellableProductRepository.Update(id, sellableProduct); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
