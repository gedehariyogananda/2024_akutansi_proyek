package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"gorm.io/gorm"
)

type (
	IWaitingListRepository interface {
		InsertWaitingList(request *Dto.MakeWaitingListRequest) (err error)
	}

	WaitingListRepository struct {
		DB *gorm.DB
	}
)

func WaitingListRepositoryProvider(db *gorm.DB) *WaitingListRepository {
	return &WaitingListRepository{DB: db}
}

func (r *WaitingListRepository) InsertWaitingList(request *Dto.MakeWaitingListRequest) (err error) {
	waitingList := Models.WaitingList{
		Name:         request.Name,
		Company:      request.Company,
		Phone:        request.Phone,
		Email:        request.Email,
		BusinessType: request.BusinessType,
	}

	if err := r.DB.Create(&waitingList).Error; err != nil {
		return err
	}

	return nil
}
