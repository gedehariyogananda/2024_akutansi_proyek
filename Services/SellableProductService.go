package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"net/http"
)

type (
	ISellableProductService interface {
		GetAll(companyID string, query *Common.Query, onlyActive bool) (response []*Response.SellableResponse, meta Common.Meta, statusCode int, err error)
		UpdateStock(id string, request *Dto.SellableProductDTO) (statusCode int, err error)
	}

	SellableProductService struct {
		SellableProductRepository Repositories.ISellableProductRepository
		PromoItemRepository       Repositories.IPromoItemRepository
	}
)

func SellableProductServiceProvider(sellableProductRepository Repositories.ISellableProductRepository, promoItemRepository Repositories.IPromoItemRepository) *SellableProductService {
	return &SellableProductService{
		SellableProductRepository: sellableProductRepository,
		PromoItemRepository:       promoItemRepository,
	}
}

func (service *SellableProductService) GetAll(companyID string, query *Common.Query, onlyActive bool) (response []*Response.SellableResponse, meta Common.Meta, statusCode int, err error) {

	sellableProducts, totalData, err := service.SellableProductRepository.GetAll(companyID, onlyActive, query)
	if err != nil {
		return nil, Common.Meta{}, http.StatusInternalServerError, err
	}

	meta = Common.Meta{
		TotalData: totalData,
		Limit:     query.Limit,
		Page:      query.Page,
	}

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

		if onlyActive {
			res = append(res, &Response.SellableResponse{
				ID:              sellableProduct.ID,
				Name:            &sellableProduct.Name,
				Image:           &sellableProduct.Image,
				Description:     &sellableProduct.Description,
				CurrentQuantity: &sellableProduct.CurrentQuantity,
				Price:           &sellableProduct.Price,
			})
			continue
		} else {
			res = append(res, &Response.SellableResponse{
				ID:              sellableProduct.ID,
				Name:            &sellableProduct.Name,
				CompanyID:       &sellableProduct.CompanyID,
				SmallestUnitID:  &sellableProduct.SmallestUnitID,
				CategoryID:      &sellableProduct.CategoryID,
				Image:           &sellableProduct.Image,
				Description:     &sellableProduct.Description,
				Status:          sellableProduct.Status,
				StatusDisplay:   &status,
				HasReceipt:      &sellableProduct.HasReceipt,
				CurrentQuantity: &sellableProduct.CurrentQuantity,
				Price:           &sellableProduct.Price,
				CreatedAt:       sellableProduct.CreatedAt,
				UpdatedAt:       sellableProduct.UpdatedAt,
				DeletedAt:       sellableProduct.DeletedAt,
				Unit:            sellableProduct.Unit,
				Category:        sellableProduct.Category,
				PromoItems:      sellableProduct.PromoItems,
			})
	
		}	
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
