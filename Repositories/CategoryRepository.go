package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Utils"

	"gorm.io/gorm"
)

type (
	ICategoryRepository interface {
		FindAll(company_id string, query *Common.Query) (category []*Models.Category, totalData int64, err error)
		FindByNames(category_names []string) (category []Models.Category, err error)
		Create(category *Models.Category) (*Models.Category, error)
		Update(category *Models.Category, id string) (*Models.Category, error)
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

func (r *CategoryRepository) FindAll(company_id string, query *Common.Query) (categories []*Models.Category, totalData int64, err error) {

	err = r.DB.Scopes(Helper.FilterCompanyID(company_id), Utils.Paginate(query.Page, query.Limit), Helper.FilterStatus(query.Status), Helper.FilterSearch(*query.Search)).Find(&categories).Error

	if err != nil {
		return nil, 0, err
	}

	err = r.DB.Model(&categories).Scopes(Helper.FilterCompanyID(company_id), Utils.Paginate(query.Page, query.Limit), Helper.FilterStatus(query.Status), Helper.FilterSearch(*query.Search)).Count(&totalData).Error

	if err != nil {
		return nil, 0, err
	}

	return categories, totalData, nil
}

func (r *CategoryRepository) FindByNames(category_names []string) (categories []Models.Category, err error) {
	if err := r.DB.Where("category_name IN (?)", category_names).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) Create(category *Models.Category) (*Models.Category, error) {
	if err := r.DB.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) Update(category *Models.Category, id string) (*Models.Category, error) {
	if err := r.DB.Where("id=?", id).Updates(category).Update("status", category.Status).Error; err != nil {
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
