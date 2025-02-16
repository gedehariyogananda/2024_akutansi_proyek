package Services

import "2024_akutansi_project/Repositories"

type (
	IDashboardService interface {
	}

	DashboardService struct {
		productRepository Repositories.ISellableProductRepository
	}
)

func DashboardServiceProvider(productRepository Repositories.ISellableProductRepository) *DashboardService {
	return &DashboardService{productRepository: productRepository}
}
