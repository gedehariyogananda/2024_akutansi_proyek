package Services

import (
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"fmt"
	"net/http"
)

type (
	ISaleableProductService interface {
		FindAllSaleableProducts(company_id string, categoryQuery []string) (saleableProduct *[]Response.SaleableResponseDTO, err error, statusCode int)
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

// func (s *SaleableProductService) FindAllSaleableProducts(company_id string, categoryQuery string) (saleableProduct *[]Response.SaleableResponseDTO, err error, statusCode int) {

// 	saleableProduct = &[]Response.SaleableResponseDTO{}

// 	if categoryQuery == "Others" {
// 		materialProductInit, err := s.MaterialProductRepository.FindByAvailableForSale(company_id)
// 		if err != nil {
// 			return nil, err, http.StatusNotFound
// 		}

// 		if materialProductInit != nil {
// 			for _, item := range *materialProductInit {
// 				materialData := Response.SaleableResponseDTO{
// 					ID:           item.ID,
// 					ProductName:  item.MaterialProductName,
// 					UnitPrice:    item.UnitPriceForSelling,
// 					CategoryName: "",
// 				}

// 				*saleableProduct = append(*saleableProduct, materialData)
// 			}
// 		}

// 		return saleableProduct, nil, http.StatusOK
// 	}

// 	// Handle case when category is empty (get all products)
// 	if categoryQuery == "" {
// 		saleableProductInit, err := s.SaleableProductRepository.FindAll(company_id)
// 		if err != nil {
// 			return nil, err, http.StatusInternalServerError
// 		}

// 		if saleableProductInit != nil {
// 			for _, item := range *saleableProductInit {
// 				saleableData := Response.SaleableResponseDTO{
// 					ID:           item.ID,
// 					ProductName:  item.ProductName,
// 					UnitPrice:    item.UnitPrice,
// 					CategoryName: item.Category.CategoryName,
// 				}

// 				*saleableProduct = append(*saleableProduct, saleableData)
// 			}
// 		}

// 		materialProductInit, err := s.MaterialProductRepository.FindByAvailableForSale(company_id)
// 		if err != nil {
// 			return nil, err, http.StatusInternalServerError
// 		}

// 		if materialProductInit != nil {
// 			for _, item := range *materialProductInit {
// 				materialData := Response.SaleableResponseDTO{
// 					ID:           item.ID,
// 					ProductName:  item.MaterialProductName,
// 					UnitPrice:    item.UnitPriceForSelling,
// 					CategoryName: "",
// 				}

// 				*saleableProduct = append(*saleableProduct, materialData)
// 			}
// 		}

// 		return saleableProduct, nil, http.StatusOK
// 	}

// 	category, err := s.CategoryRepository.FindByName(categoryQuery)
// 	if err != nil {
// 		return nil, err, http.StatusNotFound
// 	}

// 	saleableProductInit, err := s.SaleableProductRepository.FindByCategory(company_id, category.ID)
// 	if err != nil {
// 		return nil, err, http.StatusInternalServerError
// 	}

// 	if saleableProductInit != nil {
// 		for _, item := range *saleableProductInit {
// 			saleableData := Response.SaleableResponseDTO{
// 				ID:           item.ID,
// 				ProductName:  item.ProductName,
// 				UnitPrice:    item.UnitPrice,
// 				CategoryName: item.Category.CategoryName,
// 			}

// 			*saleableProduct = append(*saleableProduct, saleableData)
// 		}
// 	}

// 	return saleableProduct, nil, http.StatusOK

// }

func (s *SaleableProductService) FindAllSaleableProducts(company_id string, categoryQueries []string) (saleableProduct *[]Response.SaleableResponseDTO, err error, statusCode int) {
	saleableProduct = &[]Response.SaleableResponseDTO{}

	if len(categoryQueries) == 0 {
		saleableProductInit, err := s.SaleableProductRepository.FindAll(company_id)
		if err != nil {
			return nil, err, http.StatusInternalServerError
		}

		// Pengolahan saleable products tetap sama
		if saleableProductInit != nil {
			for _, item := range *saleableProductInit {
				saleableData := Response.SaleableResponseDTO{
					ID:           item.ID,
					ProductName:  item.ProductName,
					UnitPrice:    item.UnitPrice,
					CategoryName: item.Category.CategoryName,
				}

				*saleableProduct = append(*saleableProduct, saleableData)
			}
		}

		materialProductInit, err := s.MaterialProductRepository.FindByAvailableForSale(company_id)
		if err != nil {
			return nil, err, http.StatusInternalServerError
		}

		if materialProductInit != nil {
			for _, item := range *materialProductInit {
				materialData := Response.SaleableResponseDTO{
					ID:           item.ID,
					ProductName:  item.MaterialProductName,
					UnitPrice:    item.UnitPriceForSelling,
					CategoryName: "",
				}

				*saleableProduct = append(*saleableProduct, materialData)
			}
		}

		return saleableProduct, nil, http.StatusOK
	}

	categories, err := s.CategoryRepository.FindByNames(categoryQueries)
	if err != nil || len(categories) == 0 {
		return nil, fmt.Errorf("categories not found"), http.StatusNotFound
	}

	var categoryIDs []string
	for _, category := range categories {
		categoryIDs = append(categoryIDs, category.ID)
	}

	saleableProductInit, err := s.SaleableProductRepository.FindByCategory(company_id, categoryIDs)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	if saleableProductInit != nil {
		for _, item := range *saleableProductInit {
			saleableData := Response.SaleableResponseDTO{
				ID:           item.ID,
				ProductName:  item.ProductName,
				UnitPrice:    item.UnitPrice,
				CategoryName: item.Category.CategoryName,
			}

			*saleableProduct = append(*saleableProduct, saleableData)
		}
	}

	return saleableProduct, nil, http.StatusOK
}
