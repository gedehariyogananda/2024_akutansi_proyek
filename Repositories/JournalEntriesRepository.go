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
		FindByCompanyID(companyID string) (journalEntry *Models.JournalEntry, err error)
		CalculateAmountByType(companyID string, prefixType Models.JournalType) (countAmount float64, err error)
		CheckupUnbalance(companyID string) (bool, error)
		Insert(journalEntries []Models.JournalEntry, trx *gorm.DB) error
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

func (repository *JournalEntriesRepository) FindByCompanyID(companyID string) (journalEntry *Models.JournalEntry, err error) {
	journalEntry = &Models.JournalEntry{}

	if err := repository.DB.
		Where("company_id = ?", companyID).
		First(journalEntry).Error; err != nil {
		return nil, err
	}

	return journalEntry, nil
}

func (repository *JournalEntriesRepository) Insert(journalEntries []Models.JournalEntry, trx *gorm.DB) error {
	db := trx
	if db == nil {
		db = repository.DB
	}

	if err := db.Create(&journalEntries).Error; err != nil {
		return err
	}

	return nil
}

func (repository *JournalEntriesRepository) CalculateAmountByType(companyID string, prefixType Models.JournalType) (countAmount float64, err error) {

	if err := repository.DB.
		Model(&Models.JournalEntry{}).
		Where("company_id = ? AND type = ?", companyID, prefixType).
		Select("sum(amount)").
		Scan(&countAmount).Error; err != nil {
		return 0, err
	}

	return countAmount, nil
}

func (repository *JournalEntriesRepository) CheckupUnbalance(companyID string) (bool, error) {
	var countDebit float64
	var countCredit float64

	countDebit, err := repository.CalculateAmountByType(companyID, Models.DEBIT)
	if err != nil {
		return false, err
	}

	countCredit, err = repository.CalculateAmountByType(companyID, Models.CREDIT)
	if err != nil {
		return false, err
	}

	return countDebit == countCredit, nil
}
