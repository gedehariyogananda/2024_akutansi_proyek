package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Utils"
	"gorm.io/gorm"
)

type (
	IJournalEntriesRepository interface {
		FindAll(request Dto.GetJournalRequest) ([]*Models.JournalEntry, int64, error)
		Insert(journalEntries []Models.JournalEntry) error
	}

	JournalEntriesRepository struct {
		DB *gorm.DB
	}
)

func JournalEntriesProvider(db *gorm.DB) *JournalEntriesRepository {
	return &JournalEntriesRepository{DB: db}
}

func (repository *JournalEntriesRepository) FindAll(request Dto.GetJournalRequest) ([]*Models.JournalEntry, int64, error) {
	var journalEntries []*Models.JournalEntry
	var totalData int64

	query := repository.DB.Model(&Models.JournalEntry{}).Preload("Account")

	if request.AccountID != "" {
		query = query.Where("account_id = ?", request.AccountID)
	}

	if request.StartDate != "" {
		query = query.Where("date >= ?", request.StartDate)
	}

	if request.EndDate != "" {
		query = query.Where("date <= ?", request.EndDate)
	}

	query.Count(&totalData)

	query = query.Scopes(
		Utils.Paginate(request.Page, request.Limit)).
		Order("date desc")

	if err := query.Find(&journalEntries).Error; err != nil {
		return nil, 0, err
	}

	return journalEntries, totalData, nil
}

func (repository *JournalEntriesRepository) Insert(journalEntries []Models.JournalEntry) error {
	return repository.DB.Create(&journalEntries).Error
}
