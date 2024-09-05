package Services

import (
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Models/Mapper"
	"2024_akutansi_project/Repositories"
)

type (
	ISaleableProductService interface {
		FindAllSaleableProducts(company_id int) (saleableProduct *[]Response.SaleableResponseDTO, materialProduct *[]Response.MaterialResponseDTO, err error)
	}

	SaleableProductService struct {
		SaleableProductService Repositories.ISaleableProductRepository
		MaterialProductService Repositories.IMaterialProductRepository
	}
)

func SaleableProductServiceProvider(saleableProductRepository Repositories.ISaleableProductRepository, materialProductRepository Repositories.IMaterialProductRepository) *SaleableProductService {
	return &SaleableProductService{
		SaleableProductService: saleableProductRepository,
		MaterialProductService: materialProductRepository,
	}
}

func (s *SaleableProductService) FindAllSaleableProducts(company_id int) (saleableProduct *[]Response.SaleableResponseDTO, materialProduct *[]Response.MaterialResponseDTO, err error) {
	saleableProductInit, err := s.SaleableProductService.FindAll(company_id)
	if err != nil {
		return nil, nil, err
	}

	// for mapping
	saleableData := []Response.SaleableResponseDTO{}
	for _, sp := range *saleableProductInit {
		saleable := Mapper.ToSaleableProductResponsDTO(sp)
		saleableData = append(saleableData, saleable)
	}

	materialProductInit, err := s.MaterialProductService.FindByAvailableForSale(company_id)

	if err != nil {
		return nil, nil, err
	}

	// for mapping
	materialData := []Response.MaterialResponseDTO{}
	for _, mp := range *materialProductInit {
		material := Mapper.ToMaterialProductResponseDTO(mp)
		materialData = append(materialData, material)
	}

	return &saleableData, &materialData, nil

}
