package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Repositories"
)

type (
	ICategoryService interface {
		FindAllCategory(company_id string) (category *[]Models.Category, err error)
	}

	CategoryService struct {
		CategoryRepository Repositories.ICategoryRepository
	}
)

func CategoryServiceProvider(categoryRepository Repositories.ICategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepository: categoryRepository}
}

func (s *CategoryService) FindAllCategory(company_id string) (category *[]Models.Category, err error) {
	category, err = s.CategoryRepository.FindAll(company_id)

	if err != nil {
		return nil, err
	}

	return category, nil
}
