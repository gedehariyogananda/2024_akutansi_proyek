package Services

import (
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Models/Mapper"
	"2024_akutansi_project/Repositories"
	"net/http"
)

type (
	ISaleableProductService interface {
		FindAllSaleableProducts(company_id string, categoryQuery string) (saleableProduct *[]Response.SaleableResponseDTO, materialProduct *[]Response.MaterialResponseDTO, err error, statusCode int)
	}

	SaleableProductService struct {
		SaleableProductRepository Repositories.ISaleableProductRepository
		MaterialProductRepository Repositories.IMaterialProductRepository
		CategoryRepository        Repositories.ICategoryRepository
	}
)

func SaleableProductServiceProvider(SaleableProductRepository Repositories.ISaleableProductRepository, MaterialProductRepository Repositories.IMaterialProductRepository, CategoryRepository Repositories.ICategoryRepository) *SaleableProductService {
	return &SaleableProductService{
		SaleableProductRepository: SaleableProductRepository,
		MaterialProductRepository: MaterialProductRepository,
		CategoryRepository:        CategoryRepository,
	}
}

func (s *SaleableProductService) FindAllSaleableProducts(company_id string, categoryQuery string) (saleableProduct *[]Response.SaleableResponseDTO, materialProduct *[]Response.MaterialResponseDTO, err error, statusCode int) {
	if categoryQuery == "Lainnya" {
		materialProductInit, err := s.MaterialProductRepository.FindByAvailableForSale(company_id)
		if err != nil {
			return nil, nil, err, http.StatusInternalServerError
		}

		materialData := []Response.MaterialResponseDTO{}
		if materialProductInit != nil {
			for _, mp := range *materialProductInit {
				material := Mapper.ToMaterialProductResponseDTO(mp)
				materialData = append(materialData, material)
			}
		}

		return nil, &materialData, nil, http.StatusOK
	}

	// Handle case when category is empty (get all products)
	if categoryQuery == "" {
		saleableProductInit, err := s.SaleableProductRepository.FindAll(company_id)
		if err != nil {
			return nil, nil, err, http.StatusInternalServerError
		}

		saleableData := []Response.SaleableResponseDTO{}
		if saleableProductInit != nil {
			for _, sp := range *saleableProductInit {
				saleable := Mapper.ToSaleableProductResponsDTO(sp)
				saleableData = append(saleableData, saleable)
			}
		}

		materialProductInit, err := s.MaterialProductRepository.FindByAvailableForSale(company_id)
		if err != nil {
			return nil, nil, err, http.StatusInternalServerError
		}

		materialData := []Response.MaterialResponseDTO{}
		if materialProductInit != nil {
			for _, mp := range *materialProductInit {
				material := Mapper.ToMaterialProductResponseDTO(mp)
				materialData = append(materialData, material)
			}
		}

		return &saleableData, &materialData, nil, http.StatusOK
	}

	if categoryQuery != "" {
		category, err := s.CategoryRepository.FindByName(categoryQuery)
		if err != nil {
			return nil, nil, err, http.StatusNotFound
		}

		saleableProductInit, err := s.SaleableProductRepository.FindByCategory(company_id, category.ID)
		if err != nil {
			return nil, nil, err, http.StatusInternalServerError
		}

		saleableData := []Response.SaleableResponseDTO{}
		if saleableProductInit != nil {
			for _, sp := range *saleableProductInit {
				saleable := Mapper.ToSaleableProductResponsDTO(sp)
				saleableData = append(saleableData, saleable)
			}
		}

		return &saleableData, nil, nil, http.StatusOK
	}

	return nil, nil, nil, http.StatusOK
}
