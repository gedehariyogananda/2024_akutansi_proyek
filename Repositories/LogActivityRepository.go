package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	ILogActivityRepository interface {
		FindByUserId(userID string) (logActivity []*Models.LogActivity, err error)
		Create(logActivity *Models.LogActivity) (err error)
	}

	LogActivityRepository struct {
		DB *gorm.DB
	}
)

func LogActivityRepositoryProvider(db *gorm.DB) *LogActivityRepository {
	return &LogActivityRepository{DB: db}
}

func (h *LogActivityRepository) FindByUserId(userID string) (logActivity []*Models.LogActivity, err error) {
	if err := h.DB.
		Where("user_id = ?", userID).
		Find(&logActivity).Error; err != nil {
		return nil, err
	}

	return logActivity, nil
}

func (h *LogActivityRepository) Create(logActivity *Models.LogActivity) (err error) {
	if err := h.DB.Create(logActivity).Error; err != nil {
		return err
	}

	return nil
}
