package Mapper

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto/Response"
)

func ToSaleableProductResponsDTO(saleableProduct Models.SaleableProduct) Response.SaleableResponseDTO {
	return Response.SaleableResponseDTO{
		ID:                saleableProduct.ID,
		CompanyID:         saleableProduct.CompanyID,
		ProductName:       saleableProduct.ProductName,
		UnitPrice:         saleableProduct.UnitPrice,
		CategoryName:      saleableProduct.Category.CategoryName,
		IsSaleableProduct: true,
	}
}

func ToMaterialProductResponseDTO(materialProduct Models.MaterialProduct) Response.MaterialResponseDTO {
	return Response.MaterialResponseDTO{
		ID:                materialProduct.ID,
		CompanyID:         materialProduct.CompanyID,
		ProductName:       materialProduct.MaterialProductName,
		UnitPrice:         materialProduct.UnitPriceForSelling,
		CategoryName:      "",
		IsSaleableProduct: false,
	}
}
