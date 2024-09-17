package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	ICategoryRepository interface {
		FindAll(company_id string) (category *[]Models.Category, err error)
		FindByName(category_name string) (category *Models.Category, err error)
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

func (r *CategoryRepository) FindByName(category_name string) (category *Models.Category, err error) {
	if err := r.DB.Where("category_name = ?", category_name).First(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
