package Response

import "2024_akutansi_project/Models"

type Account struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	CompanyID string `json:"company_id"`
	Code      string `json:"code"`
	IsLocked  bool   `json:"is_locked"`
	Status    bool   `json:"status"`
}

func ToAccount(account *Models.Account) *Account {
	return &Account{
		ID:        account.ID,
		Name:      account.Name,
		Type:      string(account.Type),
		CompanyID: account.CompanyID,
		Code:      account.Code,
		IsLocked:  account.IsLocked,
		Status:    account.Status,
	}
}

func ToAccountSlice(accounts []*Models.Account) []*Account {
	accountResponses := []*Account{}
	for _, account := range accounts {
		accountResponses = append(accountResponses, ToAccount(account))
	}
	return accountResponses
}
