package Repositories

import (
	"2024_akutansi_project/Models"

	"gorm.io/gorm"
)

type (
	IAccountRepository interface {
		FindByID(id string) (account *Models.Account, err error)
		Create(account *Models.Account) (*Models.Account, error)
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
