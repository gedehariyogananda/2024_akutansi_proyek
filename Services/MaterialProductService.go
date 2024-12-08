package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
)

type (
	IMaterialProductService interface {
		Create(dto *Dto.CreateMaterialProductDto) (*Response.MaterialProduckResponse, error)
	}
	MaterialProductService struct {
		materialProductRepository    Repositories.IMaterialProductRepository
		materialConversionRepository Repositories.IMaterialConversionRepository
	}
)

func MaterialProductServiceProvider(materialProductRepository Repositories.IMaterialProductRepository, materialConversionRepository Repositories.IMaterialConversionRepository) *MaterialProductService {
	return &MaterialProductService{materialProductRepository: materialProductRepository, materialConversionRepository: materialConversionRepository}
}

func (s *MaterialProductService) Create(dto *Dto.CreateMaterialProductDto) (*Response.MaterialProduckResponse, error) {

	materialProduct := &Models.MaterialProduct{
		Name:           dto.Name,
		CompanyID:      dto.CompanyID,
		Status:         dto.Status,
		Sku:            dto.Sku,
		SmallestUnitID: dto.SmallestUnitID,
		CategoryID:     dto.CategoryID,
	}

	materialProduct, err := s.materialProductRepository.Create(materialProduct)
	if err != nil {
		return nil, err
	}

	for _, conversion := range dto.MaterialConversions {
		modelConversion := &Models.MaterialConversion{
			MaterialProductID: materialProduct.ID,
			UnitID:            conversion.UniID,
			Quantity:          conversion.Quantity,
		}
		_, err := s.materialConversionRepository.Create(modelConversion)
		if err != nil {
			return nil, err
		}
	}

	res := Response.ToMaterialProduckResponse(*materialProduct)

	return &res, nil
}
