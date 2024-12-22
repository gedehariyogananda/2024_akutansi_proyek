package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"fmt"
)

type (
	IJournalEntriesService interface {
		FindAll(request Dto.GetJournalRequest) (res []*Models.JournalEntry, meta Common.Meta, err error)
		InsertJournalCashier(params Common.JournalEntryParams, isPaid bool) (err error)
	}

	JournalEntriesService struct {
		JournalEntriesRepository Repositories.IJournalEntriesRepository
		AccountRepository        Repositories.IAccountRepository
	}
)

func JournalEntriesProvider(journalRepo Repositories.IJournalEntriesRepository, accountRepo Repositories.IAccountRepository) *JournalEntriesService {
	return &JournalEntriesService{
		JournalEntriesRepository: journalRepo,
		AccountRepository:        accountRepo,
	}
}

func (service *JournalEntriesService) FindAll(request Dto.GetJournalRequest) (res []*Models.JournalEntry, meta Common.Meta, err error) {
	res, totalData, err := service.JournalEntriesRepository.FindAll(request)
	if err != nil {
		return nil, Common.Meta{}, err
	}

	meta = Common.PaginateMetadata(nil, totalData, request.Limit, request.Page)

	return res, meta, nil
}

func (service *JournalEntriesService) InsertJournalCashier(params Common.JournalEntryParams, isPaid bool) (err error) {
	params = Common.JournalEntryParams{
		AccountID:       params.AccountID,
		SubTotal:        params.SubTotal,
		Tax:             params.Tax,
		CompanyID:       params.CompanyID,
		Note:            params.Note,
		TransactionCode: params.TransactionCode,
		AdditionalData:  params.AdditionalData,
	}

	cashAccount, _ := service.FindAccountIDByCode(params.CompanyID, Models.AccountCashCode)
	outputTaxAccount, _ := service.FindAccountIDByCode(params.CompanyID, Models.AccountOutputTaxCode)
	revenueAccount, _ := service.FindAccountIDByCode(params.CompanyID, Models.AccountRevenueCode)

	// checkup unbalance
	if _, err := service.UnbalanceCheckup(params.CompanyID); err != nil {
		return err
	}

	cashType := Models.DEBIT
	outputTaxType := Models.CREDIT
	revenueType := Models.CREDIT

	if !isPaid {
		cashType = Models.CREDIT
		outputTaxType = Models.DEBIT
		revenueType = Models.DEBIT
	}

	res := []Models.JournalEntry{
		{
			AccountID:       cashAccount.ID, // kas
			Amount:          params.SubTotal + params.Tax,
			Type:            cashType,
			CompanyID:       params.CompanyID,
			Note:            params.Note,
			TransactionCode: params.TransactionCode,
			AdditionalData:  params.AdditionalData,
		},
		{
			AccountID:       outputTaxAccount.ID, // Pajak Luaran
			Amount:          params.Tax,
			Type:            outputTaxType,
			CompanyID:       params.CompanyID,
			Note:            params.Note,
			TransactionCode: params.TransactionCode,
			AdditionalData:  params.AdditionalData,
		},
		{
			AccountID:       revenueAccount.ID, // Pendapatan
			Amount:          params.SubTotal,
			Type:            revenueType,
			CompanyID:       params.CompanyID,
			Note:            params.Note,
			TransactionCode: params.TransactionCode,
			AdditionalData:  params.AdditionalData,
		},
	}

	// Insert Journal Entry
	if err := service.JournalEntriesRepository.Insert(res, nil); err != nil {
		return err
	}

	return nil
}

func (service *JournalEntriesService) FindAccountIDByCode(companyID string, code string) (res *Models.Account, err error) {
	res, err = service.AccountRepository.FindByCode(companyID, code)
	if err != nil {
		return nil, fmt.Errorf("E_ACCOUNT_NOT_FOUND")
	}

	return res, nil
}

func (service *JournalEntriesService) UnbalanceCheckup(companyID string) (safety bool, err error) {
	unbalance, err := service.JournalEntriesRepository.CheckupUnbalance(companyID)
	if err != nil {
		return false, err
	}

	if !unbalance {
		return false, fmt.Errorf("E_JOURNAL_ENTRY_UNBALANCE")
	}

	return true, nil
}
