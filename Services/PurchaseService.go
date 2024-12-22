package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
)

type (
	IPurchaseService interface {
		GetDropdown(companyID string) (res Response.DropDwonPurchase, err error)
	}

	PurchaseService struct {
		ProductRepository  Repositories.ISellableProductRepository
		MaterialRepository Repositories.IMaterialProductRepository
	}
)

func PurchaseServiceProvider(productRepository Repositories.ISellableProductRepository, materialRepository Repositories.IMaterialProductRepository) *PurchaseService {
	return &PurchaseService{ProductRepository: productRepository, MaterialRepository: materialRepository}
}

func (p *PurchaseService) GetDropdown(companyID string) (res Response.DropDwonPurchase, err error) {
	products, err := p.ProductRepository.GetProductWhithoutReceipt(companyID)
	if err != nil {
		return res, err
	}

	materials, err := p.MaterialRepository.FindByCompany(companyID)
	if err != nil {
		return res, err
	}

	productsValues := make([]Models.SellableProduct, len(products))
	for i, product := range products {
		productsValues[i] = *product
	}

	materialsValues := make([]Models.MaterialProduct, len(materials))
	for i, material := range materials {
		materialsValues[i] = *material
	}

	res = Response.ToDropDownPurchase(productsValues, materialsValues)

	return res, nil
}
