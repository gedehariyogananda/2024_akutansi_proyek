package Dto

import "2024_akutansi_project/Models/Common"

type GetJournalRequest struct {
	Common.Query
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	AccountID string `json:"account_id"`
}
