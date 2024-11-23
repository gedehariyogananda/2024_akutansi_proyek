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
	ICategoryService interface {
		FindAllCategory(company_id string) (category *[]Models.Category, err error)
		Create(Category *Dto.CreateCategory, company_id string) (res *Response.Category, statusCode int, err error)
		Update(request *Dto.UpdateCategory, id string) (res *Response.Category, statusCode int, err error)
		FindByID(id string) (res *Response.Category, statusCode int, err error)
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

func (s *CategoryService) Create(dto *Dto.CreateCategory, id string) (res *Response.Category, statusCode int, err error) {
	category := &Models.Category{
		Name:      dto.Name,
		CompanyID: id,
		Type:      dto.Type,
		Status:    dto.Status,
		Code:      dto.Code,
	}

	category, err = s.CategoryRepository.Create(category)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res = Response.ToCategory(category)

	return res, http.StatusOK, nil
}

func (s *CategoryService) Update(request *Dto.UpdateCategory, id string) (res *Response.Category, statusCode int, err error) {
	_, err = s.CategoryRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	category := &Models.Category{
		Name:   request.Name,
		Type:   request.Type,
		Status: request.Status,
		Code:   request.Code,
	}
	category.ID = id

	category, err = s.CategoryRepository.Update(category, id)

	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	res = Response.ToCategory(category)

	return res, http.StatusOK, nil
}

func (s *CategoryService) FindByID(id string) (res *Response.Category, statusCode int, err error) {
	category, err := s.CategoryRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// fmt.Println(err)
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res = Response.ToCategory(category)

	return res, http.StatusOK, nil
}

func (s *CategoryService) DeleteCategory(id string) (statusCode int, err error) {
	err = s.CategoryRepository.Delete(id)

	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
