package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Utils"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type (
	IJournalEntriesService interface {
		FindAll(request Dto.GetJournalRequest) (res []*Models.JournalEntry, meta Common.Meta, err error)
		InsertJournalCashier(params Common.JournalEntryParams, isPaid bool, trx *gorm.DB) (err error)
		InsertJournalPurchase(params Common.JournalEntryParams, isPaid bool, trx *gorm.DB) (err error)
		InsertJournalOtherTransaction(ctx context.Context) (err error)
	}

	JournalEntriesService struct {
		JournalEntriesRepository Repositories.IJournalEntriesRepository
		AccountRepository        Repositories.IAccountRepository
		TransactionRepository    Repositories.ITransactionRepository
	}
)

func JournalEntriesProvider(journalRepo Repositories.IJournalEntriesRepository, accountRepo Repositories.IAccountRepository, transactionRepo Repositories.ITransactionRepository) *JournalEntriesService {
	return &JournalEntriesService{
		JournalEntriesRepository: journalRepo,
		AccountRepository:        accountRepo,
		TransactionRepository:    transactionRepo,
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

func (service *JournalEntriesService) InsertJournalCashier(params Common.JournalEntryParams, isPaid bool, trx *gorm.DB) (err error) {
	isContinued, err := service.unbalanceCheckup(params.CompanyID)
	if err != nil {
		return err
	}

	if !isContinued {
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

	if err := service.JournalEntriesRepository.Insert(res, trx); err != nil {
		return err
	}

	return nil
}

func (service *JournalEntriesService) InsertJournalPurchase(params Common.JournalEntryParams, isPaid bool, trx *gorm.DB) (err error) {

	isContinued, err := service.unbalanceCheckup(params.CompanyID)
	if err != nil {
		return err
	}

	if !isContinued {
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

	if err := service.JournalEntriesRepository.Insert(res, trx); err != nil {
		return err
	}

	return nil
}

func (service *JournalEntriesService) InsertJournalOtherTransaction(ctx context.Context) (err error) {
	now := time.Now()
	currentDate := now.Format("2006-01-02")

	transactions, err := service.TransactionRepository.GetByDate(currentDate, &[]string{"id", "payment_type", "payment_method", "amount", "company_id", "title"})
	if err != nil {
		return err
	}

	for _, transaction := range transactions {
		var entries []Models.JournalEntry

		note := transaction.Title + "_" + transaction.PaymentType + "_" + transaction.PaymentMethod

		params := Common.JournalEntryParams{
			CompanyID:      transaction.CompanyID,
			Note:           note,
			Date:           &transaction.Date,
			AdditionalData: transaction.AdditionalData,
		}

		acc, err := service.allAccountCompanyUser(transaction.CompanyID)
		if err != nil {
			return err
		}

		if transaction.PaymentType == "" {
			return fmt.Errorf("tipe transaksi kosong!")
		}

		switch Models.TypeTransactionClient(transaction.PaymentType) {
		case Models.TIPE_WITHDRAWAL: // Case 1: Prive
			entries = []Models.JournalEntry{
				service.createJournalEntry(acc[Models.AccountWithdrawalCode].ID, params, Models.DEBIT, transaction.Amount), // Debit: Prive
				service.createJournalEntry(acc[Models.AccountCashCode].ID, params, Models.CREDIT, transaction.Amount),      // Kredit: Kas
			}
		case Models.TIPE_CAPITAL_ADDITION: // Case 2: Penambahan Modal
			entries = []Models.JournalEntry{
				service.createJournalEntry(acc[Models.AccountCashCode].ID, params, Models.DEBIT, transaction.Amount),         // Debit: Kas
				service.createJournalEntry(acc[Models.AccountBusinessCapital].ID, params, Models.CREDIT, transaction.Amount), // Kredit: Modal Usaha
			}
		case Models.TIPE_DEBT_PAYMENT: // Case 3: Pembayaran Hutang
			entries = []Models.JournalEntry{
				service.createJournalEntry(acc[Models.AccountBusinessDebtCode].ID, params, Models.DEBIT, transaction.Amount), // Debit: Hutang Usaha
				service.createJournalEntry(acc[Models.AccountCashCode].ID, params, Models.CREDIT, transaction.Amount),        // Kredit: Kas
			}
		case Models.TIPE_RECEIVABLE_PAYMENT: // Case 4: Pembayaran Piutang
			entries = []Models.JournalEntry{
				service.createJournalEntry(acc[Models.AccountCashCode].ID, params, Models.DEBIT, transaction.Amount),         // Debit: Kas
				service.createJournalEntry(acc[Models.AccountReceivablesCode].ID, params, Models.CREDIT, transaction.Amount), // Kredit: Piutang Usaha
			}
		case Models.TIPE_SALE: // Case 5 & 8: Penjualan
			if transaction.PaymentMethod == string(Models.PAYMENT_METHOD_CASH) { // Metode Cash
				entries = []Models.JournalEntry{
					service.createJournalEntry(acc[Models.AccountCashCode].ID, params, Models.DEBIT, transaction.Amount),     // Debit: Kas
					service.createJournalEntry(acc[Models.AccountRevenueCode].ID, params, Models.CREDIT, transaction.Amount), // Kredit: Pendapatan
				}
			} else { // Metode Hutang
				entries = []Models.JournalEntry{
					service.createJournalEntry(acc[Models.AccountReceivablesCode].ID, params, Models.DEBIT, transaction.Amount), // Debit: Piutang Usaha
					service.createJournalEntry(acc[Models.AccountRevenueCode].ID, params, Models.CREDIT, transaction.Amount),    // Kredit: Pendapatan
				}
			}
		case Models.TIPE_PURCHASE: // Case 6 & 7: Pembelian
			if transaction.PaymentMethod == string(Models.PAYMENT_METHOD_CASH) {
				entries = []Models.JournalEntry{
					service.createJournalEntry(acc[Models.AccountAssetsCode].ID, params, Models.DEBIT, transaction.Amount), // Debit: Aset
					service.createJournalEntry(acc[Models.AccountCashCode].ID, params, Models.CREDIT, transaction.Amount),  // Kredit: Kas
				}
			} else {
				entries = []Models.JournalEntry{
					service.createJournalEntry(acc[Models.AccountAssetsCode].ID, params, Models.DEBIT, transaction.Amount),        // Debit: Aset
					service.createJournalEntry(acc[Models.AccountBusinessDebtCode].ID, params, Models.CREDIT, transaction.Amount), // Kredit: Hutang Usaha
				}
			}
		case Models.TIPE_TAX_PAYMENT: // Case 9: Pembayaran Pajak
			entries = []Models.JournalEntry{
				service.createJournalEntry(acc[Models.AccountOutputTaxCode].ID, params, Models.DEBIT, transaction.Amount), // Debit: Pajak Keluaran
				service.createJournalEntry(acc[Models.AccountCashCode].ID, params, Models.CREDIT, transaction.Amount),     // Kredit: Kas
			}
		case Models.TYPE_EXPENSE: // Case 10 & 11: Pembayaran Beban
			if transaction.PaymentMethod == string(Models.PAYMENT_METHOD_CASH) { // Metode Cash
				entries = []Models.JournalEntry{
					service.createJournalEntry(acc[Models.AccountCompanyExpenseCode].ID, params, Models.DEBIT, transaction.Amount), // Debit: Beban Perusahaan
					service.createJournalEntry(acc[Models.AccountCashCode].ID, params, Models.CREDIT, transaction.Amount),          // Kredit: Kas
				}
			} else { // Metode Hutang
				entries = []Models.JournalEntry{
					service.createJournalEntry(acc[Models.AccountCompanyExpenseCode].ID, params, Models.DEBIT, transaction.Amount), // Debit: Beban Perusahaan
					service.createJournalEntry(acc[Models.AccountBusinessDebtCode].ID, params, Models.CREDIT, transaction.Amount),  // Kredit: Hutang Usaha
				}
			}
		default:
			fmt.Printf("E_UNSUPPORT_TYPE: %s\n", transaction.PaymentType)
			continue
		}

		if err := service.JournalEntriesRepository.Insert(entries, nil); err != nil {
			return err
		}
	}

	return nil
}

func (service *JournalEntriesService) createJournalEntry(accountID string, params Common.JournalEntryParams, journalType Models.JournalType, amount float64) Models.JournalEntry {

	now := time.Now()
	if params.Date != nil {
		now = *params.Date
	}

	creditAt := params.CreditAt
	debitAt := params.DebitAt

	if journalType == Models.CREDIT {
		creditAt = &amount
	} else {
		debitAt = &amount
	}

	dataDate := Utils.SeperateDate(now.Format("2006-01-02"))

	return Models.JournalEntry{
		AccountID:       accountID,
		Amount:          amount,
		Type:            journalType,
		CompanyID:       params.CompanyID,
		Note:            params.Note,
		Date:            now,
		CreditAt:        creditAt,
		DebitAt:         debitAt,
		TransactionCode: "TRX-" + Utils.GenerateUniqueSuffix() + "-" + fmt.Sprintf("%d", *dataDate.Year),
		AdditionalData:  params.AdditionalData,
	}
}

func (service *JournalEntriesService) allAccountCompanyUser(companyID string) (map[string]*Models.Account, error) {
	accounts, err := service.AccountRepository.GetAll(companyID)
	if err != nil {
		return nil, fmt.Errorf("E_ACCOUNT_NOT_FOUND")
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

func (service *JournalEntriesService) unbalanceCheckup(companyID string) (isContinued bool, err error) {
	data, err := service.JournalEntriesRepository.FindByCompanyID(companyID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
	}

	if data != nil {
		// checkup unbalance
		unbalance, err := service.JournalEntriesRepository.CheckupUnbalance(companyID)
		if err != nil {
			return false, err
		}

		if !unbalance {
			return false, fmt.Errorf("E_JOURNAL_ENTRY_UNBALANCE")
		}
	}

	return true, nil
}
