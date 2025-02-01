package Response

import "2024_akutansi_project/Models"

type JurnalEntriesResponse struct {
	Name    string  `json:"name"`
	Ammount float64 `json:"ammount"`
	Types   string  `json:"type"`
}

type ProfitlLossResponse struct {
	JurnalEntry  []JurnalEntriesResponse `json:"jurnal_entry"`
	TotalIncome  float64                 `json:"total_income"`
	TotalExpense float64                 `json:"total_expense"`
	Amount       float64                 `json:"amount"`
	IsProfit     bool                    `json:"is_profit"`
}

func ToProfilLossResponse(jurnalEntries []*Models.JournalEntry, Income, Expense float64) ProfitlLossResponse {
	jurnalEntriesResponse := []JurnalEntriesResponse{}
	for _, jurnalEntry := range jurnalEntries {
		jurnalEntriesResponse = append(jurnalEntriesResponse, JurnalEntriesResponse{
			Name:    "dont'know",
			Ammount: jurnalEntry.Amount,
			Types:   string(jurnalEntry.Account.Type),
		})
	}
	return ProfitlLossResponse{
		JurnalEntry:  jurnalEntriesResponse,
		TotalIncome:  Income,
		TotalExpense: Expense,
		Amount:       Income - Expense,
		IsProfit:     Income > Expense,
	}
}
