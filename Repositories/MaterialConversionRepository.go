package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IMaterialConversionRepository interface {
		Create(materialConversion *Models.MaterialConversion) (*Models.MaterialConversion, error)
	}
	MaterialConversionRepository struct {
		DB *gorm.DB
	}
)

func MaterialConversionRepositoryProvider(db *gorm.DB) *MaterialConversionRepository {
	return &MaterialConversionRepository{DB: db}
}

func (r *MaterialConversionRepository) Create(materialConversion *Models.MaterialConversion) (*Models.MaterialConversion, error) {
	if err := r.DB.Create(materialConversion).Error; err != nil {
		return nil, err
	}

	return materialConversion, nil
}
