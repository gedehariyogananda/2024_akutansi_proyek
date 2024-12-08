package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type (
	IMaterialProductService interface {
		Create(dto *Dto.CreateMaterialProductDto) (*Response.MaterialProduckResponse, error)
		FindById(id string) (*Response.MaterialProduckResponse, int, error)
		Delete(id string) (statusCode int, err error)
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

func (s *MaterialProductService) FindById(id string) (*Response.MaterialProduckResponse, int, error) {
	materialProduct, err := s.materialProductRepository.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res := Response.ToMaterialProduckResponse(*materialProduct)

	return &res, http.StatusOK, nil
}

func (s *MaterialProductService) Delete(id string) (statusCode int, err error) {
	_, err = s.materialProductRepository.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, err
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = s.materialProductRepository.Delete(id)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = s.materialConversionRepository.DeleteMany(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
