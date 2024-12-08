package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Utils"

	"gorm.io/gorm"
)

type (
	IAccountRepository interface {
		FindByID(id string) (account *Models.Account, err error)
		Create(account *Models.Account) (*Models.Account, error)
		Delete(id string) error
		Update(account *Models.Account, id string) (*Models.Account, error)
		FindAll(companyID string, qeury *Common.Query) (accounts []*Models.Account, totalData int64, err error)
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
	if err := r.DB.Model(account).Where("id = ?", id).Updates(account).Error; err != nil {
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
		Helper.FilterIslock(query.IsLocked)).
		Find(&accounts).Error

	if err != nil {
		return nil, 0, err
	}

	err = r.DB.Model(&Models.Account{}).Scopes(Helper.FilterStatus(query.Status),
		Helper.FilterCompanyID(companyID),
		Helper.FilterSearch(*query.Search),
		Helper.FilterIslock(query.IsLocked),
	).
		Count(&totalData).Error

	if err != nil {
		return nil, 0, err
	}

	return accounts, totalData, nil
}
