package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"

	"gorm.io/gorm"
)

type (
	ICategoryRepository interface {
		FindAll(company_id string) (category *[]Models.Category, err error)
		FindByNames(category_names []string) (category []Models.Category, err error)
		Create(request *Dto.CreateCategoryRequestDTO, company_id string) (category *Models.Category, err error)
		Update(request *Dto.UpdateCategoryRequestDTO, id string, company_id string) (category *Models.Category, err error)
		Delete(id string) (err error)
		FindById(id string) (category *Models.Category, err error)
	}

	CategoryRepository struct {
		DB *gorm.DB
	}
)

func CategoryRepositoryProvider(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) FindAll(company_id string) (category *[]Models.Category, err error) {
	if err := r.DB.Where("company_id = ?", company_id).Find(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) FindByNames(category_names []string) (categories []Models.Category, err error) {
	if err := r.DB.Where("category_name IN (?)", category_names).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) Create(request *Dto.CreateCategoryRequestDTO, company_id string) (category *Models.Category, err error) {
	category = &Models.Category{
		CategoryName: request.CategoryName,
		CompanyID:    company_id,
	}

	if err := r.DB.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) Update(request *Dto.UpdateCategoryRequestDTO, id string, company_id string) (category *Models.Category, err error) {
	category = &Models.Category{
		CategoryName: request.CategoryName,
		CompanyID:    company_id,
	}

	if err := r.DB.Model(category).Where("id = ?", id).Updates(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) Delete(id string) (err error) {
	if err := r.DB.Delete(&Models.Category{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) FindById(id string) (category *Models.Category, err error) {
	category = &Models.Category{}

	if err := r.DB.Where("id = ?", id).First(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
