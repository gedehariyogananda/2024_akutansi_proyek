package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Utils"
	"fmt"
	"time"
)

type (
	IJournalEntriesService interface {
		FindAll(request Dto.GetJournalRequest) (res []*Models.JournalEntry, meta Common.Meta, err error)
		InsertJournalCashier(params Common.JournalEntryParams, isPaid bool) (err error)
		InsertJournalPurchase(params Common.JournalEntryParams, isPaid bool) (err error)
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
	// checkup unbalance
	if err := service.unbalanceCheckup(params.CompanyID); err != nil {
		return err
	}

	acc, err := service.allAccountCompanyUser(params.CompanyID)
	if err != nil {
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
		service.createJournalEntry(acc[Models.AccountCashCode].ID, params, cashType, params.SubTotal+params.Tax), // kas
		service.createJournalEntry(acc[Models.AccountOutputTaxCode].ID, params, outputTaxType, params.Tax),       // pajak luaran
		service.createJournalEntry(acc[Models.AccountRevenueCode].ID, params, revenueType, params.SubTotal),      // pendapatan
	}

	// Insert Journal Entry
	if err := service.JournalEntriesRepository.Insert(res, nil); err != nil {
		return err
	}

	return nil
}

func (service *JournalEntriesService) InsertJournalPurchase(params Common.JournalEntryParams, isPaid bool) (err error) {
	// checkup unbalance
	if err := service.unbalanceCheckup(params.CompanyID); err != nil {
		return err
	}

	acc, err := service.allAccountCompanyUser(params.CompanyID)
	if err != nil {
		return err
	}

	productMaterialType := Models.DEBIT
	inputTaxType := Models.DEBIT
	cashType := Models.CREDIT
	bussinessDebtType := Models.CREDIT

	res := []Models.JournalEntry{
		service.createJournalEntry(acc[Models.AccountProductMaterialCode].ID, params, productMaterialType, params.SubTotal), // bahan produk
		service.createJournalEntry(acc[Models.AccountInputTaxCode].ID, params, inputTaxType, params.Tax),                    // pajak masukan
	}

	if isPaid {
		res = append(res, service.createJournalEntry(acc[Models.AccountCashCode].ID, params, cashType, params.SubTotal+params.Tax)) // kas
	} else {
		res = append(res, service.createJournalEntry(acc[Models.AccountBusinessDebtCode].ID, params, bussinessDebtType, params.SubTotal+params.Tax)) // hutang usaha
	}

	// Insert Journal Entry
	if err := service.JournalEntriesRepository.Insert(res, nil); err != nil {
		return err
	}

	return nil
}

func (service *JournalEntriesService) createJournalEntry(accountID string, params Common.JournalEntryParams, journalType Models.JournalType, amount float64) Models.JournalEntry {

	now := time.Now()
	if params.Date != nil {
		now = *params.Date
	}

	dataDate := Utils.SeperateDate(now.Format("2006-01-02"))

	return Models.JournalEntry{
		AccountID:       accountID,
		Amount:          amount,
		Type:            journalType,
		CompanyID:       params.CompanyID,
		Note:            params.Note,
		Date:            now,
		TransactionCode: "TRX-" + Utils.GenerateUniqueSuffix() + "-" + fmt.Sprintf("%d", *dataDate.Year),
		AdditionalData:  params.AdditionalData,
	}
}

func (service *JournalEntriesService) allAccountCompanyUser(companyID string) (map[string]*Models.Account, error) {
	accounts, _, err := service.AccountRepository.FindAll(companyID, nil)
	if err != nil {
		return nil, err
	}

	acc := make(map[string]*Models.Account)

	for _, account := range accounts {
		acc[account.Code] = account
	}

	if len(acc) == 0 {
		return nil, fmt.Errorf("E_ACCOUNT_NOT_FOUND")
	}

	return acc, nil
}

func (service *JournalEntriesService) unbalanceCheckup(companyID string) (err error) {
	unbalance, err := service.JournalEntriesRepository.CheckupUnbalance(companyID)
	if err != nil {
		return err
	}

	if !unbalance {
		return fmt.Errorf("E_JOURNAL_ENTRY_UNBALANCE")
	}

	return nil
}
