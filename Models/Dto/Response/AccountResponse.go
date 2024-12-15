package Response

import "2024_akutansi_project/Models"

type Profile struct {
	ID          string  `json:"id"`
	Username    string  `json:"username,omitempty"`
	Phone       string  `json:"phone,omitempty"`
	Name        string  `json:"name,omitempty"`
	Email       string  `json:"email,omitempty"`
	CompanyID   string  `json:"company_id"`
	CompanyName string  `json:"company_name"`
	EmployeeKey *string `json:"employee_key,omitempty"`
	IsEmployee  bool    `json:"is_employee"`
}

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
