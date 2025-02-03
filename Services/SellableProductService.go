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
	"time"

	"gorm.io/gorm"
)

type (
	ISellableProductService interface {
		Create(request *Dto.CreateSellableProductDTO) (res *Response.SellableResponse, err error)
		GetAll(companyID string, query *Dto.GetSellableProduct) (response []*Response.SellableResponse, meta Common.Meta, statusCode int, err error)
		AssignMaterial(request *Dto.AssignMaterialDtos) (statusCode int, err error)
		UnAssignMaterial(request *Dto.UnAssignMaterialDto) (statusCode int, err error)
		FindById(id string, setWithMaterial bool) (res *Response.SellableResponse, statusCode int, err error)
		Delete(id string) (statusCode int, objectKey string, err error)
		Update(request *Dto.UpdateSellableProductDTO, id string) (statusCode int, oldImage string, err error)
		UpdateStock(id string, request *Dto.SellableProductDTO, companyID string) (statusCode int, addMessage *string, err error)
	}

	SellableProductService struct {
		SellableProductRepository Repositories.ISellableProductRepository
		SellableStockRepository   Repositories.ISellableStockRepository
		PromoItemRepository       Repositories.IPromoItemRepository
		ReceiptRepository         Repositories.IReceiptRepository
		StorageService            IStorageService
		DB                        *gorm.DB
	}
)

func SellableProductServiceProvider(sellableProductRepository Repositories.ISellableProductRepository, promoItemRepository Repositories.IPromoItemRepository, receiptRepository Repositories.IReceiptRepository, sellableStockRepository Repositories.ISellableStockRepository, DB *gorm.DB, storageService IStorageService) *SellableProductService {
	return &SellableProductService{
		SellableProductRepository: sellableProductRepository,
		PromoItemRepository:       promoItemRepository,
		StorageService:            storageService,
		ReceiptRepository:         receiptRepository,
		SellableStockRepository:   sellableStockRepository,
		DB:                        DB,
	}
}

func (service *SellableProductService) presignedURL(objectKey string) (string, error) {
	presignedURL, err := service.StorageService.SignedUrl(Dto.StorageRequest{
		ObjectKey: objectKey,
	})

	if err != nil {
		return "", err
	}

	return presignedURL, nil
}

func (service *SellableProductService) GetAll(companyID string, query *Dto.GetSellableProduct) (response []*Response.SellableResponse, meta Common.Meta, statusCode int, err error) {
	sellableProducts, totalData, err := service.SellableProductRepository.GetAll(companyID, query)

	if err != nil {
		return nil, Common.Meta{}, http.StatusInternalServerError, err
	}

	var res []*Response.SellableResponse

	for _, sellableProduct := range sellableProducts {
		status := ""
		if *sellableProduct.Status && sellableProduct.CurrentQuantity > 0 {
			status = "Aktif"
		} else if sellableProduct.CurrentQuantity == 0 {
			status = "Habis"
		} else {
			status = "Non-Aktif"
		}

		var promo *Models.Promo
		if len(sellableProduct.PromoItems) > 0 {
			promo = sellableProduct.PromoItems[0].Promo
		}

		presignedURL, err := service.presignedURL(sellableProduct.Image)
		if err != nil {
			return nil, Common.Meta{}, http.StatusInternalServerError, err
		}

		res = append(res, &Response.SellableResponse{
			ID:              sellableProduct.ID,
			Name:            &sellableProduct.Name,
			Image:           &presignedURL,
			CurrentQuantity: &sellableProduct.CurrentQuantity,
			Price:           &sellableProduct.Price,
			Category:        sellableProduct.Category,
			StatusDisplay:   &status,
			Promo:           promo,
		})
	}

	meta = Common.PaginateMetadata(nil, totalData, query.Limit, query.Page)

	return res, meta, http.StatusOK, nil
}

func (service *SellableProductService) UpdateStock(id string, request *Dto.SellableProductDTO, companyID string) (statusCode int, addMessage *string, err error) {
	var status bool

	product, err := service.SellableProductRepository.FindByID(id, false)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, nil, err
	}

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if !product.HasReceipt {
		if request.CurrentQuantity != nil && product.CurrentQuantity != *request.CurrentQuantity {
			return http.StatusBadRequest, nil, errors.New("product tidak bisa diupdate, harus melakukan pembelian terlebih dahulu")
		}
	}

	if request.PromoID == nil {
		_, isExist, _ := service.PromoItemRepository.FindBySellableID(id)

		if isExist {
			if err := service.PromoItemRepository.DeleteBySellableProductID(id); err != nil {
				return http.StatusInternalServerError, nil, err
			}
		}
	} else {
		promoItems := &Models.PromoItem{
			SellableProductID: id,
		}

		if request.PromoID != nil {
			promoItems.PromoID = *request.PromoID
		}

		_, _, err = service.PromoItemRepository.UpdateOrCreate(promoItems)

		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		promoItem, _, err := service.PromoItemRepository.FindBySellableID(id)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		if time.Now().After(promoItem.Promo.EndDate) {
			prefMessage := "Promo yang anda masukkan sudah berakhir di tanggal " + promoItem.Promo.EndDate.Format("02-01-2006") + ", Promo tidak akan tampil di layar kasir."
			addMessage = &prefMessage
		} else if time.Now().Before(promoItem.Promo.StartDate) {
			prefMessage := "Promo yang anda masukkan belum dimulai!, Promo akan bisa dipakai dan ditampilkan di kasir di tanggal " + promoItem.Promo.StartDate.Format("02-01-2006")
			addMessage = &prefMessage
		}
	}

	if request.Status != nil {
		status = *request.Status
	}

	if err = service.SellableProductRepository.Update(id,
		&Models.SellableProduct{
			CurrentQuantity: *request.CurrentQuantity,
			Status:          &status,
			Description:     *request.Description,
		}); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, addMessage, nil
}

