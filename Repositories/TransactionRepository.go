package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Utils"

	"gorm.io/gorm"
)

type (
	ITransactionRepository interface {
		GetAllByCompany(query *Common.Query) (data []*Models.Transaction, totalData int64, err error)
		Store(transaction *Models.Transaction) (data *Models.Transaction, err error)
		CountByDateCompanyID(companyID string, date string) (int64, error)
		GetByDate(date string, selectedFields *[]string) (transactions []*Models.Transaction, err error)
	}

	TransactionRepository struct {
		DB *gorm.DB
	}
)

func TransactionRepositoryProvider(DB *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: DB}
}

func (r *TransactionRepository) GetAllByCompany(query *Common.Query) (data []*Models.Transaction, totalData int64, err error) {
	if err := r.DB.Model(&Models.Transaction{}).
		Where("company_id = ?", query.CompanyID).
		Scopes(Helper.FilterTransaksiLain(query.Search)).
		Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Model(&Models.Transaction{}).
		Where("company_id = ?", query.CompanyID).
		Scopes(
			Utils.Paginate(query.Page, query.Limit),
			Helper.FilterTransaksiLain(query.Search),
		).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, totalData, nil

}

func (r *TransactionRepository) Store(transaction *Models.Transaction) (data *Models.Transaction, err error) {
	if err := r.DB.Create(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *TransactionRepository) CountByDateCompanyID(companyID string, date string) (int64, error) {
	var total int64

	if err := r.DB.Model(&Models.Transaction{}).
		Where("company_id = ?", companyID).
		Where("date = ?", date).
		Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (r *TransactionRepository) GetByDate(date string, selectedFields *[]string) (transactions []*Models.Transaction, err error) {

	query := r.DB.Model(&Models.Transaction{}).
		Where("DATE(created_at) = ?", date)

	if selectedFields != nil {
		query = query.Select(*selectedFields)
	}

	if err := query.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
