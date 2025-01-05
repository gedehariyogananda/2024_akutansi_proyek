package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Utils"
	"fmt"

	"gorm.io/gorm"
)

type (
	IAccountRepository interface {
		FindByID(id string) (account *Models.Account, err error)
		Create(account *Models.Account) (*Models.Account, error)
		Delete(id string) error
		Update(account *Models.Account, id string) (*Models.Account, error)
		FindAll(companyID string, qeury *Common.Query) (accounts []*Models.Account, totalData int64, err error)
		GetAll(companyID string) ([]*Models.Account, error)
		InsertDefaultAccounts(companyID string) error
		FindByCode(companyID string, code string) (*Models.Account, error)
	}
	AccountRepository struct {
		DB *gorm.DB
	}
)

func AccountProvider(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (r *AccountRepository) FindByID(id string) (account *Models.Account, err error) {
	account = &Models.Account{}

	if err := r.DB.Where("id = ?", id).First(account).Error; err != nil {
		return nil, err
	}

	return account, nil
}

func (r *AccountRepository) Create(account *Models.Account) (*Models.Account, error) {
	if err := r.DB.Create(account).Error; err != nil {
		return nil, err
	}

	return account, nil
}

func (r *AccountRepository) Update(account *Models.Account, id string) (*Models.Account, error) {
	if err := r.DB.Model(account).Where("id = ?", id).Updates(account).Update("status", account.Status).Error; err != nil {
		return nil, err
	}

	return account, nil
}

func (r *AccountRepository) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&Models.Account{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) FindAll(companyID string, query *Common.Query) (accounts []*Models.Account, totalData int64, err error) {
	err = r.DB.Scopes(Utils.Paginate(query.Page, query.Limit),
		Helper.FilterCompanyID(companyID),
		Helper.FilterStatus(query.Status),
		Helper.FilterSearch(*query.Search),
		Helper.FilterIslock(query.IsLocked),
		Helper.FilterTypeAccount(*query.TypeAccount)).
		Find(&accounts).Error

	if err != nil {
		return nil, 0, err
	}

	fmt.Println(len(accounts))

	err = r.DB.Model(&Models.Account{}).Scopes(Helper.FilterStatus(query.Status),
		Helper.FilterCompanyID(companyID),
		Helper.FilterSearch(*query.Search),
		Helper.FilterIslock(query.IsLocked),
		Helper.FilterTypeAccount(*query.TypeAccount),
	).
		Count(&totalData).Error

	if err != nil {
		return nil, 0, err
	}

	return accounts, totalData, nil
}

func (r *AccountRepository) InsertDefaultAccounts(companyID string) error {
	accounts := []*Models.Account{
		{
			Name:      Models.AccountCash,
			Type:      Models.ASSET,
			CompanyID: companyID,
			Code:      Models.AccountCashCode,
			Status:    true,
		},
		{
			Name:      Models.AccountRevenue,
			Type:      Models.REVENUE,
			CompanyID: companyID,
			Code:      Models.AccountRevenueCode,
			Status:    true,
		},
		{
			Name:      Models.AccountOutputTax,
			Type:      Models.LIABILITY,
			CompanyID: companyID,
			Code:      Models.AccountOutputTaxCode,
			Status:    true,
		},
		{
			Name:      Models.AccountInputTax,
			Type:      Models.ASSET,
			CompanyID: companyID,
			Code:      Models.AccountInputTaxCode,
			Status:    true,
		},
		{
			Name:      Models.AccountProductMaterial,
			Type:      Models.ASSET,
			CompanyID: companyID,
			Code:      Models.AccountProductMaterialCode,
			Status:    true,
		},
		{
			Name:      Models.AccountBusinessDebt,
			Type:      Models.LIABILITY,
			CompanyID: companyID,
			Code:      Models.AccountBusinessDebtCode,
			Status:    true,
		},
		{
			Name:      Models.AccountBusinessCapital,
			Type:      Models.EQUITY,
			CompanyID: companyID,
			Code:      Models.AccountBusinessCapitalCode,
			Status:    true,
		},
		{
			Name:      Models.AccountReceivables,
			Type:      Models.ASSET,
			CompanyID: companyID,
			Code:      Models.AccountReceivablesCode,
			Status:    true,
		},
		{
			Name:      Models.AccountAssets,
			Type:      Models.ASSET,
			CompanyID: companyID,
			Code:      Models.AccountAssetsCode,
			Status:    true,
		},
		{
			Name:      Models.AccountCompanyExpense,
			Type:      Models.EXPENSE,
			CompanyID: companyID,
			Code:      Models.AccountCompanyExpenseCode,
			Status:    true,
		},
		{
			Name:      Models.AccountWithdrawal,
			Type:      Models.EQUITY,
			CompanyID: companyID,
			Code:      Models.AccountWithdrawalCode,
			Status:    true,
		},
	}

	for _, account := range accounts {
		if err := r.DB.Create(account).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *AccountRepository) FindByCode(companyID string, code string) (*Models.Account, error) {
	account := &Models.Account{}

	if err := r.DB.Where("code = ? AND company_id = ?", code, companyID).
		First(account).Error; err != nil {
		return nil, err
	}

	return account, nil
}

func (r *AccountRepository) GetAll(companyID string) ([]*Models.Account, error) {
	var accounts []*Models.Account

	if err := r.DB.Where("company_id = ?", companyID).Find(&accounts).Error; err != nil {
		return nil, err
	}

	return accounts, nil
}
