package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"net/http"
)

type (
	ICategoryService interface {
		FindAllCategory(company_id string) (category *[]Models.Category, err error)
		CreateCategory(request *Dto.CreateCategoryRequestDTO, company_id string) (category *Models.Category, statusCode int, err error)
		UpdateCategory(request *Dto.UpdateCategoryRequestDTO, id string, company_id string) (category *Models.Category, statusCode int, err error)
		DeleteCategory(id string) (statusCode int, err error)
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

func (s *CategoryService) CreateCategory(request *Dto.CreateCategoryRequestDTO, company_id string) (category *Models.Category, statusCode int, err error) {
	category, err = s.CategoryRepository.Create(request, company_id)

	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return category, http.StatusOK, nil
}

func (s *CategoryService) UpdateCategory(request *Dto.UpdateCategoryRequestDTO, id string, company_id string) (category *Models.Category, statusCode int, err error) {
	category, err = s.CategoryRepository.FindById(id)

	if err != nil {
		return nil, http.StatusNotFound, err
	}

	category, err = s.CategoryRepository.Update(request, id, company_id)

	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return category, http.StatusOK, nil
}

func (s *CategoryService) DeleteCategory(id string) (statusCode int, err error) {
	err = s.CategoryRepository.Delete(id)

	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
