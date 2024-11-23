package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type (
	ICategoryService interface {
		FindAll(company_id string, query *Common.Query) (res []*Response.Category, meta Common.Meta, err error)
		Create(Category *Dto.CreateCategory, company_id string) (res *Response.Category, statusCode int, err error)
		Update(request *Dto.UpdateCategory, id string) (res *Response.Category, statusCode int, err error)
		FindByID(id string) (res *Response.Category, statusCode int, err error)
		Delete(id string) (statusCode int, err error)
	}

	CategoryService struct {
		CategoryRepository Repositories.ICategoryRepository
	}
)

func CategoryServiceProvider(categoryRepository Repositories.ICategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepository: categoryRepository}
}

func (s *CategoryService) FindAll(company_id string, query *Common.Query) (res []*Response.Category, meta Common.Meta, err error) {
	categories, totalData, err := s.CategoryRepository.FindAll(company_id, query)

	if err != nil {
		return nil, meta, err
	}

	res = Response.ToCategorySlice(categories)

	meta = Common.Meta{
		TotalData: totalData,
		Page:      query.Page,
		Limit:     query.Limit,
	}

	return res, meta, nil

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

func (s *CategoryService) Delete(id string) (statusCode int, err error) {

	_, err = s.CategoryRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, err
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = s.CategoryRepository.Delete(id)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
