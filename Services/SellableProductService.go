package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Repositories"
	"net/http"
)

type (
	ISellableProductService interface {
		GetAll(companyID string, query *Common.Query) (sellableProducts []*Models.SellableProduct, meta Common.Meta, statusCode int, err error)
	}

	SellableProductService struct {
		SellableProductRepository Repositories.ISellableProductRepository
	}
)

func SellableProductServiceProvider(sellableProductRepository Repositories.ISellableProductRepository) *SellableProductService {
	return &SellableProductService{SellableProductRepository: sellableProductRepository}
}

func (service *SellableProductService) GetAll(companyID string, query *Common.Query) (sellableProducts []*Models.SellableProduct, meta Common.Meta, statusCode int, err error) {
	sellableProducts, totalData, err := service.SellableProductRepository.GetAll(companyID, nil, query)
	if err != nil {
		return nil, Common.Meta{}, http.StatusInternalServerError, err
	}

	meta = Common.Meta{
		TotalData: totalData,
		Limit:     query.Limit,
		Page:      query.Page,
	}

	return sellableProducts, meta, http.StatusOK, nil
}
