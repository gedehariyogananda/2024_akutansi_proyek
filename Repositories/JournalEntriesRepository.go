package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Utils"
	"time"

	"gorm.io/gorm"
)

type (
	IJournalEntriesRepository interface {
		FindAll(request Dto.GetJournalRequest, selectedFields *[]string, isFinancialReport bool) ([]*Models.JournalEntry, int64, error)
		FindByCompanyID(companyID string) (journalEntry *Models.JournalEntry, err error)
		CalculateAmountByType(companyID string, prefixType Models.JournalType) (countAmount float64, err error)
		CheckupUnbalance(companyID string) (bool, error)
		Insert(journalEntries []Models.JournalEntry, trx *gorm.DB) error
		FindByPeriode(year int) ([]*Models.JournalEntry, error)
		GetIncomeOrExpenseYear(year int, types string) (float64, error)
	}

	JournalEntriesRepository struct {
		DB *gorm.DB
	}
)

func JournalEntriesProvider(db *gorm.DB) *JournalEntriesRepository {
	return &JournalEntriesRepository{DB: db}
}

func (repository *JournalEntriesRepository) FindAll(request Dto.GetJournalRequest, selectedFields *[]string, isFinancialReport bool) ([]*Models.JournalEntry, int64, error) {
	var journalEntries []*Models.JournalEntry
	var totalData int64

	query := repository.DB.Model(&Models.JournalEntry{}).
		Where("journal_entries.company_id = ?", request.CompanyID).
		Joins("LEFT JOIN accounts ON journal_entries.account_id = accounts.id")

	if isFinancialReport {
		query = query.Where("accounts.type IN ?", []string{
			string(Models.ASSET), string(Models.LIABILITY), string(Models.EQUITY),
		})
	}

	if selectedFields != nil {
		query = query.Select(*selectedFields)
	}

	query = query.Preload("Account", func(accPayload *gorm.DB) *gorm.DB {
		return accPayload.Select("id", "name", "type")
	})

	if request.AccountID != "" {
		query = query.Where("journal_entries.account_id = ?", request.AccountID)
	}

	if request.StartDate != "" {
		query = query.Where("DATE(journal_entries.date) >= ?", request.StartDate)
	}

	if request.EndDate != "" {
		query = query.Where("DATE(journal_entries.date) <= ?", request.EndDate)
	}

	if request.PeriodDate != "" {
		query = query.Where("DATE(journal_entries.date) <= ?", request.PeriodDate)
	}

	query.Count(&totalData)

	query = query.Scopes(
		Utils.Paginate(request.Page, request.Limit)).
		Order("journal_entries.date desc")

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

func (repository *JournalEntriesRepository) FindByPeriode(year int) ([]*Models.JournalEntry, error) {
	var journalEntries []*Models.JournalEntry

	startOfYear := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := time.Date(year, time.December, 31, 23, 59, 59, 999999999, time.UTC)

	if err := repository.DB.
		Preload("Account").
		Where("date BETWEEN ? AND ?", startOfYear, endOfYear).
		Find(&journalEntries).Error; err != nil {
		return nil, err
	}

	return journalEntries, nil
}

func (repository *JournalEntriesRepository) GetIncomeOrExpenseYear(year int, types string) (float64, error) {
	var totalIncome float64

	startOfYear := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := time.Date(year, time.December, 31, 23, 59, 59, 999999999, time.UTC)

	if err := repository.DB.
		Preload("Account", "type = ?", types).
		Model(&Models.JournalEntry{}).
		Where("date BETWEEN ? AND ?", startOfYear, endOfYear).
		Select("sum(amount)").
		Scan(&totalIncome).Error; err != nil {
		return 0, err
	}

	return totalIncome, nil
}