func (service *SellableProductService) Create(request *Dto.CreateSellableProductDTO) (res *Response.SellableResponse, err error) {
	trx := service.DB.Begin()

	if trx.Error != nil {
		return nil, trx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
			err = fmt.Errorf("panic occurred: %v", r)
		} else if err != nil {
			trx.Rollback()
		} else {
			trx.Commit()
		}
	}()

	sellableProduct := &Models.SellableProduct{
		Name:           request.Name,
		SmallestUnitID: request.SmallestUnitID,
		CategoryID:     request.CategoryID,
		Description:    request.Description,
		Status:         request.Status,
		HasReceipt:     request.HasReceipt,
		Price:          request.Price,
		Sku:            request.Sku,
		Image:          request.Image,
		CompanyID:      request.CompanyID,
	}

	product, err := service.SellableProductRepository.Create(sellableProduct)

	if err != nil {
		return res, err
	}

	if !request.HasReceipt && request.MaterialsObj == nil {
		return res, errors.New("if don't have receipt can't send materials")
	}

	if request.MaterialsObj != nil {
		for _, material := range *request.MaterialsObj {
			receipts := &Models.Receipt{
				SellableProductID: product.ID,
				MaterialProductID: material.MaterialID,
				Quantity:          int(material.Quantity),
			}
			err = service.ReceiptRepository.Create(receipts)

			if err != nil {
				return res, err
			}
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

func (s *SellableProductService) FindById(id string, setWithMaterial bool) (res *Response.SellableResponse, statusCode int, err error) {
	sellableProduct, err := s.SellableProductRepository.FindByID(id, setWithMaterial)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	presignedURL, err := s.presignedURL(sellableProduct.Image)
	if err != nil {
		return nil, http.StatusInternalServerError, err

	}

	var promo string
	if len(sellableProduct.PromoItems) > 0 {
		promo = sellableProduct.PromoItems[0].PromoID
	}

	isExpired := false
	if time.Now().After(sellableProduct.PromoItems[0].Promo.EndDate) {
		isExpired = true
	}

	if setWithMaterial {
		res = Response.ToSellableResponse(sellableProduct)
	} else {
		res = &Response.SellableResponse{
			ID:              sellableProduct.ID,
			Name:            &sellableProduct.Name,
			Image:           &presignedURL,
			Status:          sellableProduct.Status,
			CurrentQuantity: &sellableProduct.CurrentQuantity,
			Description:     &sellableProduct.Description,
			Category:        sellableProduct.Category,
			HasReceipt:      &sellableProduct.HasReceipt,
			IsExpiredPromo:  &isExpired,
		}

		if promo != "" {
			res.PromoID = &promo
		}
	}

	return res, http.StatusOK, nil
}

func (s *SellableProductService) Delete(id string) (statusCode int, objectKey string, err error) {
	trx := s.DB.Begin()

	if trx.Error != nil {
		return http.StatusInternalServerError, "", trx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
			err = fmt.Errorf("panic occurred: %v", r)
		} else if err != nil {
			trx.Rollback()
		} else {
			trx.Commit()
		}
	}()

	data, err := s.SellableProductRepository.FindByID(id, false)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, "", err
	}

	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	err = s.SellableProductRepository.Delete(id)

	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	err = s.ReceiptRepository.DeleteByProductId(id)

	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	return http.StatusOK, data.Image, nil
}

func (s *SellableProductService) Update(dto *Dto.UpdateSellableProductDTO, id string) (statusCode int, oldImage string, err error) {
	product, err := s.SellableProductRepository.FindByID(id, false)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, "", err
	}
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	var image string

	oldImage = ""

	if dto.Image == "" {
		image = product.Image
	} else if dto.Image != product.Image {
		oldImage = product.Image
		image = dto.Image

		err = Utils.DeleteFile(product.Image)
		if err != nil {
			return http.StatusInternalServerError, "", err
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
		HasReceipt:     dto.HasReceipt,
	}

	if err = s.SellableProductRepository.Update(id, sellableProduct); err != nil {
		return http.StatusInternalServerError, "", err
	}

	return http.StatusOK, oldImage, nil
}
