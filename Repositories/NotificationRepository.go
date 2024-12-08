package Repositories

import (
	"2024_akutansi_project/Models"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type (
	INotificationRepository interface {
		StoreNotification(ctx context.Context, payload Models.Notification) error
		UpdateNotification(ctx context.Context, payload Models.Notification) error
		FindNotification(ctx context.Context, payload Models.Notification) *Models.Notifications
	}

	NotificationRepository struct {
		DB *gorm.DB
	}
)

func NotificationRepositoryProvider(DB *gorm.DB) *NotificationRepository {
	return &NotificationRepository{DB: DB}
}

func (r *NotificationRepository) StoreNotification(ctx context.Context, payload Models.Notification) error {
	tx := r.DB.Begin()

	if err := tx.WithContext(ctx).Create(payload).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *NotificationRepository) UpdateNotification(ctx context.Context, payload Models.Notification) error {
	tx := r.DB.Begin()

	if err := tx.WithContext(ctx).
		Model(&Models.Notification{}).
		Where("id = ?", payload.ID).
		Updates(payload).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *NotificationRepository) FindNotification(ctx context.Context, payload Models.Notification) *Models.Notifications {
	var dataEntity Models.Notifications

	db := r.DB.Session(&gorm.Session{
		Logger: logger.Default.LogMode(logger.Info),
	}).WithContext(ctx)

	findData := db.WithContext(ctx).Model(&Models.Notification{})

	if payload.UserID != "" {
		findData.Where("user_id = ?", payload.UserID)
	}

	if payload.Scheme != "" {
		findData.Where("scheme = ?", payload.Scheme)
	}

	if payload.Status != "" {
		findData.Where("status = ?", payload.Status)
	}

	findData.Where("sent_at IS NULL")

	findData = findData.Where("scheduled_at BETWEEN ? AND ?",
		time.Now().Add(-5*time.Minute),
		time.Now().Add(5*time.Minute))

	findData = findData.Debug()

	if err := findData.Find(&dataEntity).Error; err != nil {
		return nil
	}

	return &dataEntity
}
